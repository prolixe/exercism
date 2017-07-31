package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

const testVersion = 1

var plants = map[string]string{
	"V": "violets",
	"R": "radishes",
	"G": "grass",
	"C": "clover",
}

type Garden struct {
	plants   [][]string
	children []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	trimmedDiagram := strings.Trim(diagram, "\n")
	if trimmedDiagram == diagram {
		return nil, errors.New("expected starting newline")
	}
	rows := strings.Split(trimmedDiagram, "\n")
	// Check format
	if len(rows) != 2 || len(rows[0]) != len(rows[1]) || len(rows[0])%2 != 0 {
		return nil, errors.New("invalid diagram")
	}

	//Check cup code
	for _, code := range trimmedDiagram {
		if code == rune('\n') {
			continue
		}
		if !strings.ContainsRune(string(code), 'V') &&
			!strings.ContainsRune(string(code), 'R') &&
			!strings.ContainsRune(string(code), 'G') &&
			!strings.ContainsRune(string(code), 'C') {
			return nil, errors.New("invalid cup code")

		}
	}

	var g Garden
	g.children = append(g.children, children...)
	sort.Strings(g.children)

	for i := 0; i < len(g.children)-1; i++ {
		if g.children[i] == g.children[i+1] {
			return nil, errors.New("duplicate")
		}
	}

	g.plants = make([][]string, len(children))

	for _, r := range rows {
		for p := 0; p < len(r) || p/2 < len(children); p += 2 {
			g.plants[p/2] = append(g.plants[p/2], plants[string(r[p])], plants[string(r[p+1])])
		}
	}
	return &g, nil

}
func (g *Garden) Plants(child string) ([]string, bool) {

	pos, found := g.findChild(child)
	if found {
		return g.plants[pos], found
	}
	return nil, found
}

func (g *Garden) findChild(child string) (position int, found bool) {
	for i, c := range g.children {
		if c == child {
			position = i
			found = true
			return
		}

	}
	return
}
