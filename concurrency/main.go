package concurrency

import (
	"log"
	"time"
)

func track(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, execution time %s\n", name, time.Since(start))
	}
}

func CheckWebsitesConcurrent(wc WebsiteChecker, urls []string) map[string]bool {
	defer track("check website")()
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	return results
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	defer track("check website")()
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
