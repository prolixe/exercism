package protein

const testVersion = 1

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(input string) (proteins []string) {

	var pos int
	for {

		if pos+3 > len(input) {
			return
		}
		codon := FromCodon(input[pos : pos+3])
		if codon == "STOP" {
			return
		}
		proteins = append(proteins, codon)
		pos += 3

	}
}

func FromCodon(input string) string {
	return codonMap[input]
}
