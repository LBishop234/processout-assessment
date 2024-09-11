package bank

import (
	"main/domain"
	"math/rand"
	"time"
)

type naiveBank struct{}

func NewNaiveBank() *naiveBank {
	return &naiveBank{}
}

func (b *naiveBank) SynchronousPayment(req *domain.Transaction) error {
	req.SetState(domain.Pending)

	// Add delay to mock request latency
	time.Sleep(10 * time.Millisecond)

	if rand.Float64() < 0.05 {
		req.SetState(domain.Declined)
	}
	req.SetState(domain.Completed)

	return nil
}
