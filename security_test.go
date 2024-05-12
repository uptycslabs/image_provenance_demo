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
