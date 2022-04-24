package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3
	println(c)
	println(d)
	var f float32 = 32799438743.8297

	fmt.Printf("f %f\n", f)
	fmt.Printf("g %g\n", f)
	fmt.Println(f)
	fmt.Printf("v %v\n", f)
	fmt.Printf("T %T\n", f)

	var e = 324.13455
	var g = 3.14

	fmt.Printf("%05.2f\n", e) // 00324.13
	fmt.Printf("%08.3g\n", e) // 03.2e+02
	fmt.Printf("%8.5g\n", e)  // 324.13
	fmt.Printf("%f\n", g)     // 3.140000

}
