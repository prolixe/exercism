package railfence

// nextPosSeq generate the sequence of increments in the position of a
// character depending on its starting position (its level) and the size of the
// rails.
// For the edges, the increment is always the same: (rails-1)*2 - startPos*2 or startPos*2
// whatever the one that isn't 0.
// And every other sequence is a sequence of those 2 increment, alternated.
func nextPosSeq(startPos, rails int) func() int {
	flipflop := true
	return func() int {
		flipflop = !flipflop
		if (flipflop || startPos == (rails-1)) && startPos != 0 {
			// Edges cases
			return startPos * 2
		}
		return ((rails - 1) - startPos) * 2
	}
}

func Encode(message string, rails int) string {
	output := ""
	for startPos := 0; startPos < rails; startPos++ {
		nextPos := nextPosSeq(startPos, rails)
		for pos := startPos; pos < len(message); pos += nextPos() {
			output += message[pos : pos+1]
		}
	}
	return output
}

func Decode(message string, rails int) string {
	output := make([]byte, len(message))
	messagePos := 0
	for startPos := 0; startPos < rails; startPos++ {
		nextPos := nextPosSeq(startPos, rails)
		for pos := startPos; pos < len(message); pos += nextPos() {
			output[pos] = message[messagePos]
			messagePos++
		}
	}
	return string(output)
}
