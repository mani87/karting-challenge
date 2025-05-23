package utils

import (
	"bufio"
	"compress/gzip"
	"os"
	"sync"
)

var (
	promoCache map[string]int
	once       sync.Once
)

func loadPromoCache() {
	files := []string{
		"../data/couponbase1.gz",
		"../data/couponbase2.gz",
		"../data/couponbase3.gz",
	}

	promoCache = make(map[string]int)
	wg := sync.WaitGroup{}
	fileResults := make(chan map[string]bool, len(files))

	// Load each file concurrently into a local map
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			localCodes := make(map[string]bool)

			f, err := os.Open(filePath)
			if err != nil {
				fileResults <- localCodes
				return
			}
			defer f.Close()

			gz, err := gzip.NewReader(f)
			if err != nil {
				fileResults <- localCodes
				return
			}
			defer gz.Close()

			scanner := bufio.NewScanner(gz)
			for scanner.Scan() {
				localCodes[scanner.Text()] = true
			}
			fileResults <- localCodes
		}(file)
	}

	wg.Wait()
	close(fileResults)

	for localSet := range fileResults {
		for code := range localSet {
			promoCache[code]++
		}
	}
}

func ValidatePromoCode(code string) bool {
	if len(code) < 8 || len(code) > 10 {
		return false
	}

	once.Do(loadPromoCache)
	return promoCache[code] >= 2
}
