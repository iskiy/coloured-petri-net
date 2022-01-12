package pkg

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Place struct {
	PlacesID int
	Tokens   []string
}

func (p *Place) containsColour(colour string) (int, bool) {
	for i, t := range p.Tokens {
		if t == colour {
			return i, true
		}
	}
	return -1, false
}

type Arc struct {
	Colour string
	// IDs place/transition
	FromID int
	ToID   int
}

func (a Arc) fromArcString() string {
	return fmt.Sprintf("colour: %s, from placeID: %d", a.Colour, a.FromID)
}

func (a Arc) toArcString() string {
	return fmt.Sprintf("colour: %s, to placeID: %d", a.Colour, a.ToID)
}

type Transition struct {
	TransitionID int
	ToArcs       []Arc
	FromArcs     []Arc
}

func (t Transition) String() string {
	var res string
	res += fmt.Sprintf("TransitionID: %d", t.TransitionID)
	res += "\n\t\tFromArcs: "
	for _, a := range t.FromArcs {
		res += fmt.Sprintf("\n\t\t\t %s;", a.fromArcString())
	}
	res += "\n\t\tToArcs: "
	for _, a := range t.ToArcs {
		res += fmt.Sprintf("\n\t\t\t %s;", a.toArcString())
	}
	return res
}

type PetriNet struct {
	Places      []Place
	Transitions []Transition
	Limit       int
}

func (pn PetriNet) String() string {
	var res = "Places: "
	res += pn.PlacesToString()
	res += "\nTransitions:"
	for _, t := range pn.Transitions {
		res += fmt.Sprintf("\n\t%s", t)
	}
	return res
}

func (pn PetriNet) PlacesToString() string {
	res := ""
	for _, p := range pn.Places {
		res += fmt.Sprintf("\n\t ID: %d, Tokens: [%s]", p.PlacesID, strings.Join(p.Tokens, " "))
	}
	return res
}

func (pn *PetriNet) Run() {
	for i := 0; i < pn.Limit; i++ {
		transToExec := pn.FindExecTransitions()
		if len(transToExec) == 0 {
			break
		}
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(transToExec))
		fmt.Println(fmt.Sprintf("Activation of %d transition", transToExec[randomIndex]))
		pn.execTransition(transToExec[randomIndex])
		fmt.Println(pn.PlacesToString() + "\n")
	}
}

func (pn *PetriNet) execTransition(idx int) {
	for _, a := range pn.Transitions[idx].FromArcs {
		id, _ := pn.Places[a.FromID].containsColour(a.Colour)
		pl, _ := subToken(id, pn.Places[a.FromID])
		pn.Places[a.FromID] = pl
	}
	for _, a := range pn.Transitions[idx].ToArcs {
		pn.Places[a.ToID].Tokens = append(pn.Places[a.ToID].Tokens, a.Colour)
	}
}

func (pn *PetriNet) FindExecTransitions() []int {
	var res []int
	for i := range pn.Transitions {
		if pn.isExecutableTransitions(i) {
			res = append(res, i)
		}
	}
	return res
}

func (pn PetriNet) isExecutableTransitions(idx int) bool {
	placesCopy := createTempPlaces(pn.Places)
	if len(placesCopy) == 0 {
		return false
	}
	for _, arc := range pn.Transitions[idx].FromArcs {
		if !pn.checkToArc(arc, placesCopy) {
			return false
		}
	}
	return true
}

func (pn PetriNet) checkToArc(arc Arc, places []Place) bool {
	if arc.FromID >= len(pn.Places) || arc.ToID >= len(pn.Transitions) {
		return false
	}
	if len((places)[arc.FromID].Tokens) == 0 {
		return false
	}
	idx, ok := (places)[arc.FromID].containsColour(arc.Colour)
	if !ok {
		return false
	}
	pl, err := subToken(idx, (places)[arc.FromID])
	if err != nil {
		return false
	}
	(places)[arc.FromID] = pl
	return true
}

func subToken(idx int, p Place) (Place, error) {
	if idx >= len(p.Tokens) || idx < 0 {
		return p, fmt.Errorf("idx out of bounds")
	}
	p.Tokens = append(p.Tokens[:idx], p.Tokens[idx+1:]...)
	return p, nil
}

func createTempPlaces(places []Place) []Place {
	var res []Place
	for _, p := range places {
		var tokensTemp = make([]string, len(p.Tokens))
		copy(tokensTemp, p.Tokens)
		res = append(res, Place{PlacesID: p.PlacesID, Tokens: tokensTemp})
	}
	return res
}
