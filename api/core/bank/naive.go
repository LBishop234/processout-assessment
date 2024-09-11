package bank

import (
	"main/core/domain/transaction"
	"math/rand"
	"time"
)

type naiveBank struct{}

func NewNaiveBank() *naiveBank {
	return &naiveBank{}
}

func (b *naiveBank) SynchronousPayment(req *transaction.Transaction) error {
	req.SetState(transaction.Pending)

	// Add delay to mock request latency
	time.Sleep(10 * time.Millisecond)

	// Non-deterministic transaction outcome to simulate real-world conditions
	if rand.Float64() < 0.05 {
		req.SetState(transaction.Declined)
	}
	req.SetState(transaction.Completed)

	return nil
}
