package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatePromoCodeValid(t *testing.T) {
	valid := ValidatePromoCode("HAPPYHRS")
	require.True(t, valid)
}

func TestValidatePromoCodeInvalid(t *testing.T) {
	valid := ValidatePromoCode("SUPER100")
	require.False(t, valid)
}
