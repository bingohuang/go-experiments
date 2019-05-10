package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	data := 3
	for i, v := range slice {
		if v == data {
			fmt.Println(i)
			fmt.Println(slice[:i])

			fmt.Println(slice[i+1:]) //  容量=0 ok
			fmt.Println(slice[i+1:i+1]) //  容量=0 ok

			fmt.Println(slice[i+1:i+2]) // error

			fmt.Println(slice[i+2:i+2]) // error

			fmt.Println(slice[i+1]) // error
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(slice)
}
