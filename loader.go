package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func parseConfig(path string) ([]TestCase, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	var tests []TestCase

	switch ext {
	case ".json":
		err = json.Unmarshal(content, &tests)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, &tests)
	case ".xml":
		var xmlCfg XMLConfig
		err = xml.Unmarshal(content, &xmlCfg)
		tests = xmlCfg.Tests
	default:
		return nil, fmt.Errorf("format de fichier non supporté: %s", ext)
	}

	return tests, err
}
