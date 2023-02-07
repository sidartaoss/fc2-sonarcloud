package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// arrange
	expected := 30
	// act
	actual := sum(10, 20)
	// assert
	assert.Equal(t, expected, actual)
}
