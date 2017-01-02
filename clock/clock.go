// Package clock provide a clock that handles times without dates
package clock

import (
	"fmt"
)

const testVersion = 4

// Clock struct with hour and min as int
type Clock struct{ h, m int }

// New : constructor
func New(hour, minute int) Clock {
	var c Clock

	//add minutes converted to hours
	c.h += minute / 60

	//add hour
	c.h += hour

	//remainder of minutes
	c.m = minute % 60

	normalizeclock(&c.h, &c.m)

	return c
}

// String : Return formatted string time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add : add or subtract minutes to clock
func (c Clock) Add(minutes int) Clock {

	// no adjustments if no minutes
	if minutes == 0 {
		return c
	}

	//add existing and new minutes
	c.m += minutes

	//convert minutes passed to hours
	hour := c.m / 60

	// adjust hour/min of clock accordingly
	if hour != 0 {
		c.h += hour
		c.m = c.m % 60
	}

	normalizeclock(&c.h, &c.m)

	return c
}

func normalizeclock(h, m *int) {

	//roll over hour
	*h %= 24

	//adjust negative minutes
	if *m < 0 {
		*h--
		*m += 60
	}

	//negative hours to be minused from 24
	if *h < 0 {
		*h += 24
	}

	return
}
