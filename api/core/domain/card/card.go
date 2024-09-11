package card

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	CardNoLength     int = 16
	CardCVVLength    int = 3
	CardMinDigitIncl int = 0
	CardMaxDigitIncl int = 9
)

var (
	ErrInvalidCardNoLength error = errors.New("invalid card number length")
	ErrInvalidCardNoDigit  error = errors.New("invalid card number digit")
)

type (
	CardNo  string
	CardCVV string
)

func NewCardNo(cardNo string) (CardNo, error) {
	aCardNo := CardNo(cardNo)
	return aCardNo, aCardNo.Validate()
}

func RndCardNo() CardNo {
	cardNo, err := NewCardNo(rndIntString(CardNoLength))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardNo
}

func (n CardNo) Validate() error {
	if len(n) != CardNoLength {
		return ErrInvalidCardNoLength
	}

	for i := 0; i < CardNoLength; i++ {
		iInt, err := strconv.Atoi(string(n[i]))
		if err != nil {
			return err
		}

		if iInt < CardMinDigitIncl || iInt > CardMaxDigitIncl {
			return ErrInvalidCardNoDigit
		}
	}

	return nil
}

func NewCardCVV(cvv string) (CardCVV, error) {
	aCardCvv := CardCVV(cvv)
	return aCardCvv, aCardCvv.Validate()
}

func RndCardCVV() CardCVV {
	cardCVV, err := NewCardCVV(rndIntString(CardCVVLength))
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
		iInt, err := strconv.Atoi(string(c[i]))
		if err != nil {
			return err
		}

		if iInt < CardMinDigitIncl || iInt > CardMaxDigitIncl {
			return ErrInvalidCardNoDigit
		}
	}

	return nil
}

func rndIntString(n int) string {
	a := ""
	for i := 0; i < n; i++ {
		a = fmt.Sprintf("%s%d", a, rand.Intn(10))
	}

	return a
}
