package main

import (
	"fmt"
)

type state struct {
	edgd1 *state
	edge2 *state
}

func poregtonfa(postfix) {

}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}