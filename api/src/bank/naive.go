package bank

import (
	"main/src/transactions"
	"math/rand"
	"time"
)

type naiveBank struct{}

func NewNaiveBank() *naiveBank {
	return &naiveBank{}
}

func (b *naiveBank) SynchronousPayment(req *transactions.Transaction) error {
	req.SetState(transactions.Pending)

	// Add delay to mock request latency
	time.Sleep(10 * time.Millisecond)

	// Non-deterministic transaction outcome to simulate real-world conditions
	if rand.Float64() < 0.05 {
		req.SetState(transactions.Declined)
	}
	req.SetState(transactions.Completed)

	return nil
}
