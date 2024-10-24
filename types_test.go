package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount(
		"Arman",
		"Sarvardin",
		"arman123",
	)
	assert.Nil(t, err)

	fmt.Printf("\n%+v", acc)
}