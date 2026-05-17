package checker

import (
	"context"
	"net/http"
	"time"

	"github.com/Arya911/healthCli/internal/model"
)

func checkUrl(ctx context.Context, url string) model.Result {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)

	latency := time.Since(start)

	if err != nil {
		return model.Result{
			URL:     url,
			Status:  "DOWN",
			Latency: latency,
			Error:   err.Error(),
			Checked: time.Now(),
		}
	}

	defer resp.Body.Close()

	status := "UP"
	if resp.StatusCode >= 400 {
		status = "DOWNGRADE"
	}

	return model.Result{
		URL:     url,
		Status:  status,
		Latency: latency,
		Checked: time.Now(),
	}

}
