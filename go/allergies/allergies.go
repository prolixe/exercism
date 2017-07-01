package allergies

const testVersion = 1

type allergyType uint

const (
	eggs allergyType = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var allergyTypeToString = map[allergyType]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats"}

func Allergies(input uint) (allergies []string) {
	for allergy := range allergyTypeToString {
		if input&uint(allergy) != 0 {
			allergies = append(allergies, allergyTypeToString[allergy])
		}
	}
	return allergies
}

func AllergicTo(input uint, allergen string) bool {
	allergies := Allergies(input)
	for _, a := range allergies {
		if a == allergen {
			return true
		}
	}
	return false
}
