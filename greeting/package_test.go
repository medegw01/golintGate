package greeting

import (
	"fmt"
	"os"
	"testing"
)

var coverage = 0.9

func TestMain(m *testing.M) {
	rc := m.Run()

	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < coverage {
			fmt.Printf("All Tests passed but coverage failed at: %f%%\n", (c * 100))
			rc = -1
		}
	}
	os.Exit(rc)
}
