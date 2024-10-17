// gref_test.go
package gref

import (
	"fmt"
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
	Teacher  []Person
}

type Dst struct {
	Name    string
	Age     float64
	Friends []Person
	Parent  Person
	Brother *Person
	Sister  []*Person
	Teacher []int
}

// example
func TestCopy(t *testing.T) {
	src := &Src{
		Name:     "A",                             // same type and name as dst , will be copied
		Age:      "12",                            // same name but different type, will be converted to dst type
		Password: "123456",                        // dst has no Password field, will be ignored
		Friends:  make([]PersonContainAddress, 0), // dst has same Friends field but different type, and fewer fields, the same fields will be copied and converted
		Parent: PersonContainAddress{ // same name but different type, will be converted to dst type
			Name: "parent-A",
			Age:  32,
			Addr: "address-A",
		},
		Brother: &PersonContainAddress{ // dst has data, will be ignored
			Name: "brother-A",
			Age:  13,
			Addr: "address-A",
		},
	}
	src.Friends = append(src.Friends, PersonContainAddress{Name: "B", Age: 13, Addr: "address-B"})
	src.Friends = append(src.Friends, PersonContainAddress{Name: "C", Age: 14, Addr: "address-C"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A1", Age: 15, Addr: "address-E"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A2", Age: 16, Addr: "address-D"})
	src.Teacher = append(src.Teacher, Person{Name: "teacher-A", Age: 30})

	dst := &Dst{
		Brother: &Person{
			Name: "",
			Age:  10,
		},
	}

	err := Copy(src, dst)
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
	fmt.Println(len(dst.Teacher))
}
