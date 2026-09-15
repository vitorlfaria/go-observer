package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/vitorlfaria/go-observer/internal/client"
	"github.com/vitorlfaria/go-observer/internal/types"
)

func main() {
	configArg := flag.String("config", "", "Config file containing the URLs to observe")

	flag.Parse()

	if *configArg == "" {
		log.Fatal("Error: config is required")
	}

	var configFile types.Config

	configFileBytes, err := os.ReadFile(*configArg)

	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err.Error())
	}

	if err := json.Unmarshal(configFileBytes, &configFile); err != nil {
		log.Fatalf("Error parsing JSON: %s\n", err.Error())
	}

	client := client.NewClient()
	var wg sync.WaitGroup
	resultsChan := make(chan types.Result, len(configFile.URLs))
	for _, url := range configFile.URLs {
		wg.Go(func() {
			start := time.Now()
			callResult, err := client.Get(url)
			duration := time.Since(start)

			result := types.Result{
				URL:      url,
				Status:   "",
				Duration: duration,
				Error:    nil,
			}

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
