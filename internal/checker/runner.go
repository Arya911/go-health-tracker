package checker

import (
	"context"
	"sync"
	"time"

	"github.com/Arya911/healthCli/internal/model"
)

func runChecks(urls []string, out chan<- model.Result) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			out <- checkUrl(ctx, u)

		}(url)
	}
	wg.Wait()
}
