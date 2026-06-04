package main

import (
	"fmt"
	"reflect"
	"reflection/employe"
	"reflection/person"
)

func display(s interface{}) {

	// fmt.Printf("type of struct is %T and value is %v\n", s, s)
	// fmt.Println("type: ", reflect.TypeOf(s))
	// fmt.Println("type of kind: ", reflect.TypeOf(s).Kind())
	// fmt.Println("value: ", reflect.ValueOf(s))
	// fmt.Println("value of kind: ", reflect.ValueOf(s).Kind())

	if reflect.ValueOf(s).Kind() == reflect.Struct {
		v := reflect.ValueOf(s)
		t := reflect.TypeOf(s)
		fmt.Println("struct name: ", t.Name())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("field: %d, value: %v, type: %T\n", i, v.Field(i), v.Field(i))
			// fmt.Println("kind of field: ", v.Field(i).Kind())
		}
	} else {
		fmt.Println("invalid type")
	}

}

func main() {

	Employe := employe.New("Vaibhav", 50000, "SDE")
	Person := person.New("Vaibhav", 25, "Sports, coding, travelling")

	display(45)
	display("name")
	display(Employe)
	display(Person)
}
