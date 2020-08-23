package main

import "fmt"

func main() {
	fmt.Print("Hello World\n")
	s := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}

}
