package letsgo

import "fmt"

// Simple Add function

func add(x int, y int) int {
	return x + y
}

// Multiple return with basic swap

func swap(x, y int) (int, int) {
	return y, x
}

// Named Return
func declare() (xx int, yy int) {
	xx = 100
	yy = 200
	return
}

// Simple variable declaration
func saysomething() string {
	var word string
	fmt.Println("Say", word)

	// shorthand variable
	newWord := "Hello"
	fmt.Println("Say again", newWord)

	return newWord
}

// Switch case
func simpleCheckBill(x int) {
	switch x {
	case 10:
		fmt.Println("too bad")
	case 20:
		fmt.Println("normal")
	}
}

// If case
func simpleCheckBillwithIf(x int) {
	if x < 10 {
		fmt.Println("less than 10")
	} else if x < 20 {
		fmt.Println("less than 20")
	} else {
		fmt.Println("else")
	}
}

// Defer
func deferCase(x int) {
	fmt.Println("Open Connection", x)
	defer fmt.Println("Defer connection")
	defer fmt.Println("Defer another connection")
}

// Pointer
func mutate(i *int) {
	fmt.Printf("in mutate function : %p \n", i)
	*i = 200
}

// Struct
type Coordinate struct {
	X int
	Y int
}

// Array
func arrReceiver(arr *[3]int) {
	fmt.Printf("in arrReceiver >>> %p\n", arr)
}
