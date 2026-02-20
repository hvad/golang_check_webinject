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

## Usage

This plugin is designed to execute web scenarios and return results compatible 
with monitoring systems like **Nagios**, **Icinga**, or **Centreon**.

### Running the Plugin

Once you have the binary for your architecture, you can run it from the command line 
by providing a configuration file (usually XML or JSON or YAML).

```bash
./golang_check_webinject-linux-amd64 -c scenario.xml
```

### Command Line Arguments

| Argument | Description |
| :--- | :--- |
| `-c` | Path to the WebInject scenario file (required). |
| `-k` | Support insecure HTTPS (optional). |

### Monitoring Integration (Nagios / Centreon)

To use this plugin within a monitoring environment, define a command in your configuration:

**1. Define the Command**

```nagios
define command {
    command_name    check_webinject
    command_line    $USER1$/golang_check_webinject-linux-amd64 -c $ARG1$
}
```

**2. Define the Service**

```nagios
define service {
    use                     generic-service
    host_name               your_web_server
    service_description     Web_Scenario_Check
    check_command           check_webinject!/etc/nagios/scenarios/login_test.xml
}
```

### Exit Codes (Nagios Standards)

The plugin follows standard Nagios exit codes to report status:

* **0 (OK)**: All tests in the scenario passed successfully.
* **1 (WARNING)**: At least one test failed or a non-critical threshold was met.
* **2 (CRITICAL)**: Connection failure or multiple critical test failures.
* **3 (UNKNOWN)**: Configuration error or scenario file not found.
