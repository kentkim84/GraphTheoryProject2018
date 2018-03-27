package main

import (
	"fmt"
)

// linked list data structure
type state struct {
	symbol rune
	edge1  *state // point to other state
	edge2  *state // point to other state, only used if there is two edges
}

//
type nfa struct {
	initial *state // track the fragment of initial state
	accept  *state // track the fragment of final/accept state
}

// postfix regular expression to nfa
func poregToNfa(postfix string) *nfa {
	nfaStack := []*nfa{} // create nfa stack

	for _, r := range postfix {
		switch r {
		case '.': // The Concatenation Operator - N.M
			frag2 := nfaStack[len(nfaStack)-1]    // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one
			frag1 := nfaStack[len(nfaStack)-1]    // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one

			frag1.accept.edge1 = frag2.initial // join the accept state of fragement 1 to the initial state of fragment 2

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) // push the states of two fragments
		case '|': // The Alternation Operator (union/or) - N|M
			frag2 := nfaStack[len(nfaStack)-1]    // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one
			frag1 := nfaStack[len(nfaStack)-1]    // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one

			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept // join the accept state of fragment 1 to the final accept state
			frag2.accept.edge1 = &accept // join the accept state of fragment 2 to the final accept state

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the states of two fragments
		case '*': // The Match-zero-or-more Operator (kleene star) - N*
			frag := nfaStack[len(nfaStack)-1]     // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			// this fragment has two edges
			frag.accept.edge1 = frag.initial // join the accept state of fragment to the initial state
			frag.accept.edge2 = &accept      // join the accept state of fragment to the accept state

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the states of a fragment
		default: // any characters, but special ones
			accept := state{}
			initial := state{symbol: r, edge1: &accept}                           // read a standard charater or empty string (r)
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the states of a fragment
		}
	}
	if len(nfaStack) != 1 {
		fmt.Println("HEY! somthing wrong!: ", len(nfaStack), nfaStack)
	}

	return nfaStack[0]
}

func addState(listOfState []*state, initialState *state, acceptState *state) []*state {
	listOfState = append(listOfState, initialState)

	if initialState != acceptState && initialState.symbol == 0 { // check that not in the accept state and get E arrow coming from
		listOfState = addState(listOfState, initialState.edge1, acceptState)
		if initialState.edge2 != nil { // check if there is the second edge
			listOfState = addState(listOfState, initialState.edge2, acceptState)
		}
	}

	return listOfState
}

func poMatch(postfix string, matchString string) bool {
	isMatch := false             // initial return value
	poNfa := poregToNfa(postfix) // get nfa from the postfix string

	current := []*state{}
	next := []*state{}

	current = addState(current[:], poNfa.initial, poNfa.accept)

	for _, r := range matchString {
		for _, c := range current {
			if c.symbol == r { // check if current character 'c' matches a character 'r' from matching string
				next = addState(next[:], c.edge1, poNfa.accept)
			}
		}
		// swap the current state to the next state
		// create a new next state
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == poNfa.accept {
			isMatch = true
			break
		}
	}

	return isMatch
}

func main() {
	fmt.Println("Postfix regex: ab.c*| matching string: abc is", poMatch("ab.c*|", "abc"))
	fmt.Println("Postfix regex: ab.c*| matching string: empty string is", poMatch("ab.c*|", ""))
	fmt.Println("Postfix regex: ab.c*| matching string: cccc is", poMatch("ab.c*|", "cccc"))
}
