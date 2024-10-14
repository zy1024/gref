// gref_test.go
package gref

import (
	"fmt"
	"reflect"
	"testing"
)

type PersonContainAddress struct {
	Name string
	Age  int
	Addr string
}

type Person struct {
	Name string
	Age  int
}

type Src struct {
	Name     string
	Age      string
	Password string
	Friends  []PersonContainAddress
	Parent   PersonContainAddress
	Brother  *PersonContainAddress
	Sister   []*PersonContainAddress
}

type Dst struct {
	Name    string
	Age     float64
	Friends []Person
	Parent  Person
	Brother *Person
	Sister  []*Person
}

// example
func TestCopyStructFields(t *testing.T) {
	src := &Src{
		Name:     "A",
		Age:      "12",
		Password: "123456",
		Friends:  make([]PersonContainAddress, 0),
		Parent: PersonContainAddress{
			Name: "parent-A",
			Age:  32,
			Addr: "address-A",
		},
		Brother: &PersonContainAddress{
			Name: "brother-A",
			Age:  13,
			Addr: "address-A",
		},
	}
	src.Friends = append(src.Friends, PersonContainAddress{Name: "B", Age: 13, Addr: "address-B"})
	src.Friends = append(src.Friends, PersonContainAddress{Name: "C", Age: 14, Addr: "address-C"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A1", Age: 15, Addr: "address-E"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A2", Age: 16, Addr: "address-D"})

	dst := &Dst{}
	err := CopyStructFields(src, dst)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(src)
	fmt.Println("---------------------------------------")
	fmt.Println(dst)
	fmt.Println(dst.Brother)
	for _, person := range dst.Sister {
		fmt.Println(person)
	}
}

func TestCopyValue(t *testing.T) {
	var a = 1
	var b string
	var c = []int32{1, 2, 3}
	var d []string

	err := CopyBasicValue(&a, &b)
	if err != nil {
		t.Error(err)
	}

	err = CopyBasicValue(&c, &d)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(d, reflect.TypeOf(d))
}
