# Terraform Provider OpenZeppelin Defender

[![GoDoc](https://pkg.go.dev/badge/github.com/liquid-collective/terraform-provider-openzeppelin-defender.svg)](https://pkg.go.dev/github.com/liquid-collective/terraform-provider-openzeppelin-defender)
[![Go Report Card](https://goreportcard.com/badge/github.com/liquid-collective/terraform-provider-openzeppelin-defender)](https://goreportcard.com/report/github.com/liquid-collective/terraform-provider-openzeppelin-defender)
[![Release](https://img.shields.io/github/v/release/liquid-collective/terraform-provider-openzeppelin-defender?logo=terraform&include_prereleases&style=flat-square)](https://github.com/liquid-collective/terraform-provider-openzeppelin-defender/releases)
[![Codecov](https://img.shields.io/codecov/c/github/liquid-collective/terraform-provider-openzeppelin-defender?logo=codecov&style=flat-square)](https://codecov.io/gh/liquid-collective/terraform-provider-openzeppelin-defender)
[![License](https://img.shields.io/github/license/liquid-collective/terraform-provider-openzeppelin-defender.svg?logo=fossa&style=flat-square)](https://github.com/liquid-collective/terraform-provider-openzeppelin-defender/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/liquid-collective/terraform-provider-openzeppelin-defender/Main/master?logo=github&style=flat-square)](https://github.com/liquid-collective/terraform-provider-openzeppelin-defender/actions?query=branch%3Amaster)

OpenZeppelin Defender Terraform Provider is a plugin for managing Admin proposals on OpenZeppelin using
[Terraform](https://www.terraform.io/) tool.

---

## Documentation

- [Official Docs](https://registry.terraform.io/providers/liquid-collective/openzeppelin-defender/latest/docs)

## Getting Started

### Requirements

- [Terraform](https://www.terraform.io/downloads)
- A [OpenZeppelin Defender](https://defender.openzeppelin.com/) account

### Installation

This provider is available on [Terraform Registry](https://registry.terraform.io/). 

To use it this provider, copy and paste the following code into your Terraform configuration.

```terraform
terraform {
  required_providers {
    defender = {
      source = "liquidcollective.io/admin/openzeppelin-defender"
    }
  }
}

provider "defender" {
  api_key = "<api-key>"
  api_secret = "<api-key>"
}
```

Then, run 

```sh
$ terraform init
```
