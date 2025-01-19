package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Println("FOR LOOP")
	sum := 0
	// the init statement(optional) ; the condition expression; the post statement(optional) *optional see forAnotherStyle func
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	forAnotherStyle()

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		ifPow(3, 2, 10),
		ifPow(3, 3, 20),
	)

	fmt.Println(
		ifPow2(3, 2, 10),
		ifPow2(3, 3, 20),
	)

	// https://go.dev/tour/flowcontrol/8
	fmt.Println(Sqrt(16))

	// https://go.dev/tour/flowcontrol/9
	switchStatement()

	// https://go.dev/tour/flowcontrol/10
	switchStatement2()

	// https: //go.dev/tour/flowcontrol/11
	switchStatement3()

	// https://go.dev/tour/flowcontrol/12
	testDefer()

	// https://go.dev/tour/flowcontrol/13
	stackDefer()

}

func forAnotherStyle() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	for {
	}
}

// sample if statment. condition doesnt need to be surounded by () but braces {} are needed
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Like for, the if statement can start with a short statement to execute before the condition.

// Variables declared by the statement are only in scope until the end of the if.

// (Try using v in the last return statement.)

func ifPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// If and else
// Variables declared inside an if short statement are also available inside any of the else blocks.

// (Both calls to pow return their results before the call to fmt.Println in main begins.)

func ifPow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// https://go.dev/tour/flowcontrol/8
func Sqrt(x float64) float64 {
	var z float64
	for z = 1.0; z <= 10; z++ {
		value := z - (z*z-x)/(2*z)
		fmt.Println("value")
		fmt.Println(value)
	}
	return z

}

// https://go.dev/tour/flowcontrol/9
func switchStatement() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// https://go.dev/tour/flowcontrol/10
func switchStatement2() {

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}

// Switch with no condition
// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-else chains.

func switchStatement3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer
// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

func testDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
// To learn more about defer statements read this blog post.
// https://go.dev/blog/defer-panic-and-recover

// https://go.dev/tour/flowcontrol/13
func stackDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
