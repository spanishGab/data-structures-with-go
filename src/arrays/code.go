package main

import "fmt"

func _main() {
	var marks [4][4]int

	for i := 0; i < len(marks); i++ {
		fmt.Printf("Mark %d: %d\n", i, marks[i])
	}

}
