// gref_test.go
package gref

import (
	"fmt"
	"testing"
	"time"
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
	Name     string                  // same type and name as dst, will be copied directly
	Age      string                  // same name but different type, will be converted to dst type (int)
	Password string                  // dst has no Password field, will be ignored
	Friends  []PersonContainAddress  // both are slices of structs, but structs have different fields; will copy and convert the same fields
	Parent   PersonContainAddress    // both are structs, but have different fields; will copy and convert the same fields
	Brother  *PersonContainAddress   // both are pointers to structs, but structs have different fields; will copy and convert the same fields
	Sister   []*PersonContainAddress // both are slices of pointers to structs, but structs have different fields; will copy and convert the same fields
	Teacher  []Person                // both are slices, but have different data types; this will not be assigned and will be ignored
}

type Dst struct {
	Name    string
	Age     int
	Friends []Person
	Parent  Person
	Brother *Person
	Sister  []*Person
	Teacher []string
}

// example
func TestCopy(t *testing.T) {
	src := &Src{
		Name:     "A",
		Age:      "12",
		Password: "123456",
		Parent: PersonContainAddress{
			Name: "parent-A",
			Age:  32,
			Addr: "address-A1",
		},
		Brother: &PersonContainAddress{
			Name: "brother-A",
			Age:  13,
			Addr: "address-A1",
		},
	}
	src.Friends = append(src.Friends, PersonContainAddress{Name: "friend-A1", Age: 13, Addr: "address-A2"})
	src.Friends = append(src.Friends, PersonContainAddress{Name: "friend-A2", Age: 14, Addr: "address-A3"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A1", Age: 15, Addr: "address-A1"})
	src.Sister = append(src.Sister, &PersonContainAddress{Name: "sister-A2", Age: 16, Addr: "address-A1"})
	src.Teacher = append(src.Teacher, Person{Name: "teacher-A", Age: 30})

	dst := &Dst{
		Brother: &Person{ // dst has data, src data will be ignored
			Name: "Ignore",
			Age:  10,
		},
	}

	startTime := time.Now()

	// Operating efficiency test
	//for i := 0; i < 1; i++ {
	//	srcBytes, err := json.Marshal(src)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	err = json.Unmarshal(srcBytes, dst)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	err := Copy(src, dst)
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//}

	err := Copy(src, dst)
	if err != nil {
		t.Error(err)
		return
	}

	endTime := time.Now()

	fmt.Println("time:", endTime.Sub(startTime))

	fmt.Println(src)
	fmt.Println("---------------------------------------")
	fmt.Println(dst)
	fmt.Println(dst.Brother)
	for _, person := range dst.Sister {
		fmt.Println(person)
	}
	fmt.Println(len(dst.Teacher))
}
