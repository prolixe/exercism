package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

type teamTally struct {
	name           string
	mp, w, d, l, p int
}

func (t *teamTally) AddWin() {
	t.w++
	t.mp++
	t.p += 3
}

func (t *teamTally) AddLoss() {
	t.mp++
	t.l++
}

func (t *teamTally) AddDraw() {
	t.mp++
	t.d++
	t.p++
}

func teamTallyLess(teams []teamTally) func(i, j int) bool {
	return func(i, j int) bool {
		if teams[i].p == teams[j].p {
			return teams[i].name < teams[j].name
		}
		return teams[i].p > teams[j].p
	}
}

func Tally(r io.Reader, w io.Writer) error {

	var writeBuf bytes.Buffer
	var readBuf bytes.Buffer

	_, err := readBuf.ReadFrom(r)

	if err != nil {
		return err
	}

	input := readBuf.String()
	teams := make(map[string]*teamTally)
	// Parse the input
	for _, line := range strings.Split(input[1:], "\n") {
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		results := strings.Split(line, ";")
		if len(results) != 3 {
			return errors.New("invalid input format")
		}

		team1, team2, outcome := results[0], results[1], results[2]
		t1 := teams[team1]
		if t1 == nil {
			t1 = &teamTally{name: team1}
			teams[team1] = t1
		}
		t2 := teams[team2]
		if t2 == nil {
			t2 = &teamTally{name: team2}
			teams[team2] = t2
		}

		switch outcome {
		case "win":
			t1.AddWin()
			t2.AddLoss()
		case "loss":
			t1.AddLoss()
			t2.AddWin()
		case "draw":
			t1.AddDraw()
			t2.AddDraw()
		default:
			return errors.New("invalid outcome")

		}
	}
	// Put the team map into a slice and sort it
	teamsSlice := make([]teamTally, 0)
	for _, team := range teams {
		teamsSlice = append(teamsSlice, *team)
	}
	sort.Slice(teamsSlice, teamTallyLess(teamsSlice))

	// then add each element of the slice to the byte writer
	writeBuf.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range teamsSlice {
		writeBuf.WriteString(
			fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",
				team.name,
				team.mp,
				team.w,
				team.d,
				team.l,
				team.p))
	}
	_, err = w.Write(writeBuf.Bytes())

	if err != nil {
		return err
	}
	return nil
}
