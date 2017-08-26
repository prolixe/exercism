package variablelengthquantity

import "errors"

const testVersion = 4

// DecodeVarint decode a slice of byte into a slice of uint32 integers
func DecodeVarint(input []byte) ([]uint32, error) {
	var d uint32
	var complete bool
	decodedInts := make([]uint32, 0)
	for _, b := range input {
		d += uint32(b & 0x7F)
		complete = (b&0x80 == 0)
		if complete {
			decodedInts = append(decodedInts, d)
			d = 0
			continue
		}
		d <<= 7
	}
	if !complete {
		return nil, errors.New("incomplete sequence")
	}
	return decodedInts, nil
}

// EncodeVarint encode a slice of uint32 integers into a slice of byte
func EncodeVarint(input []uint32) []byte {

	encoded := make([]byte, 0)
	for _, i := range input {
		e := []byte{byte(i % 128)}
		for i >>= 7; i != 0; i >>= 7 {
			e = append([]byte{128 + byte(i%128)}, e...)
		}
		encoded = append(encoded, e...)
	}
	return encoded
}
