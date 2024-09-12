package transaction

type TransactionStatus struct {
	ID    string           `json:"id"`
	State TransactionState `json:"state"`
}

func NewTransactionStatus(id string, state TransactionState) *TransactionStatus {
	return &TransactionStatus{
		ID:    id,
		State: state,
	}
}
