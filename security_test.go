package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuditUserName(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Name: []rune("John"),
	}
	assert.NoError(t, s.auditUserName(u))

}

func TestAuditAddress(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Name: []rune("John"),
	}
	assert.NoError(t, s.auditAddress(u))

}

func TestAuditEmail(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Email: "abc@abc.com",
	}
	f := s.auditEmail(u.Email)
	assert.Equal(t, f, true)
}

func TestNegativeAuditEmail(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Email: "abc.com",
	}
	f := s.auditEmail(u.Email)
	assert.Equal(t, f, false)
}
