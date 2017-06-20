package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(inputs []string) FreqMap {

	m := FreqMap{}
	cm := make(chan FreqMap, len(inputs))
	for _, input := range inputs {
		go func(s string) {
			cm <- Frequency(s)
		}(input)

	}

	for range inputs {
		for k, v := range <-cm {
			m[k] += v
		}
	}

	return m
}
