
package main 

import "fmt"
import "strconv"

func main(){
	var m []string
	if nil == m {
		m = make([]string,10)
		println("init m ok")
	}

	for i := 0; i<10; i++ {
		m[i] = "Num:" + strconv.Itoa(i)
	}

	for i, s := range m {
		fmt.Println(s)
		switch i {
			case 0: fallthrough
			case 1: fallthrough
			case 2, 3: fmt.Println("switch 1, 2, 3")
		}
		switch {
			case i == 4: 
				fmt.Println("switch 4")
				continue
			case i == 9: goto ENDF
			default: continue
		}

		for ;; {
			fmt.Println("dead loop")
		}
	
	}

ENDF:
	fmt.Println("Class is over, Bye")
}