package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

const testVersion = 2

var replacer = strings.NewReplacer(" ", "", "-", "", ".", "", "(", "", ")", "", "+", "")

func Number(input string) (string, error) {
	strippedInput := replacer.Replace(input)
	if strippedInput[0:1] == "1" {
		strippedInput = strippedInput[1:]
	}
	if len(strippedInput) != 10 {
		return "", errors.New("Invalid number!")
	}
	if strippedInput[3] == '0' || strippedInput[3] == '1' {
		return "", errors.New("Invalid if exchange code does not start with 2-9")
	}
	return strippedInput, nil

}

func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
