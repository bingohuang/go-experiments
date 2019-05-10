#!/bin/sh

umask 027

USER=$(whoami)
OPERIP=$(whoami | awk '{print $NF}' | sed 's/[()]//g')
base_file=$(basename $0)
current_dir=$(cd "$(dirname "$0")"; pwd)
SCRIPT_PATH=$0
SCRIPT_OWNER="`stat -c '%U' ${SCRIPT_PATH}`"
CURRENT_USER="`/usr/bin/id -u -n`"

if [ "${SCRIPT_OWNER}" != "${CURRENT_USER}" ]
then
    echo "only ${SCRIPT_OWNER} can execute this script."
    exit 1
fi

start_ngx() {
    . ${current_dir}/init.sh nginx

    ${PYTHONHOME}/bin/python ${NGINX_ROOT}/bin/bus_adm.pyc
}

start_ra() {
    return 0
}

start_ra_health() {
     return 0
}

start_nginx_health() {
    . ${current_dir}/init.sh nginx

    ret=`netstat -apn|grep ${MGRMT_PORT}|grep LISTEN |wc -l`
    if [ $ret -ne 0 ];then
        pid=`ps -ef | grep -v grep | grep $APP_ROOT | grep master | sed -n '1P' | awk '{print $2}'`
        pids=`ps -ef |grep $pid |grep -v grep  | awk '{print $2}'`
        for procid in $pids
        do
            dstate=`cat /proc/${procid}/status | grep State | awk '{print $2}'`
            if [[ ("${dstate}x" == "Zx") || ("${dstate}x" == "Tx") ]]; then
                /bin/echo "check process is unhealthy"
                return 1
            fi
        done
    fi
    return 0
}

start_monitor(){
    . ${current_dir}/init.sh nginx

    ${PYTHONHOME}/bin/python ${ROUTERAGENT_ROOT}/rtsp/hrscommon/tools/pyscript/bus/tracemonitor.pyc $ROLE
}

certificate_check() {
    . ${current_dir}/init.sh python

    ret=`${PYTHONHOME}/bin/python ${ROUTERAGENT_ROOT}/bin/script/certificate_check.pyc`

    if [ ! -f $SSL_ROOT/cert_file ];then
        touch $SSL_ROOT/cert_file
    fi
    cert=`cat $SSL_ROOT/cert_file`

    if [ -z "$cert" -a -n "$ret" ]; then
        echo $ret > $SSL_ROOT/cert_file
        return 0
    fi

    if [ "$cert" != "$ret" ]; then
        echo $ret > $SSL_ROOT/cert_file
        return 1
    else
        return 0
    fi

}

monitor_certificate(){
    /bin/sh ${current_dir}/start.sh certificate_check
    ret=$?
    if [ $ret -ne 0 ]; then
        ps -ef|grep -v grep |grep routeragent-0-0 | awk '{print $2}' | xargs kill -9
        /bin/sh ${current_dir}/start.sh python

        isKill=$1
        if [ $isKill -eq 1 ];then
            ps -ef|grep -v grep | grep nginx |awk '{print $2}' | xargs kill -9
            /bin/sh ${current_dir}/start.sh nginx
        fi
    fi

}

is_master()
{
    local_ip=`ip addr |grep -w ${HOST_FLOAT_IP} |awk '{print $2}'`
    if [[ $local_ip != "" ]]; then
       return 1
    else
       return 0
    fi
 }
 
do_loop_check()
{
        log "Successful, do loop check."
        while [ 1 ]; do
            TIME_INFO="`date +%m-%d` `date +%H:%M:%S`"
            /bin/echo "$TIME_INFO start monitor..."

            /bin/sh ${current_dir}/start.sh monitor

            /bin/sh ${current_dir}/start.sh nginx_health
            ret=$?
            if [ $ret -ne 0 ]; then
                ps -ef|grep -v grep | grep nginx |awk '{print $2}' | xargs kill -9
                /bin/sh ${current_dir}/start.sh nginx

            fi
            j=`ps -ef|grep -v grep | grep master`
            if [ -z "${j}" ]; then
	        ps -ef|grep -v grep | grep nginx |awk '{print $2}' | xargs kill -9
                /bin/sh ${current_dir}/start.sh nginx
            fi

            sleep 30
        done
}

do_active_standby_check()
{
        log "Successful,do active standby check."
        num=0
        while [ 1 ]; do
            TIME_INFO="`date +%m-%d` `date +%H:%M:%S`"
            /bin/echo "$TIME_INFO start monitor..."
            num=`expr $num + 5`
            if [ $num -gt 30 ]; then

                /bin/sh ${current_dir}/start.sh monitor

                is_master
                ret=$?
                if [ $ret -eq 1 ];then
                    /bin/sh ${current_dir}/start.sh nginx_health
                    result=$?
                    if [ $result -ne 0 ]; then
                        ps -ef|grep -v grep | grep nginx |awk '{print $2}' | xargs kill -9
                        /bin/sh ${current_dir}/start.sh nginx

                    fi
                fi
                num=0
            fi
            
            is_master
            ret=$?
            if [ $ret -eq 1 ]; then
               j=`ps -ef|grep -v grep | grep nginx | grep master`
               if [ -z "${j}" ]; then
	           ps -ef|grep -v grep | grep nginx |awk '{print $2}' | xargs kill -9
                   /bin/sh ${current_dir}/start.sh nginx
               fi
            else
               j=`ps -ef|grep -v grep | grep nginx | grep master`
               if [ "${j}" ]; then
                   /bin/sh ${current_dir}/stop.sh nginx
               fi
            fi
            sleep 5
        done
}

log() {
    logger -t $USER -p local0.info "${base_file};$1;${OPERIP:-127.0.0.1};Start Sidecar"
}

main() {
    case "$1" in
    nginx)
        start_ngx
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    python)
        start_ra
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    monitor)
        start_monitor
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    ra_health)
        start_ra_health
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    nginx_health)
        start_nginx_health
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    certificate_check)
        certificate_check
        ret=$?
        if [ $ret -ne 0 ]; then
            exit $ret
        fi
        ;;
    *)
      if [ -z ${HOST_FLOAT_IP} ];then
         /bin/echo "HOST_FLOAT_IP is empty"
      else
         IS_ACTIVE_STANDBY=true
         export LOCALHOST_FLOATSERVERNAME=${HOST_FLOAT_IP}
      fi

        /bin/sh ${current_dir}/start.sh nginx
        ret=$?
        if [ $ret -ne 0 ]; then
            is_master
            IS_MASTER=$?
            if [ -z $IS_ACTIVE_STANDBY ]||[ $IS_MASTER -ne 0 ];then
               log "Failed"
               exit $ret
            fi
        fi
        if [ -z $IS_ACTIVE_STANDBY ];then
            do_loop_check
        else
            do_active_standby_check
        fi
        ;;
    esac
}

main $1

exit 0
