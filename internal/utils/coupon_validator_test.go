package utils

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestValidatePromoCodeValid(t *testing.T) {
	stepUpDir()
	valid := ValidatePromoCode("HAPPYHRS")
	require.True(t, valid)
}

func TestValidatePromoCodeInvalid(t *testing.T) {
	stepUpDir()
	valid := ValidatePromoCode("SUPER100")
	require.False(t, valid)
}

func stepUpDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	parent := filepath.Dir(cwd)
	grandParent := filepath.Dir(parent)

	err = os.Chdir(grandParent)
	if err != nil {
		panic(err)
	}
}
