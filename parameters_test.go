package matomo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	Setup()
	// start with just the recommended parameters
	params := Parameters{
		RecommendedParameters: &RecommendedParameters{},
	}

	encoded := params.encode()
	assert.NotNil(t, encoded)
}
