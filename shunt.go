package main

import (
	"fmt"
)

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	postfix := []rune{} // characters as utf8
	stack := []rune{} // stack

	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // everything in stack except the last element
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // everything in stack except the last element
			}
			stack = append(stack, r)
		default:
			postfix = append(postfix, r)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1]) // takes top(last) of the element then put into postfix
		stack = stack[:len(stack)-1] // everything in stack except the last element
	}

	return string(postfix)
}

func main() {
	// Answer: ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	// Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	// Anser: abd|.c*.
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	// Answer: abb.+c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
}