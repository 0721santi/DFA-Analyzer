package dfaanalyzer

import (
	"errors"
	"fmt"
	"strings"
)

type State struct {
	id           string
	isFinalState bool
}

type DFA struct {
	Alphabet        []string
	Qs              []State
	Q0              State //Start State
	FinalStates     []State
	TransitionFuncs map[State]map[string]State
}

func GenDFA() *DFA {
	return &DFA{
		Alphabet:        make([]string, 0),
		Qs:              make([]State, 0),
		Q0:              State{id: "NA", isFinalState: false},
		FinalStates:     make([]State, 0),
		TransitionFuncs: make(map[State]map[string]State),
	}
}
func (dfa *DFA) AddSymbol(s string) error {
	exist := false
	for _, symbol := range dfa.Alphabet {
		if symbol == s {
			exist = true
			break
		}
	}
	if exist {
		return errors.New("ya existe el simbolo")
	}
	dfa.Alphabet = append(dfa.Alphabet, s)
	return nil
}

func (dfa *DFA) AddState(id string, iFS bool) (*State, error) {
	exist := false
	for _, val := range dfa.Qs {
		if val.id == id {
			exist = true
			break
		}
	}
	if exist {
		return nil, errors.New("ya existe el estado")
	}
	s := State{id: id, isFinalState: iFS}
	dfa.Qs = append(dfa.Qs, s)
	if iFS {
		dfa.FinalStates = append(dfa.FinalStates, s)
	}
	return &s, nil
}

func (dfa *DFA) SetStartState(id string) error {
	exist := false
	for _, val := range dfa.Qs {
		if val.id == id {
			dfa.Q0 = val
			exist = true
			break
		}
	}
	if !exist {
		return errors.New("no se encontro el estado")
	}
	return nil
}

func (dfa *DFA) AddTransitionFunc(qI, qF State, values []string) string {
	added := 0
	for _, value := range values {
		for _, symbol := range dfa.Alphabet {
			if value == symbol {
				added++
				dfa.TransitionFuncs[qI][value] = qF
				break
			}
		}
	}
	return fmt.Sprintf("Se agregaron %d funciones.", added)
}

func (dfa DFA) delta(qI State, val string) (State, bool) {
	qN, ok := dfa.TransitionFuncs[qI][val]
	return qN, ok
}

func (dfa DFA) AnalizeChain(str string) bool {
	if str == "E" {
		return false
	}
	actualState := dfa.Q0
	vals := strings.Split(str, "")
	for i := 0; i < len(vals); i++ {
		qN, existNext := dfa.delta(actualState, vals[i])
		if !existNext {
			return false
		}
		actualState = qN
	}
	if actualState.isFinalState {
		return true
	} else {
		return false
	}
}
