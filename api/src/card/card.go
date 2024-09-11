package card

import (
	"errors"
	"math/rand"
)

const (
	CardNoLength     int  = 16
	CardCVVLength    int  = 3
	CardMinDigitIncl int8 = 0
	CardMaxDigitIncl int8 = 9
)

var (
	ErrInvalidCardNoLength = errors.New("invalid card number length")
	ErrInvalidCardNoDigit  = errors.New("invalid card number digit")
)

type (
	CardNo  []int8
	CardCVV []int8
)

func NewCardNo(cardNo []int8) (CardNo, error) {
	aCardNo := CardNo(cardNo)

	if err := aCardNo.Validate(); err != nil {
		return nil, err
	}

	return aCardNo, nil
}

func RndCardNo() CardNo {
	cardNo, err := NewCardNo(rndInt8Array(CardNoLength))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardNo
}

func (c CardNo) Validate() error {
	if len(c) != CardNoLength {
		return ErrInvalidCardNoLength
	}

	for i := 0; i < CardNoLength; i++ {
		if c[i] < CardMinDigitIncl || c[i] > CardMaxDigitIncl {
			return ErrInvalidCardNoDigit
		}
	}

	return nil
}

func NewCardCVV(cvv []int8) (CardCVV, error) {
	if len(cvv) != CardCVVLength {
		return nil, ErrInvalidCardNoLength
	}

	return CardCVV(cvv), nil
}

func RndCardCVV() CardCVV {
	cardCVV, err := NewCardCVV(rndInt8Array(CardCVVLength))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardCVV
}

func (c CardCVV) Validate() error {
	if len(c) != CardCVVLength {
		return ErrInvalidCardNoLength
	}

	for i := 0; i < CardCVVLength; i++ {
		if c[i] < CardMinDigitIncl || c[i] > CardMaxDigitIncl {
			return ErrInvalidCardNoDigit
		}
	}

	return nil
}

func rndInt8Array(n int) []int8 {
	a := make([]int8, n)
	for i := 0; i < n; i++ {
		a[i] = int8(rand.Intn(10))
	}

	return a
}
