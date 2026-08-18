package domain

import (
	"net/mail"
	"strings"
)

func isUpperLetter(r rune) bool { return r >= 'A' && r <= 'Z' }
func isDigit(r rune) bool       { return r >= '0' && r <= '9' }

func NormalizeInitial(initial string) (string, error) {
	trimmed := strings.TrimSpace(initial)
	initialRune := []rune(strings.ToUpper(trimmed))
	const lengthInitial = 6 //ex: WE25-1, NP25-1, MX25-1
	if len(initialRune) != lengthInitial {
		return "", ErrInvalidInitial
	}
	
	if !isUpperLetter(initialRune[0]) || !isUpperLetter(initialRune[1]) {
		return "", ErrInvalidInitial
	}
	if !isDigit(initialRune[2]) || !isDigit(initialRune[3]) {
		return "", ErrInvalidInitial
	}
	if initialRune[4] != '-' {
		return "", ErrInvalidInitial
	}
	if !isDigit(initialRune[5]) {
		return "", ErrInvalidInitial
	}

	return string(initialRune), nil
}

func NormalizeEmail(email string) (string, error) {
	trimmed := strings.TrimSpace(email)
	if trimmed == "" {
		return "", ErrInvalidEmail
	}

	addr, err := mail.ParseAddress(trimmed)
	if err != nil {
		return "", ErrInvalidEmail
	}

	at := strings.LastIndexByte(addr.Address, '@')
	if at < 0 || at == len(addr.Address)-1 {
		return "", ErrInvalidEmail
	}

	domainPart := addr.Address[at+1:]
	if !strings.Contains(domainPart, "gmail.com") && !strings.Contains(domainPart, "binus.ac.id") && !strings.Contains(domainPart, "binus.edu") {
		return "", ErrInvalidEmail
	}

	return strings.ToLower(addr.Address), nil
}