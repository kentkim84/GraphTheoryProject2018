package main

import (
	"fmt"
)
// ref : https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func intopostRegex(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8} // regular expression precedence	
	postfix := []rune{} // characters as utf8
	stack := []rune{} // stack

	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r) // push token to stack
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1]) // push token to output
				stack = stack[:len(stack)-1] // everything in stack except the last element
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] { // move higher precedence to the left
				postfix = append(postfix, stack[len(stack)-1]) // push token to output
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

func intopostArith(infix string) string {	
	specials := map[rune]int{'^': 4, '*': 3, '/': 3, '+': 2, '-': 2} // arithmetic expression precedence
	postfix := []rune{} // characters as utf8
	stack := []rune{} // stack

	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r) // push token to stack
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1]) // push token to output
				stack = stack[:len(stack)-1] // everything in stack except the last element
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] { // move higher precedence to the left
				postfix = append(postfix, stack[len(stack)-1]) // push token to output
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
	fmt.Println("Regular expression precedence")
	// Answer: ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopostRegex("a.b.c*"))

	// Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopostRegex("(a.(b|d))*"))

	// Anser: abd|.c*.
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopostRegex("a.(b|d).c*"))

	// Answer: abb.+c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopostRegex("a.(b.b)+.c"))

	/* ----------------------------------------------- */
	fmt.Println("Arithmetic expression precedence")
	// Answer: abc*+
	fmt.Println("Infix: ", "a+b*c")
	fmt.Println("Postfix: ", intopostArith("a+b*c"))

	// Answer: abd+*
	fmt.Println("Infix: ", "a*(b+d)")
	fmt.Println("Postfix: ", intopostArith("a*(b+d)"))

	// Anser: abd-/c*
	fmt.Println("Infix: ", "a/(b-d)*c")
	fmt.Println("Postfix: ", intopostArith("a/(b-d)*c"))

	// Answer: abb*/c+
	fmt.Println("Infix: ", "a/(b*b)+c")
	fmt.Println("Postfix: ", intopostArith("a/(b*b)+c"))
}