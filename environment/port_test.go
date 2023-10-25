package environment_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/ncg/environment"
)

func TestPort(t *testing.T) {
	t.Run("Returns the port defined in the PORT environment variable.", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		defer os.Setenv("PORT", oldPort)

		os.Setenv("PORT", "4000")

		port, err := environment.Port(3000)
		assert.NoError(t, err)
		assert.Equal(t, 4000, port)
	})

	t.Run("Returns the default port if the PORT environment variable is not set.", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		defer os.Setenv("PORT", oldPort)

		os.Unsetenv("PORT")

		port, err := environment.Port(3000)
		assert.NoError(t, err)
		assert.Equal(t, 3000, port)
	})

	t.Run("Returns the default port if the PORT environment variable is empty.", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		defer os.Setenv("PORT", oldPort)

		os.Setenv("PORT", "")

		port, err := environment.Port(3000)
		assert.NoError(t, err)
		assert.Equal(t, 3000, port)
	})

	t.Run("Returns an error if the PORT environment variable is not a number.", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		defer os.Setenv("PORT", oldPort)

		os.Setenv("PORT", "foo")

		_, err := environment.Port(3000)
		assert.Error(t, err)
	})
}
