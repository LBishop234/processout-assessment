package card

import "errors"

var (
	ErrInvalidExpiryMonth = errors.New("invalid expiry month")
	ErrInvalidExpiryYear  = errors.New("invalid expiry year")
)

type CardExpiry struct {
	Month int8 `json:"month"`
	Year  int  `json:"year"`
}

func NewCardExpiry(month int8, year int) (CardExpiry, error) {
	aCardExpiry := CardExpiry{
		Month: month,
		Year:  year,
	}

	return aCardExpiry, aCardExpiry.Validate()
}

func (c *CardExpiry) Validate() error {
	if c.Month < 1 || c.Month > 12 {
		return ErrInvalidExpiryMonth
	}

	if c.Year < 0 {
		return ErrInvalidExpiryYear
	}

	return nil
}
