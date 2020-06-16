package errcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(123))
	assert.Panics(t, func() {
		assert.Nil(t, New(123))
	})
}
