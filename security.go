package main

import (
	"fmt"
	"unicode"
)

type SecurityAudit struct {
}

type User struct {
	Name    []rune
	Address []rune
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
