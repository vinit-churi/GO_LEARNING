package main

import (
	"reflect"
	"testing"
)

func TestDescribeStruct(t *testing.T) {
	if got := DescribeStruct(Service{}); got != "Service:2" {
		t.Fatalf("DescribeStruct() = %q", got)
	}
}

func TestFieldNames(t *testing.T) {
	got := FieldNames(Service{})
	want := []string{"Name", "Active"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FieldNames() = %#v, want %#v", got, want)
	}
}

func TestIsZeroValue(t *testing.T) {
	if !IsZeroValue(Service{}) {
		t.Fatal("IsZeroValue(Service{}) = false")
	}
	if IsZeroValue(Service{Name: "billing"}) {
		t.Fatal("IsZeroValue(non-zero) = true")
	}
}
