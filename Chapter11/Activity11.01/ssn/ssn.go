package ssn

import (
	"errors"
)

var (
	ErrInvalidSSNLength  = errors.New("ssn is not nine characters long")
	ErrInvalidSSNNumbers = errors.New("ssn has non-numeric digits")
	ErrInvalidSSNPrefix  = errors.New("ssn has three zeros as a prefix")
	ErrInvalidDigitPlace = errors.New("ssn starts with a 9 requires 7 or 9 in the fourth place")
)

type SSN struct {
	digits [9]uint8
}

func NewSSN(s string) (*SSN, error) {
	var ssn SSN
	if err := checkLenght(s); err != nil {
		return nil, err
	}
	if d, err := convertToInts(s); err != nil {
		return nil, err
	} else {
		ssn.digits = *d
	}
	if err := ssn.checkPrefix(); err != nil {
		return nil, err
	}
	if err := ssn.checkDigitsInPlace(); err != nil {
		return nil, err
	}
	return &ssn, nil
}

func checkLenght(s string) error {
	if len(s) != 9 {
		return ErrInvalidSSNLength
	}
	return nil
}

func convertToInts(s string) (*[9]uint8, error) {
	var rsl [9]uint8
	for i, r := range s {
		if r < '0' || r > '9' {
			return nil, ErrInvalidSSNNumbers
		}
		rsl[i] = uint8(r - '0')
	}
	return &rsl, nil
}

func (s *SSN) checkPrefix() error {
	if s.digits[0] == 0 && s.digits[1] == 0 && s.digits[2] == 0 {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func (s *SSN) checkDigitsInPlace() error {
	if s.digits[0] == 9 && (s.digits[3] != 9 && s.digits[3] != 7) {
		return ErrInvalidDigitPlace
	}
	return nil
}
