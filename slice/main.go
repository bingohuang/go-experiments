package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3}
	slice3 = append([]int{4, 5, 6}, slice3...)
	fmt.Println(slice3)
}
