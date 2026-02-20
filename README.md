# Golang WebInject (Nagios Plugin)

A modern, fast, and dependency-free replacement for `check_webinject`. 
This plugin allows you to execute multi-step HTTP(S) test scenarios with session management.

## Features

* **Supported Formats**: XML (WebInject compatible), JSON, and YAML.
* **Sessions**: Automatic cookie handling between test steps.
* **Variables**: Capture values via Regex to reuse them in subsequent steps.
* **HTTPS**: Full TLS support with a `-k` bypass option for insecure certificates.
* **Nagios Compliant**: Standard exit codes and PerfData for graphing.

## Installation

```bash
go mod init golang_check_webinject
go mod tidy
go build -o golang_check_webinject
```
