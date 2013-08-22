// Author: h4ckm03d
// Repository: https://github.com/h4ckm03d/sinau-go
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.

package main

import "fmt"

func Triangle(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func TriangleSimetry(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	for i := 1; i < 10; i++ {
		TriangleSimetry(i)
	}
}
