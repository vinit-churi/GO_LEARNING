//go:build solution
// +build solution

package main

import (
	"fmt"
	"reflect"
)

type Service struct {
	Name   string
	Active bool
}

func DescribeStruct(input any) string {
	t := reflect.TypeOf(input)
	if t == nil {
		return "<nil>"
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return t.String()
	}
	return fmt.Sprintf("%s:%d", t.Name(), t.NumField())
}

func FieldNames(input any) []string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	names := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		names = append(names, t.Field(i).Name)
	}
	return names
}

func IsZeroValue(input any) bool {
	return reflect.ValueOf(input).IsZero()
}

func main() {
	fmt.Println(DescribeStruct(Service{Name: "billing", Active: true}))
}
