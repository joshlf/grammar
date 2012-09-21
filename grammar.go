// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The grammar package randomly generates sentences based on a user-supplied context-free grammar.
package grammar

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/joshlf13/errlist"
	"io"
	"math/rand"
	"strings"
	"time"
)

// A Grammar object contains the definition for a single
// context-free grammar, and can randomly generate
// sentences based upon that grammar structure.
type Grammar struct {
	initialized bool
	head        *part
	parts       map[string]*part
}

// Returns a Grammar object initialized with the
// grammar definition read from rdr. Additionaly,
// an error is returned if the definition is
// ill-formatted or unavailable for reading.
// In this case, a nil *Grammar is returned
func New(rdr io.Reader) (*Grammar, error) {
	g := new(Grammar)
	err := g.init(rdr)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Generate a random sentence based upon the
// grammar structure of g, and write it to w
func (g *Grammar) Speak(w io.Writer) error {
	if g.initialized {
		return g.head.speak(w)
	}
	return errors.New("Uninitialized Grammar")
}

func (g *Grammar) init(rdr io.Reader) error {
	r := bufio.NewReader(rdr)

	// Handle first line
	for {
		b, ip, err := r.ReadLine()
		if err != nil {
			return fmt.Errorf("Error reading input: %v", err)
		} else if ip {
			for ip {
				var _b []byte
				_b, ip, err = r.ReadLine()
				if err != nil {
					return fmt.Errorf("Error reading input: %v", err)
				}
				b = append(b, _b...)
			}
		}

		strs := strings.Fields(string(b))
		length := len(strs)
		if length < 1 {
			continue
		}

		g.head = new(part)
		g.parts = make(map[string]*part)

		soln := make([]*part, length)
		for i, s := range strs {
			p, ok := g.parts[s]
			if !ok {
				g.parts[s] = new(part)
				p = g.parts[s]
				p.init(s)
			}
			soln[i] = p
		}
		g.head.addSoln(soln)
		break
	}

	var errs *errlist.Errlist

	// Handle second line on
	for {
		b, ip, err := r.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return fmt.Errorf("Error reading input: %v", err)
			}
		} else if ip {
			for ip {
				var _b []byte
				_b, ip, err = r.ReadLine()
				if err != nil {
					return fmt.Errorf("Error reading input: %v", err)
				}
				b = append(b, _b...)
			}
		}

		strs := strings.Fields(string(b))
		length := len(strs)
		if length < 2 {
			continue
		}

		err = g.addRule(strs[0], strs[1:])
		if err != nil {
			if errs == nil {
				errs = errlist.NewString("Error adding definition: " + err.Error())
			} else {
				errs.AddString("Error adding definition: " + err.Error())
			}
		}
	}
	g.initialized = true

	// Seems redundant, but avoids the issue that
	// an interface with a concrete type and nil
	// value doesn't evaluate to being equal to nil
	if errs == nil {
		return nil
	}
	return errs
}

func (g *Grammar) addRule(prt string, definition []string) error {
	head, ok := g.parts[prt]
	if !ok {
		return fmt.Errorf("Unknown part: %v", prt)
	}

	soln := make([]*part, len(definition))
	for i, d := range definition {
		p, ok := g.parts[d]
		if !ok {
			g.parts[d] = new(part)
			p = g.parts[d]
			p.init(d)
		}
		soln[i] = p
	}

	head.addSoln(soln)
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
