package card

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const (
	CardNoLength    int    = 16
	CardNoParts     int    = 4
	CardNoSeparator string = "-"

	CardCVVLength int = 3
)

var (
	ErrInvalidCardNoLength error = errors.New("invalid card number length")
	ErrInvalidCardNoValue  error = errors.New("invalid card number value")
	ErrInvalidCardCVV      error = errors.New("invalid card CVV")

	intRegexp = regexp.MustCompile(`^[0-9]$`)
)

type (
	CardNo  []string
	CardCVV string
)

func init() {
	if CardNoLength%CardNoParts != 0 {
		panic("FATAL: CardNoLength must be divisible by CardNoParts")
	}

	if CardNoSeparator == "" {
		panic("FATAL: CardNoSeparator must not be empty")
	}
}

func NewCardNo(rawCardNo string) (CardNo, error) {
	aCardNo := CardNo(strings.Split(rawCardNo, CardNoSeparator))
	return aCardNo, aCardNo.Validate()
}

func RndCardNo() CardNo {
	rndCardNoParts := make([]string, CardNoParts)
	for i := 0; i < CardNoParts; i++ {
		rndCardNoParts[i] = rndIntString(CardNoLength / CardNoParts)
	}

	cardNo, err := NewCardNo(strings.Join(rndCardNoParts, CardNoSeparator))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardNo
}

func (n CardNo) Validate() error {
	if len(n) != CardNoParts {
		return ErrInvalidCardNoLength
	}

	noDigits := 0
	for i := 0; i < CardNoParts; i++ {
		noDigits += len(n[i])
	}
	if noDigits != CardNoLength {
		return ErrInvalidCardNoLength
	}

	for p := 0; p < CardNoParts; p++ {
		for i := 0; i < CardNoLength/CardNoParts; i++ {
			if !intRegexp.MatchString(string(n[p][i])) {
				return ErrInvalidCardCVV
			}
		}
	}

	return nil
}

// Masks all but the last part of the card number
func (n CardNo) Mask() CardNo {
	maskedCardNo := make([]string, CardNoParts)
	for i := 0; i < CardNoParts-1; i++ {
		maskedCardNo[i] = ""
		for j := 0; j < CardNoLength/CardNoParts; j++ {
			maskedCardNo[i] = maskedCardNo[i] + "*"
		}
	}
	maskedCardNo[CardNoParts-1] = n[CardNoParts-1]

	return CardNo(maskedCardNo)
}

func (n CardNo) Prettify() string {
	return strings.Join(n, CardNoSeparator)
}

func NewCardCVV(cvv string) (CardCVV, error) {
	aCardCvv := CardCVV(cvv)
	return aCardCvv, aCardCvv.Validate()
}

func RndCardCVV() CardCVV {
	cardCVV, err := NewCardCVV(strconv.Itoa(int(100 + rand.Intn(900))))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardCVV
}

func (c CardCVV) Validate() error {
	if len(c) != CardCVVLength {
		return ErrInvalidCardCVV
	}

	for i := 0; i < CardCVVLength; i++ {
		if !intRegexp.MatchString(string(c[i])) {
			return ErrInvalidCardCVV
		}
	}

	return nil
}

// Masks the entire CVV
func (c CardCVV) Mask() CardCVV {
	maskedCVV := ""
	for i := 0; i < CardCVVLength; i++ {
		maskedCVV = fmt.Sprintf("%s*", maskedCVV)
	}
	return CardCVV(maskedCVV)
}

func (c CardCVV) String() string {
	return string(c)
}

func rndIntString(n int) string {
	a := ""
	for i := 0; i < n; i++ {
		a = fmt.Sprintf("%s%d", a, rand.Intn(10))
	}

	return a
}
