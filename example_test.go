// gref_test.go
package gref

import (
	"encoding/json"
	"fmt"
	"github.com/zy1024/gref/utils"
	"reflect"
	"testing"
	"time"
)

type PersonContainAddress struct {
	Name     string
	Age      int
	Addr     string
	HomeTown HomeTown
}

type Person struct {
	Name     string
	Age      int
	HomeTown HomeTown
}

type HomeTown struct {
	Name    string
	Country string
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
			HomeTown: HomeTown{
				Name:    "HomeTown-A",
				Country: "China",
			},
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
			Name:     "Ignore",
			Age:      0,
			HomeTown: HomeTown{Country: "EN"},
		},
	}

	fmt.Println("isZero:", utils.IsZero(reflect.ValueOf(dst)))

	var averageTime time.Duration
	var totalTime time.Duration

	for i := 0; i < 1000; i++ {

		startTime := time.Now()

		// Operating efficiency test
		err := Copy(src, dst)
		if err != nil {
			t.Error(err)
			return
		}

		endTime := time.Now()
		totalTime += endTime.Sub(startTime)

	}
	averageTime = totalTime / 1000
	fmt.Println("totalTime:", totalTime)
	fmt.Println("averageTime:", averageTime)

	fmt.Println(src)
	fmt.Println("---------------------------------------")
	fmt.Println("dst:", dst, reflect.TypeOf(dst))
	fmt.Println("dst.Name:", dst.Name, reflect.TypeOf(dst.Name))
	fmt.Println("dst.Age:", dst.Age, reflect.TypeOf(dst.Age))
	fmt.Println("dst.Friends:", dst.Friends, reflect.TypeOf(dst.Friends))
	fmt.Println("dst.Parent:", dst.Parent, reflect.TypeOf(dst.Parent))
	fmt.Println("dst.Parent.HomeTown:", dst.Parent.HomeTown, reflect.TypeOf(dst.Parent.HomeTown))
	fmt.Println("dst.Brother:", dst.Brother, reflect.TypeOf(dst.Brother))
	fmt.Println("dst.Sister:", dst.Sister, reflect.TypeOf(dst.Sister))
	for i, sister := range dst.Sister {
		fmt.Println("dst.Sister[", i, "]:", sister, reflect.TypeOf(sister))
	}
	fmt.Println("dst.Teacher:", dst.Teacher, reflect.TypeOf(dst.Teacher))
}

func TestCopyAndMarshalEfficiency(t *testing.T) {
	a := PersonContainAddress{
		Name: "A",
		Age:  18,
		Addr: "addr-A",
	}
	b := Person{}

	var averageTime time.Duration
	var totalTime time.Duration
	for i := 0; i < 1000; i++ {
		startTime := time.Now()
		_ = Copy(&a, &b)
		endTime := time.Now()
		totalTime += endTime.Sub(startTime)
	}
	averageTime = totalTime / 1000
	fmt.Println("Copy Spend total time:", totalTime)
	fmt.Println("Copy Spend average time:", averageTime)

	averageTime = 0
	totalTime = 0
	for i := 0; i < 1000; i++ {
		startTime := time.Now()
		aBytes, _ := json.Marshal(&a)
		_ = json.Unmarshal(aBytes, &b)
		endTime := time.Now()
		totalTime += endTime.Sub(startTime)
	}
	averageTime = totalTime / 1000
	fmt.Println("json Marshal Spend total time:", totalTime)
	fmt.Println("json Marshal Spend average time:", averageTime)
}
