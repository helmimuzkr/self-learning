package golang_validation

import (
	"errors"
	"net"
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrInvalidEmail   = errors.New("invalid email")
	ErrDomainNotFound = errors.New("domain email not found")
)

func isEmail(email string) error {
	// Check apakah formatnya email atau bukan
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	// Check apakah domain part dari sebuah email exist
	// local_part@domain_part
	dompart := strings.Split(email, "@")[1]
	_, err = net.LookupMX(dompart)
	if err != nil {
		return ErrDomainNotFound
	}

	return nil
}

func TestEmailValidation(t *testing.T) {
	testCase := []struct {
		name   string
		input  string
		expect error
	}{
		{"email is valid", "helmi@gmail.com", nil},
		{"not an email", "helmiATgmail.com", ErrInvalidEmail},
		{"domain part not found", "helmi@domain.notexist", ErrDomainNotFound},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			err := isEmail(test.input)
			assert.Equal(t, test.expect, err)
		})
	}
}
