package dna

import (
	"errors"
	"strings"
)

const testVersion = 2

type DNA string

type Histogram map[byte]int

func (dna DNA) Count(n byte) (int, error) {

	if !strings.ContainsAny(string(n), "GCTA") {
		return 0, errors.New("Invalid Nucleotide")
	}

	return strings.Count(string(dna), string(n)), nil

}

func (dna DNA) Counts() (Histogram, error) {

	G, _ := dna.Count('G')
	C, _ := dna.Count('C')
	A, _ := dna.Count('A')
	T, _ := dna.Count('T')

	if G+C+A+T != len(dna) {
		return nil, errors.New("Invalid strand")
	}
	return Histogram{'A': A, 'C': C, 'T': T, 'G': G}, nil
}
