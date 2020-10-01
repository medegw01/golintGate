package coding_test

import (
	java "golintGate/coding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Java(t *testing.T) {

	assert.Equal(t, "java", java.GetJava())
}
