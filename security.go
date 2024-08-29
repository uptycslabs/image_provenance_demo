package main

import (
	"fmt"
	"regexp"
	"unicode"
)

type SecurityAudit struct {
}

type User struct {
	Name    []rune
	Address []rune
	Email   string
}

func NewSecurityAudit() *SecurityAudit {
	return &SecurityAudit{}
}

func (s *SecurityAudit) Audit(u *User) error {
	if err := s.auditUserName(u); err != nil {
		return err
	}
	if err := s.auditAddress(u); err != nil {
		return err
	}
	if !s.auditEmail(u.Email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func (s *SecurityAudit) auditUserName(u *User) error {
	if unicode.IsLetter(u.Name[0]) {
		return nil
	}
	return fmt.Errorf("username does not start with letter")
}

func (s *SecurityAudit) auditAddress(u *User) error {
	if unicode.IsLetter(u.Name[0]) {
		return nil
	}
	return fmt.Errorf("username does not start with letter")
}

func (s *SecurityAudit) auditEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
