package utils

import (
	"bufio"
	"compress/gzip"
	"os"
	"path/filepath"
)

func ValidatePromoCode(code string) bool {
	if len(code) < 8 || len(code) > 10 {
		return false
	}

	files := []string{
		getAbsolutePath("internal/data/couponbase1.gz"),
		getAbsolutePath("internal/data/couponbase2.gz"),
		getAbsolutePath("internal/data/couponbase3.gz"),
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

func getAbsolutePath(relPath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(cwd, relPath)
}
