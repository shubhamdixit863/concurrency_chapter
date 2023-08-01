package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
}

func (us User) Hello(name string) {
	fmt.Println("Hey user,Method invoked at runtime ---------------------")
}

func ParseJson() User {
	str := []byte(`{
    "id": 20,
    "name": "Henderson Branch",
    "salary": "$30751"
  }`)
	v := User{}

	err := json.Unmarshal(str, &v)
	if err != nil {
		fmt.Println(err)
		return User{}
	}

	return v
}

// reflect package
func main() {

	// Different operations we can perform on the basis of reflection
	//1-Getting the type of value
	//x := 3.14
	//fmt.Println(reflect.TypeOf(x))

	//2-Getting value at run time
	//fmt.Println(reflect.ValueOf(x))

	// 3- Using kind method
	//v := reflect.ValueOf(x) // the value of x ,whatever the value that is lying inside x

	//fmt.Println(v.Kind())

	// Changing the valu eat runtime
	//v := reflect.ValueOf(&x)
	//fmt.Println(v) // Address where x is stored --->
	//v.Elem().SetFloat(0.29)
	//fmt.Println(x)

	//fmt.Println(ParseJson())
	v := reflect.ValueOf(ParseJson())
	method := v.MethodByName("Hello") // Getting the valu eof method by name
	args := []reflect.Value{reflect.ValueOf("Anything")}

	// Mocking a
	method.Call(args)
	t := reflect.TypeOf(ParseJson())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		nameField := t.Field(i)

		fmt.Printf("Filed %d %v %s\n", i, field, nameField.Name)
	}

}
