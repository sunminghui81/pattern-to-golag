package main

import "fmt"


func a_append_ok(){
	var a []int
	aa := make([]int, 2)
	a = append(a, 1)
	aa = append(aa, 1)
	aa[0] = 4
	fmt.Println(a, aa)
}

func a_append_err(){
	var a [2]int
	// a = append(a, 1)
	fmt.Println(a)
}


func main(){
	var a1  [5]int = [5]int{1,2,3,4,5}
	// var a2 [5]int = [1,2]
	// var a2 = [1,2,3]
	var a2 []int = []int{1,2}
	var a3 [5]int = [5]int{1,2}
	a4 := [5]int{3:1}
	fmt.Println(a1, a2, a3, a4)
	a_append_err()
	a_append_ok()
}