package secret

const testVersion = 2

func Handshake(code uint) []string {
	a := make([]string, 0)

	if code&1 == 1 {
		a = append(a, "wink")
	}

	if code&(1<<1) == 1<<1 {
		a = append(a, "double blink")
	}

	if code&(1<<2) == 1<<2 {
		a = append(a, "close your eyes")
	}

	if code&(1<<3) == 1<<3 {
		a = append(a, "jump")
	}

	if code&(1<<4) == 1<<4 {
		for i := 0; i < len(a)/2; i++ {
			a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
		}
	}

	return a
}
