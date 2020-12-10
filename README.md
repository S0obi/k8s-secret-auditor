# k8s-secret-auditor

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=S0obi_k8s-secret-auditor&metric=alert_status)](https://sonarcloud.io/dashboard?id=S0obi_k8s-secret-auditor)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

A simple Kubernetes secrets auditor that will help you finding weak passwords.

## Installation

For now, you must build k8s-secret-auditor from source :

```
make build
```

Binary will be in the _bin directory.

## Usage

```
NAME:
   k8s-secret-auditor - Audit Kubernetes secrets

USAGE:
   k8s-secret-auditor.exe [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --namespace value, -n value, --ns value  Audit a specific namespace
   --config value, -c value, --conf value   Set a specific config file (default: "config.yaml")
   --help, -h                               show help (default: false)
```