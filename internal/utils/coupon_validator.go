package utils

import (
	"bufio"
	"compress/gzip"
	"os"
)

func ValidatePromoCode(code string) bool {
	if len(code) < 8 || len(code) > 10 {
		return false
	}

	files := []string{
		"data/couponbase1.gz",
		"data/couponbase2.gz",
		"data/couponbase3.gz",
	}

	count := 0
	for _, file := range files {
		if fileHasCode(file, code) {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}

func fileHasCode(filePath string, code string) bool {
	f, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return false
	}
	defer gz.Close()

	scanner := bufio.NewScanner(gz)
	for scanner.Scan() {
		if scanner.Text() == code {
			return true
		}
	}
	return false
}
