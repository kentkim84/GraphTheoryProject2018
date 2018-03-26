package main

import (
	"fmt"
)
// linked list data structure
type state struct {
	symbol rune
	edge1 *state // point to other state
	edge2 *state // point to other state
}

// 
type nfa struct {
	initial *state
	accept	*state
}

// postfix regular expression to nfa
func poregtonfa(postfix string) *nfa {
	nfaStack := []*nfa{} // create nfa stack

	for _, r := range postfix {
		switch r {
		case '.': // concat
			frag2 := nfaStack[len(nfaStack)-1] // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one
			frag1 := nfaStack[len(nfaStack)-1] // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) // push the fragment
		case '|': // union
			frag2 := nfaStack[len(nfaStack)-1] // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one
			frag1 := nfaStack[len(nfaStack)-1] // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one
			
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge2 = &accept
			
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the fragment
		case '*': // kleene
			frag := nfaStack[len(nfaStack)-1] // pop the last element from the stack
			nfaStack = nfaStack[:len(nfaStack)-1] // everything but the last one

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the fragment
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) // push the fragment
		}
	}

	return nfaStack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}