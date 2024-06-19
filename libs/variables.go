package libs

import (
    "fmt"
	"reflect"
	"math/cmplx"
)

	
type person struct {
    Name string
    Age  int
}

var (
	TestString string = "Hello world!"
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func HowToUseVarsInGolang() {

	// Basic
	fmt.Println("---")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	var s string = "Yo!"
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", s, s)

	// Override
	fmt.Println("---")
	s = "Hola"
	fmt.Printf("Type: %T Value: %v\n", s, s)


	// Without default values
	fmt.Println("---")
	var di int
	var df float64
	var db bool
	var ds string
	fmt.Printf("%v %v %v %q\n", di, df, db, ds)

	// Multiple vars
	fmt.Println("---")
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(python), reflect.TypeOf(java))

	// Typed JSON
	fmt.Println("---")
	john := person{Name: "John Smith", Age: 23}
	person1 := []person{person{}}
	fmt.Println(john, reflect.TypeOf(john))
	fmt.Println(person1, reflect.TypeOf(person1))

	// Global vars
	fmt.Println("---")
	fmt.Printf("Type: %T Value: %v\n", TestString, TestString)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
