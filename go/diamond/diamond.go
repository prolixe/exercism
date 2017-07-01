package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

func Gen(b byte) (string, error) {
	if b < 'A' || 'Z' < b {
		return "", errors.New("Invalid")
	}

	var d []string
	diamondLength := int(b-'A')*2 + 1
	for charRow := byte('A'); charRow <= b; charRow += byte(1) {
		outerSpace := int(b - charRow)
		row := strings.Repeat(" ", outerSpace) + string(charRow)
		if charRow != byte('A') {
			innerSpace := diamondLength - outerSpace*2 - 2
			row += strings.Repeat(" ", innerSpace) + string(charRow)
		}
		row += strings.Repeat(" ", outerSpace)
		d = append(d, row)
	}
	// Copy the top of the diamond into the bottom.
	if len(d)-2 >= 0 {
		for i := len(d) - 2; i >= 0; i-- {
			d = append(d, d[i])
		}
	}

	//fmt.Println(strings.Replace(strings.Join(d, "\n"), " ", ".", -1))
	return strings.Join(d, "\n") + "\n", nil
}
