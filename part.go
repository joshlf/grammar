package grammar

import (
	"io"
	"math/rand"
)

type part struct {
	name       []byte
	isTerminal bool
	solns      [][]*part
}

func (p *part) init(name string) {
	p.name = []byte(name)
	p.isTerminal = true
}

func (p *part) addSoln(soln []*part) {
	p.isTerminal = false
	p.solns = append(p.solns, soln)
}

func (p *part) speak(w io.Writer) error {
	if p.isTerminal {
		_, err := w.Write(p.name)
		return err
	}

	soln := p.solns[rand.Intn(len(p.solns))]
	writeSpace := false
	for _, child := range soln {
		if writeSpace {
			w.Write([]byte{' '})
		} else {
			writeSpace = true
		}
		err := child.speak(w)
		if err != nil {
			return err
		}
	}
	return nil
}
