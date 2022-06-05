package main

import (
	"fmt"
	"reflect"
)

//Should reflection be used?
// Clear is better than clever. Reflection is never clear.

type Banana struct {
	Size  string `foo:"size,omitempty" json:"size,omitempty"`
	Color string `foo:"color"`
	Age   int    `foo:"age"`
	Alias string `foo:"-"`
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	structToSqlInsert(employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	})
	structToSqlInsert(90)

	readStructTags()

}

func readStructTags() {
	println("\n\n\n\n\n TAGS")
	banana := Banana{
		Size:  "XL",
		Color: "YELLOW",
		Age:   10,
		Alias: "BAN",
	}

	v := reflect.ValueOf(banana)
	for i := 0; i < v.NumField(); i++ {
		structField := reflect.TypeOf(&banana).Elem().Field(i)
		tagInfo := structField.Tag.Get("foo")
		fmt.Println(fmt.Sprintf("%s:%s", structField.Name, tagInfo))
	}
}

func structToSqlInsert(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		tableName := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", tableName)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Println(field)
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}
