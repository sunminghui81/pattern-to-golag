package main 

import "fmt"
import "errors"

func errUsage() error {
	var err error
	if nil == err{
		fmt.Println("error not init is nil")
	}
	err = errors.New("this is a new error")
	// err = fmt.Errorf("%s %v","this is another error ", 666)
	// err = errors.New(fmt.Sprintf("%s%v", "你真无聊", 2333))
	return err
}

func main(){
	fmt.Println(errUsage().Error())
}