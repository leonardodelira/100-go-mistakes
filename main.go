package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]

	s2 = append(s2, 33)

	fmt.Println(s1)
	fmt.Println(s2)
}
