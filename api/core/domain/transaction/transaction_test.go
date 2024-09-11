package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionMaskDetails(t *testing.T) {
	aTransaction := RndTransaction()

	originalCardNo := aTransaction.CardNo
	originalCVV := aTransaction.CVV

	aTransaction.MaskDetails()
	assert.NotEqual(t, aTransaction.CardNo, originalCardNo)
	assert.NotEqual(t, aTransaction.CVV, originalCVV)
}
