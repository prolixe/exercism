package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	c := Clock{h: hour, m: minute}
	c = c.adjust()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) adjust() Clock {

	for c.m >= 60 {
		c.m -= 60
		c.h++
	}
	for c.m < 0 {
		c.m += 60
		c.h--
	}

	for c.h < 0 {
		c.h += 24
	}
	for c.h >= 24 {
		c.h -= 24
	}

	return c
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	c = c.adjust()
	return c
}
