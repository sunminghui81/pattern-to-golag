
package main

import "fmt"



func dictDeclare(){
	var m1 map[string]string 
	if nil == m1 {
		fmt.Println("map not init is nil")
	}
	m1 = make(map[string]string)

	m1["a"] = "apple"
	
	fmt.Println(m1)
}

func dictNew(){
	m1 := make(map[string]string)
	m1["b"] = "boy"
	fmt.Println(m1)
}

func dictInit(){
	m1 := map[string]string{
		"a":"apple",
		"b":"boy",
	}

	m1["c"] = "cat"
	fmt.Println(m1)
}


func main(){
	dictDeclare()
	dictNew()
	dictInit()

}