package domain

import "strings"

type EmailDomainPolicy struct {
	AllowedDomains map[string]bool
}

func NewEmailDomainPolicy(domains ...string) EmailDomainPolicy {
	m := make(map[string]bool, len(domains))

	for _, domain := range domains {
		m[domain] = true
	}

	return EmailDomainPolicy{
		AllowedDomains: m,
	}
}

func (policy *EmailDomainPolicy) Check(email string) error {
	at := strings.LastIndexByte(email, '@')
	if at < 0 {
		return ErrEmailDomainNotAllowed
	}
	domainPart := strings.ToLower(email[at+1:])
	if !policy.AllowedDomains[domainPart] {
		return ErrEmailDomainNotAllowed
	}
	return nil
}
