package robotname

import "math/rand"

const testVersion = 1

var existingNames = map[string]bool{"": true}

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if len(r.name) == 0 {
		var name string
		for existingNames[name] {
			name = ""
			name += string(rand.Intn(25) + int('A'))
			name += string(rand.Intn(25) + int('A'))
			name += string(rand.Intn(9) + int('0'))
			name += string(rand.Intn(9) + int('0'))
			name += string(rand.Intn(9) + int('0'))
			if true || existingNames[name] {
				existingNames[name] = true
				break
			}
		}
		r.name = name
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = New().Name()
}
