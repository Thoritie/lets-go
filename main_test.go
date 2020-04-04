package letsgo

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test simple add function
func Test_Add(t *testing.T) {
	assert.Equal(t, add(2, 3), 5)
}

// Test function multiple return
func Test_SwapMultipleReturn(t *testing.T) {
	a, b := swap(1, 3)
	assert.Equal(t, 3, a)
	assert.Equal(t, 1, b)
}

// Test Named return
func TestSimpleAssert_CallFunctionNamedReturn(t *testing.T) {
	xx, yy := declare()
	assert.Equal(t, xx, 100)
	assert.Equal(t, yy, 200)
}

// Simple variable declaration
func Test_DeclareVarible(t *testing.T) {
	var golang string
	fmt.Println("golang", golang)

	var c int
	fmt.Println("c", c)

	var html float64
	fmt.Println("html", html)

	var javascript bool
	fmt.Println("javascript", javascript)

	//shorthand variable declaration
	php := 123
	fmt.Println("php", php)

	python := "test"
	fmt.Println("python", python)
}

func Test_Saysomething(t *testing.T) {
	actual := saysomething()
	assert.Equal(t, actual, "Hello")
}

// Type conversion
func TestSimpleAssert_TypeConversion(t *testing.T) {
	targetAsInteger := 123
	fmt.Println(reflect.TypeOf(targetAsInteger))

	targetAsFloat64 := float64(targetAsInteger)
	fmt.Println(reflect.TypeOf(targetAsFloat64))

	aToI, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("error while covert type")
		t.Fail()
	}
	fmt.Println("Atoi: ", aToI, reflect.TypeOf(aToI))
}

// Type Inference
func TestSimpleAssert_TypeInference(t *testing.T) {
	targetAsInteger := 123
	fmt.Println("targetAsInteger", reflect.TypeOf(targetAsInteger))
}

// Constant
func Test_TypeInference(t *testing.T) {
	const Pi = 3.14
	fmt.Println(Pi)
}

// For loop //
func Test_Forloop(t *testing.T) {
	sum := 0
	for i := 0; i < 20; i++ {
		sum = sum + i
	}
	fmt.Println("sum", sum)
}

// Switch case
func Test_Switchcase(t *testing.T) {
	simpleCheckBill(20)
}

// If case
func Test_simpleIf(t *testing.T) {
	simpleCheckBillwithIf(20)
}

// Defer
func Test_DeferCase(t *testing.T) {
	deferCase(10)
}

// Pointer
func Test_pointer(t *testing.T) {
	i := 100

	fmt.Printf("in main test function : %p\n", &i)
	mutate(&i)

	fmt.Println("i", i)
}

/////////////////

// Struct
func Test_Struct(t *testing.T) {
	schoolMap := Coordinate{X: 10, Y: 20}
	fmt.Println(schoolMap)

	schoolMap1 := Coordinate{X: 30, Y: 40}
	fmt.Println(schoolMap1)
}

// Array
func Test_Array(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	fmt.Printf("in main  function >>> %p\n", &arr)
	arrReceiver(&arr)
}

func Test_ArrayLooping(t *testing.T) {
	// array looping
	// for _,value := range
}

////////////////
// Slice
func Test_slice(t *testing.T) {
	sli := [5]int{1, 2, 3}
	fmt.Println(sli)
	fmt.Println("Cap:", cap(sli))
	fmt.Println("Lenght:", len(sli))
}

////////////////
// Map
func getPriceTag(tagName string) float64 {
	var priceTag = make(map[string]float64)
	priceTag["banana"] = 160.5
	return priceTag[tagName]

	// element, ok := priceTag[tagName]
	// if !ok {
	// }
}

func Test_Map(t *testing.T) {
	fmt.Println(getPriceTag("banana"))
	fmt.Println(getPriceTag("apple"))
}

/////////////
// function value

type godFunction func(int, int) int

func doSum(fn godFunction, x int, y int) int {
	return fn(x, y)
}

func Test_FunctionValue(t *testing.T) {
	sumFn := func(x, y int) int {
		return x + y
	}
	fmt.Println(doSum(sumFn, 1, 2))
}

// Method
type Student struct {
	X int
	Y int
}

func (s Student) getSum() int {
	return s.X + s.Y
}

func Test_Method(t *testing.T) {
	schoolMap := Student{X: 14, Y: 12}
	fmt.Println(schoolMap.getSum())
}

// Interface
type animal interface {
	say()
}

func doSay(fni animal) {
	fni.say()
}

type Cat struct {
	Name string
}

func (c Cat) say() {
	fmt.Println(c.Name, " is cat")
}

type Unicorn struct {
	Name string
}

func (u Unicorn) say() {
	fmt.Println(u.Name, " is unicorn")
}

func Test_Interface(t *testing.T) {
	summer := Cat{Name: "Summer"}
	doSay(summer)

	violet := Unicorn{Name: "Violet"}
	doSay(violet)
}

//////////////////
/// type Assertion

func doPrint(value interface{}) {
	fmt.Println(value)
	fmt.Println(reflect.TypeOf(value))

}
func Test_typeAssertion(t *testing.T) {
	var i interface{} = 123
	_, ok := i.(float64)
	if !ok {
		fmt.Println("is error")
	}
}
