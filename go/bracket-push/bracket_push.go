package brackets

const testVersion = 5

type enclosure int

const (
	invalid enclosure = iota
	parenthese
	brace
	bracket
)

func getEnclosure(r rune) enclosure {
	switch r {
	case '(', ')':
		return parenthese
	case '[', ']':
		return bracket
	case '{', '}':
		return brace
	}
	return invalid
}

func Bracket(input string) (bool, error) {

	stack := make([]enclosure, 0)

	for _, r := range input {
		switch r {
		case '(', '[', '{':
			//push
			stack = append(stack, getEnclosure(r))
		case ')', ']', '}':
			if len(stack) == 0 {
				return false, nil
			}
			if stack[len(stack)-1] != getEnclosure(r) {
				return false, nil
			}
			//pop
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0, nil
}
