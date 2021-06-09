package matomo

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	Setup()
	// start with just the recommended parameters
	params := Parameters{
		RecommendedParameters: &RecommendedParameters{
			ActionName: stringPtr("test_action"),
			Rand:       int64Ptr(rand.Int63n(9999999999999999)),
		},
	}

	encoded := params.encode()
	assert.NotNil(t, encoded)
	assert.NotNil(t, encoded["rand"])
	assert.NotZero(t, encoded["rand"])
	assert.Equal(t, *params.RecommendedParameters.ActionName, encoded["action_name"])
	assert.Equal(t, "1", encoded["apiv"])
}
