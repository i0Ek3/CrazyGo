package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"001"`
	}

	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("type"); ok {
		fmt.Println(catType.Tag.Get("json"))
	}

	funcVal := reflect.ValueOf(add)
	paraList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(10)}
	retList := funcVal.Call(paraList)
	fmt.Println(retList[0].Int())

}
