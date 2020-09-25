package hello_test

import (
	"golintGate/hello"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result(t *testing.T) {

	assert.Equal(t, 123, hello.GetResult())
}
