
package main

import "fmt"

func main(){
	var itf interface{}

	itf = "a"

	fmt.Println(itf)

	itf = 6

	fmt.Println(itf)

	m := map[int]string{
		666: "liuliuliu",
	}

	itf = m

	fmt.Println(itf)
}