package forth

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const testVersion = 2

var ErrEmptyStack = fmt.Errorf("empty stack")
var ErrDivByZero = fmt.Errorf("division by zero")

const (
	Add  = "+"
	Sub  = "-"
	Mul  = "*"
	Div  = "/"
	Dup  = "DUP"
	Drop = "DROP"
	Swap = "SWAP"
	Over = "OVER"
)

var words = make(map[string][]string)

type Stack struct {
	elements []int
}

func (s *Stack) Push(e int) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e, true
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.elements)
}

var isInteger = regexp.MustCompile(`^[0-9]+$`)

func Forth(input []string) ([]int, error) {

	stack := Stack{elements: make([]int, 0)}

	commandList := make([]string, 0)
	addingUserDefWord := false
	if len(input) == 0 {
		return stack.elements, nil
	}

	for _, i := range strings.Split(strings.Join(input, " "), " ") {
		switch {
		case ";" == i:
			addingUserDefWord = false
			if err := addWord(commandList); err != nil {
				return nil, err
			}
			commandList = make([]string, 0)
		case addingUserDefWord:
			commandList = append(commandList, strings.ToUpper(i))
		case isInteger.MatchString(i):
			integer, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			stack.Push(integer)
		case isUserDefinedCommand(i):
			for _, op := range getCommands(i) {
				if isInteger.MatchString(string(op)) {
					integer, err := strconv.Atoi(op)
					if err != nil {
						return nil, err
					}
					stack.Push(integer)
					continue
				}
				err := evaluate(op, &stack)
				if err != nil {
					return []int{}, err
				}
			}
		case isBuiltInOp(i):
			err := evaluate(i, &stack)
			if err != nil {
				return []int{}, err
			}
		case ":" == i:
			addingUserDefWord = true
		default:
			return nil, fmt.Errorf("invalid input %s", i)
		}
	}
	return stack.elements, nil
}

func addWord(commands []string) error {

	if len(commands) < 2 {
		return fmt.Errorf("insuficient commands to define new op %v", commands)
	}
	if isInteger.MatchString(commands[0]) {
		return fmt.Errorf("can't redefine integers")
	}
	words[commands[0]] = commands[1:]
	return nil
}

func isUserDefinedCommand(command string) bool {
	_, ok := words[strings.ToUpper(command)]
	return ok
}
func getCommands(command string) []string {
	commands, _ := words[strings.ToUpper(command)]
	return commands
}

func isBuiltInOp(input string) bool {
	switch strings.ToUpper(input) {
	case "+", "-", "*", "/", "DUP", "DROP", "SWAP", "OVER":
		return true
	}
	return false
}

func evaluate(op string, stack *Stack) error {

	switch strings.ToUpper(op) {
	case Add:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		stack.Push(op1 + op2)
	case Sub:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		stack.Push(op2 - op1)
	case Mul:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		stack.Push(op1 * op2)
	case Div:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		if op1 == 0 {
			return ErrDivByZero
		}
		stack.Push(op2 / op1)
	case Dup:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		stack.Push(op1)
		stack.Push(op1)
	case Drop:
		_, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
	case Swap:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}

		stack.Push(op1)
		stack.Push(op2)
	case Over:
		op1, ok := stack.Pop()
		if !ok {
			return ErrEmptyStack
		}
		op2, ok := stack.Pop()

		if !ok {
			return ErrEmptyStack
		}
		stack.Push(op2)
		stack.Push(op1)
		stack.Push(op2)
	default:
		return fmt.Errorf("invalid operator %s", string(op))
	}
	return nil

}
