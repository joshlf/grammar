package grammar

import (
	"math/rand"
)

type part struct {
	name       string
	isTerminal bool
	solns      [][]*part
}

func (p *part) init(name string) {
	p.name = name
	p.isTerminal = true
}

func (p *part) addSoln(soln []*part) {
	p.isTerminal = false
	p.solns = append(p.solns, soln)
}

func (p *part) mkSelf() string {
	if p.isTerminal {
		return p.name
	}

	soln := p.solns[rand.Intn(len(p.solns))]
	output := ""
	for _, child := range soln {
		output += child.mkSelf() + " "
	}

	return output[:len(output)-1]
}
