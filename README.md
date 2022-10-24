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

## 📚 Documentation

- [Official Docs](https://registry.terraform.io/providers/liquid-collective/openzeppelin-defender/latest/docs)

## 🎻 Getting Started

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

## 👋 Contributing

Feedback and contributions to this project are welcome! Before you get started, please review the following:

- [Auth0 Contribution Guidelines](https://github.com/auth0/open-source-template/blob/master/GENERAL-CONTRIBUTING.md)
- [Auth0 Contributor Code of Conduct](https://github.com/auth0/open-source-template/blob/master/CODE-OF-CONDUCT.md)
- [Contribution Guide](CONTRIBUTING.md)

## 🙇 Support & Feedback

### Raise an Issue

If you have found a bug or if you have a feature request, please raise an issue on our
[issue tracker](https://github.com/auth0/terraform-provider-auth0/issues).

### Vulnerability Reporting

Please do not report security vulnerabilities on the public GitHub issue tracker.
The [Responsible Disclosure Program](https://auth0.com/whitehat) details the procedure for disclosing security issues.


---

<div align="center">

<img alt="Auth0 logo and word-mark in black on transparent background" src="https://user-images.githubusercontent.com/28300158/183676042-b9d92893-8fff-408f-9a36-63e77b14be30.png#gh-light-mode-only"  width="20%" height="20%">

<img alt="Auth0 logo and word-mark in white on transparent background" src="https://user-images.githubusercontent.com/28300158/183676141-bea463f9-af82-40ce-b18c-3a1030183d58.png#gh-dark-mode-only"  width="20%" height="20%">

</div>

<br/>

<div align="center">

Auth0 is an easy to implement, adaptable authentication and authorization platform. To learn more checkout
[Why Auth0?](https://auth0.com/why-auth0)

This project is licensed under the MPL-2.0 license. See the [LICENSE](LICENSE) file for more info or
[auth0-terraform-provider.pdf](https://www.okta.com/sites/default/files/2022-03/auth0-terraform-provider.pdf) for a full
report.

</div>