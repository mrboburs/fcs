package main

import "fmt"

func main() {
	i := 90
	p := &i
	*p = 34
	fmt.Println(i)
	fmt.Println("i", &i)
	fmt.Println(&p)
}
