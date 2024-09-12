package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMask(t *testing.T) {
	cardNo, err := NewCardNo("1234-5678-1234-5678")
	require.NoError(t, err)
	maskedCardNo := cardNo.Mask()
	assert.Equal(t, "****-****-****-5678", maskedCardNo.Prettify())
}
