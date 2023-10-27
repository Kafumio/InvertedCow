package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	id := uint(555)
	token, err := GenerateToken(Claims{
		ID: id,
	})
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
	c, err := ParseToken(token)
	assert.Nil(t, err)
	assert.Equal(t, id, c.ID)
}
