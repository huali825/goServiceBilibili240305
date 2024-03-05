package suanfa

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int
	Name string
	Age  uint8
}

func (u *User) Call() {
	fmt.Println(" user is called function..")
	fmt.Printf("%v\n", u)
}

func DoFiledAndMethod(input any) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Println(field.Name, field.Type, value)
	}
}

func TestSuanfa004(t *testing.T) {
	u := User{
		Id:   1,
		Name: "Tom",
		Age:  28,
	}
	DoFiledAndMethod(u)
}
