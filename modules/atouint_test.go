package modules_test

import (
	"testing"

	"echo-gorm-example/modules"

	"github.com/stretchr/testify/assert"
)

func TestAtouint(t *testing.T) {
	assert := assert.New(t)

	t.Run("Positive number", func(t *testing.T) {
		u, err := modules.Atouint("123")
		assert.Equal(u, uint(123))
		assert.NoError(err)
	})

	t.Run("Negative number", func(t *testing.T) {
		u, err := modules.Atouint("-123")
		assert.Empty(u)
		assert.Error(err)
	})
}
