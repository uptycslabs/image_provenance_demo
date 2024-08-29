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

func TestNegativeUserName(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Name: []rune("1John"),
	}
	assert.Error(t, s.auditUserName(u), "username does not start with letter")
}

func TestNegativeAddress(t *testing.T) {
	s := NewSecurityAudit()
	u := &User{
		Name: []rune("1-John-Street"),
	}
	assert.Error(t, s.auditAddress(u), "address does not start with letter")
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
