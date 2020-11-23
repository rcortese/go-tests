package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

here:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here
			}
			fmt.Println(j)
			if j == 2 {
				break
			}
		}
	}

there:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if j == 1 {
				continue
			}
			fmt.Println(j)
			if j == 2 {
				break there
			}
		}
	}
}
