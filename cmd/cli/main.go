package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/vitorlfaria/go-observer/internal/client"
	"github.com/vitorlfaria/go-observer/internal/config"
	"github.com/vitorlfaria/go-observer/internal/types"
)

func main() {
	configArg := flag.String("config", "", "Config file containing the URLs to observe")

	flag.Parse()

	if *configArg == "" {
		log.Fatal("Error: config is required")
	}

	configFileBytes, err := os.ReadFile(*configArg)

	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err.Error())
	}

	configFile, err := config.Parse(configFileBytes)

	if err != nil {
		log.Fatalf("Error parsing JSON: %s\n", err.Error())
	}

	client := client.NewClient()
	var wg sync.WaitGroup
	resultsChan := make(chan types.Result, len(configFile.URLs))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for _, url := range configFile.URLs {
		wg.Go(func() {
			result := types.Result{
				URL:      url,
				Status:   "",
				Duration: 0,
				Error:    nil,
			}

			start := time.Now()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				log.Printf("Error creating request to %s: %s", url, err.Error())
				result.Error = err
				resultsChan <- result
				return
			}

			callResult, err := client.Do(req)
			duration := time.Since(start)
			result.Duration = duration

			if err != nil {
				log.Printf("Error on request to %s: %s", url, err.Error())
				result.Error = err
				resultsChan <- result
				return
			}

			result.Status = callResult.Status
			resultsChan <- result
			callResult.Body.Close()
		})
	}

	wg.Wait()
	close(resultsChan)

	fmt.Println("--- Observability Report ---")
	for result := range resultsChan {
		fmt.Printf("Url: %s\n", result.URL)
		fmt.Printf("Status: %s\n", result.Status)
		fmt.Printf("Duration: %v\n", result.Duration)
		fmt.Printf("Error: %v\n", result.Error)
		fmt.Println("----- // -----")
	}
}
