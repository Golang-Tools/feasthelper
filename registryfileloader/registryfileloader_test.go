package registryfileloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	path := "../registry.db"
	_, err := Load(path)
	if err != nil {
		assert.FailNow(t, err.Error(), "load get error")
	}
}
