package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

func main() {
	configFile := flag.String("f", "", "Scenario file (.xml, .json, .yaml)")
	warnThreshold := flag.Float64("w", 0, "Warning threshold (sec)")
	critThreshold := flag.Float64("c", 0, "Critical threshold (sec)")
	timeout := flag.Duration("t", 30*time.Second, "Global timeout")
	insecure := flag.Bool("k", false, "Ignore SSL error")
	userAgent := flag.String("A", "Go-WebInject/1.0", "User-Agent")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("UNKNOWN - Argument -f <file> require")
		os.Exit(UNKNOWN)
	}

	tests, err := parseConfig(*configFile)
	if err != nil {
		fmt.Printf("UNKNOWN - Loading scenario file error: %v\n", err)
		os.Exit(UNKNOWN)
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout: *timeout,
		Jar:     jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
		},
	}

	start := time.Now()
	for i, tc := range tests {
		if err := runStep(client, tc, *userAgent); err != nil {
			elapsed := time.Since(start).Seconds()
			fmt.Printf("CRITICAL - Step %d [%s] failed: %v | time=%.3fs\n", i+1, tc.Name, err, elapsed)
			os.Exit(CRITICAL)
		}
	}

	duration := time.Since(start).Seconds()
	perfData := fmt.Sprintf("time=%.3fs;%.3f;%.3f;0;%.3f", duration, *warnThreshold, *critThreshold, timeout.Seconds())

	if *critThreshold > 0 && duration >= *critThreshold {
		fmt.Printf("CRITICAL - Too slow scenario: %.3fs | %s\n", duration, perfData)
		os.Exit(CRITICAL)
	}
	if *warnThreshold > 0 && duration >= *warnThreshold {
		fmt.Printf("WARNING - Slow motion scenario: %.3fs | %s\n", duration, perfData)
		os.Exit(WARNING)
	}

	fmt.Printf("OK - %d Valid steps %.3fs | %s\n", len(tests), duration, perfData)
	os.Exit(OK)
}
