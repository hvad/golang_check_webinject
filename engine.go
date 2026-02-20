package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var sessionVars = make(map[string]string)

func runStep(client *http.Client, tc TestCase, defaultUA string) error {
	url := injectVars(tc.URL)
	data := injectVars(tc.Data)

	req, err := http.NewRequest(tc.Method, url, strings.NewReader(data))
	if err != nil {
		return err
	}

	// Headers configuration
	req.Header.Set("User-Agent", defaultUA)
	if tc.Method == "POST" && tc.Data != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for k, v := range tc.Headers {
		req.Header.Set(k, injectVars(v))
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	// Validate Status Code
	if tc.ExpectedStatus != 0 && resp.StatusCode != tc.ExpectedStatus {
		return fmt.Errorf("HTTP %d (attendu %d)", resp.StatusCode, tc.ExpectedStatus)
	}

	// Validate content (Regex)
	if tc.LookFor != "" {
		matched, _ := regexp.MatchString(tc.LookFor, body)
		if !matched {
			return fmt.Errorf("pattern '%s' non trouvé", tc.LookFor)
		}
	}

	// Capturing variables for use in following steps.
	if tc.VariableName != "" && tc.VariableRegex != "" {
		re := regexp.MustCompile(tc.VariableRegex)
		matches := re.FindStringSubmatch(body)
		if len(matches) > 1 {
			sessionVars[tc.VariableName] = matches[1]
		} else {
			return fmt.Errorf("échec de capture variable '%s'", tc.VariableName)
		}
	}

	return nil
}

func injectVars(input string) string {
	for k, v := range sessionVars {
		input = strings.ReplaceAll(input, "{"+k+"}", v)
	}
	return input
}
