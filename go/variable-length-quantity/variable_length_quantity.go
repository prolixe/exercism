package variablelengthquantity

import "errors"

const testVersion = 4

func DecodeVarint(input []byte) ([]uint32, error) {
	var i uint32
	decodedInts := make([]uint32, 0)
	for _, b := range input {
		i += uint32(b & 0x7F)
		if b&0x80 == 0 {
			// last byte for i
			decodedInts = append(decodedInts, i)
			i = 0
			continue
		}
		i <<= 7
	}
	if input[len(input)-1]&0x80 != 0 {
		return nil, errors.New("incomplete sequence")
	}
	return decodedInts, nil
}

func EncodeVarint(input []uint32) []byte {

	encoded := make([]byte, 0)
	for _, i := range input {
		bs := make([]byte, 0)
		bs = append(bs, byte(i%128))
		for i >>= 7; i != 0; i >>= 7 {
			bs = append(bs, 128+byte(i%128))
		}
		//reverse it.
		for j, k := 0, len(bs)-1; j < k; j, k = j+1, k-1 {
			bs[j], bs[k] = bs[k], bs[j]
		}
		encoded = append(encoded, bs...)
	}
	return encoded
}
