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
		"../data/couponbase1.gz",
		"../data/couponbase2.gz",
		"../data/couponbase3.gz",
	}

	type result struct {
		found bool
	}
	ch := make(chan result, len(files))

	for _, file := range files {
		go func(file string) {
			ch <- result{found: fileHasCode(file, code)}
		}(file)
	}

	count := 0
	for i := 0; i < len(files); i++ {
		if (<-ch).found {
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
