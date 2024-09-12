package transactions

import (
	"main/core/db"
	"main/core/domain/transaction"
)

func ReadTransaction(id string, mask bool) (*transaction.Transaction, error) {
	t, err := db.SelectTransaction(db.GetDB(), id)
	if err != nil {
		return nil, err
	}

	if mask {
		t.MaskDetailsInPlace()
	}

	return t, nil
}
