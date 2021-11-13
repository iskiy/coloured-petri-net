package main

import (
	. "coloured-petri-net/pkg"
	"fmt"
	"math/rand"
	"time"
)

const (
	blue  = "blue"
	green = "green"
	red   = "red"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	place0 := Place{PlacesID: 0, Tokens: []string{red, green, blue}}
	place1 := Place{PlacesID: 1, Tokens: []string{red, green, blue}}
	place2 := Place{PlacesID: 2, Tokens: []string{}}
	place3 := Place{PlacesID: 3, Tokens: []string{}}
	place4 := Place{PlacesID: 4, Tokens: []string{}}
	place5 := Place{PlacesID: 5, Tokens: []string{}}

	var transition0 = Transition{
		TransitionID: 0,
		FromArcs:     []Arc{{red, 0, 0}},
		ToArcs:       []Arc{{red, 0, 3}},
	}
	var transition1 = Transition{
		TransitionID: 1,
		FromArcs: []Arc{{blue, 1, 1},
			{green, 1, 1},
			{blue, 0, 1},
		},
		ToArcs: []Arc{{red, 1, 2}},
	}
	var transition2 = Transition{
		TransitionID: 2,
		FromArcs: []Arc{{red, 2, 2},
			{red, 3, 2},
		},
		ToArcs: []Arc{{blue, 2, 4},
			{green, 2, 5},
		},
	}
	petriNet := PetriNet{Places: []Place{place0, place1, place2, place3, place4, place5},
		Transitions: []Transition{transition0, transition1, transition2}}
	fmt.Println(petriNet.FindExecTransitions())
	fmt.Println(petriNet)
	fmt.Println("\n")
	petriNet.Run()

	fmt.Println(petriNet)
}
