package greeting_test

import (
	hello "golintGate/greeting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result(t *testing.T) {
	assert.Equal(t, 123, hello.GetResult())
	//assert.Equal(t, 1234, hello.GetHiResult())
}
