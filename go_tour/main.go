package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// BASIC TYPES
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

// var should alwasy be used when declaring outside a function
// package level var

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool

// Variables declared without an explicit initial value are given their zero value.

// The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

// Initializing multiple variables
var b, j int = 1, 2

func main() {

	var c int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", c, f, b, s)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// function level var
	var i int

	// Initializing multiple variables
	var isC, isPython, isJava = true, false, "no!"

	// := shortcut instead of using var
	// c, python, java := true, false, "no!"

	fmt.Println(isC, isPython, isJava)

	fmt.Println(i, c, python, java)

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// CAPITAL LETTERS ARE EXPORTED AND SMALL LETTERS ARE NOT
	fmt.Println(math.Pi)
	// THE LINE BELOW WILL RETURN AN ERROR
	// fmt.Println(math.Pi)

	fmt.Println(add(1, 2))

	firstName, lastName := swap("Prince", "Escueta")

	fmt.Println(firstName)
	fmt.Println(lastName)

	firstNum, secondNum := split(13)
	fmt.Println(firstNum)
	fmt.Println(secondNum)

	typeConversion()
	typeInference()
	constants()
	numericConstants()

}

func add(x, y int) int {
	return x + y
}

// Can return 2 variables or more
func swap(x, y string) (string, string) {
	return y, x
}

// named returned values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// NAKED RETURN AND WILL RETURN THE NAMED RETURNED VARIABLE (x, y int) should only be used for short functions to avoid readability issues
	return
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference() {
	v := "42" // change me!
	fmt.Printf("v is of type %T\n", v)

	a := 123 // change me!
	fmt.Printf("a is of type %T\n", a)

	b := 123.1 // change me!
	fmt.Printf("b is of type %T\n", b)
}

// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.
const Pi = 3.14

func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
