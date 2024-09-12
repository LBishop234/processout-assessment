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
	CardNoLength     int    = 16
	CardNoParts      int    = 4
	CardNoSeparator  string = "-"
	CardMinDigitIncl int    = 0
	CardMaxDigitIncl int    = 9

	CardCVMinIncl  int16 = 100
	CardCVVMaxIncl int16 = 999
)

func init() {
	if CardNoLength%CardNoParts != 0 {
		panic("FATAL: CardNoLength must be divisible by CardNoParts")
	}

	if CardNoSeparator == "" {
		panic("FATAL: CardNoSeparator must not be empty")
	}
}

var (
	ErrInvalidCardNoLength error = errors.New("invalid card number length")
	ErrInvalidCardNoValue  error = errors.New("invalid card number value")
	ErrInvalidCardCVV      error = errors.New("invalid card CVV")
)

type (
	CardCVV int16
)

type CardNo []string

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

func validateRawCardNo(rawCardNo string) error {
	regexpCardNo := regexp.MustCompile(`^\d{4}-\d{4}-\d{4}-\d{4}$`)
	if !regexpCardNo.MatchString(rawCardNo) {
		return ErrInvalidCardNoValue
	}
	return nil
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
			iInt, err := strconv.Atoi(string(n[p][i]))
			if err != nil {
				return err
			}

			if iInt < CardMinDigitIncl || iInt > CardMaxDigitIncl {
				return ErrInvalidCardNoValue
			}
		}
	}

	return nil
}

func (n CardNo) Mask() CardNo {
	maskedCardNo := make([]string, CardNoParts)
	for i := 0; i < CardNoParts-1; i++ {
		maskedCardNo[i] = ""
		for j := 0; j < CardNoLength/CardNoParts; j++ {
			maskedCardNo[i] = fmt.Sprintf("%s%s", maskedCardNo[i], "*")
		}
	}
	maskedCardNo[CardNoParts-1] = n[CardNoParts-1]

	return CardNo(maskedCardNo)
}

func (n CardNo) Prettify() string {
	return strings.Join(n, CardNoSeparator)
}

func NewCardCVV(cvv int16) (CardCVV, error) {
	aCardCvv := CardCVV(cvv)
	return aCardCvv, aCardCvv.Validate()
}

func RndCardCVV() CardCVV {
	cardCVV, err := NewCardCVV(int16(100 + rand.Intn(900)))
	if err != nil {
		// Panicking as this can only be the result of developer error
		panic(err)
	}

	return cardCVV
}

func (c CardCVV) Validate() error {
	if c < CardCVV(CardCVMinIncl) || c > CardCVV(CardCVVMaxIncl) {
		return ErrInvalidCardCVV
	}

	return nil
}

func (c CardCVV) Mask() CardCVV {
	return CardCVV(-999)
}

func rndIntString(n int) string {
	a := ""
	for i := 0; i < n; i++ {
		a = fmt.Sprintf("%s%d", a, rand.Intn(10))
	}

	return a
}
