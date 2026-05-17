package checker

import (
	"time"

	"github.com/Arya911/healthCli/internal/model"
)

var services = []string{
	"https://google.com",
	"https://github.com",
	"https://githuboljnholfvsndlopk.com",
}

func Start() chan model.Result {
	out := make(chan model.Result)

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		// defer ticker.Stop()

		for {
			runChecks(services, out)
			<-ticker.C
		}

	}()

	return out
}
