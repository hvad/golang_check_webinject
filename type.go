package main

import "encoding/xml"

type TestCase struct {
	Name           string            `json:"name" xml:"name" yaml:"name"`
	Method         string            `json:"method" xml:"method" yaml:"method"`
	URL            string            `json:"url" xml:"url" yaml:"url"`
	Data           string            `json:"data" xml:"data" yaml:"data"`
	ExpectedStatus int               `json:"expected_status" xml:"expected_status" yaml:"expected_status"`
	LookFor        string            `json:"look_for" xml:"look_for" yaml:"look_for"`
	VariableName   string            `json:"var_name" xml:"var_name" yaml:"var_name"`
	VariableRegex  string            `json:"var_regex" xml:"var_regex" yaml:"var_regex"`
	Headers        map[string]string `json:"headers" xml:"headers" yaml:"headers"`
}

type XMLConfig struct {
	XMLName xml.Name   `xml:"testcases"`
	Tests   []TestCase `xml:"case"`
}
