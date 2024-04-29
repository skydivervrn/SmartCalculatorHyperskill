package main

import "fmt"

func main() {
	var word string
	// the Stack and Queue structs are hidden but already declared
	var stackSolver Stack
	var queueSolver Queue

	fmt.Scan(&word)

	for _, char := range word {
		stackSolver.Push(char)
		queueSolver.Push(char)
	}

	for {
		palindrom := true
		for range word {
			s, _ := stackSolver.Pop()
			q, _ := queueSolver.Pop()
			if s != q {
				palindrom = false
				break
			}
		}

		if palindrom {
			fmt.Println("Palindrome")
			break
		}

		if !palindrom {
			fmt.Println("Not palindrome")
			break
		}
	}
}
