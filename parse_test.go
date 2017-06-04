package date_interval

import (
	"fmt"
	"testing"
)

func TestGoodParse(t *testing.T) {

	// Test good intervals
	sample := []string{
		"23 Days ",
		"3 mo",
		"4y",
		"2d 1mo 4yr",
		"2mo -3d",
		"2years 1month",
		"2y 0months 1DAY",
	}

	for i := range sample {
		y, m, d, err := Parse(sample[i])
		if err != nil {
			t.Errorf("Valid interval, but got: %v", sample[i], err)
			continue
		}
		fmt.Printf("Got expected result: %q Converts to %d y, %d m, %d d\n", sample[i], y, m, d)
	}
}

func TestBadParse(t *testing.T) {

	// Test bad intervals
	sample := []string{
		"2y 3x",
		"2years month 1",
		"1d 0mon ",
		"3d 5d ",
	}

	for i := range sample {
		y, m, d, err := Parse(sample[i])
		if err == nil {
			t.Errorf("Bad interval %q, but got result: Converts to %d y, %d m, %d d", sample[i], y, m, d)
			continue
		}
		fmt.Printf("Got expected error: %v\n", err)
	}
}
