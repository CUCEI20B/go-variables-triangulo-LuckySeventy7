package main

import "fmt"

func main()  {
	var a float64
	var b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	d := (a * b) * 0.5
	fmt.Println(d)
}