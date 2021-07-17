package main

import (
	"errors"
	"fmt"
	"math"
)

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("bro you cant give me negative values")
	}
	return math.Sqrt(x), nil
}

type person struct {
	name string
	age  int
}

func inc(x *int) {
	*x++
}

func main() {
	fmt.Println("I'm awesome")
	var x int = 10
	fmt.Println(x)
	if x > 6 {
		fmt.Println("The given number is greater than 6")
	}

	fmt.Println("Lets declare an array")
	a := [4]int{5, 4, 3, 1}
	fmt.Println(a)

	fmt.Println("Now lets dig in hashmaps")
	first_map := make(map[string]int)
	first_map["Triangle"] = 3
	first_map["Rectangle"] = 4
	fmt.Println(first_map["Triangle"])
	delete(first_map, "Rectangle")

	fmt.Println("Now coming to loops")
	for i := 0; i < 4; i++ {
		fmt.Println(a[i])
	}

	fmt.Println("While Loop")
	var i int = 0
	for i < 4 {
		fmt.Println(a[i])
		i++
	}

	fmt.Println("Ranged for loops")

	for index, value := range a {
		fmt.Println("Index", index, "Value", value)
	}

	fmt.Println("Dealing with functions")
	fmt.Println(sum(4, 6))

	// Errors and multiple return values
	fmt.Println("Dealing with errors and multiple return values as well")
	result, err := sqrt(-5)
	fmt.Println(err)
	fmt.Println(result)

	// Dealing with structs

	p := person{name: "Nithin", age: 21}
	fmt.Println(p.name)

	// Dealing with pointers
	var a_number int = 7

	fmt.Println(a_number)
	inc(&a_number)
	fmt.Println(a_number)

}
