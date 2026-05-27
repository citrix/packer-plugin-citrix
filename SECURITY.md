# Security Policy

## Supported Versions

| Version | Supported |
|---------|-----------|
| 1.0.x   | Yes       |

## Reporting a Vulnerability

Please **do not** report security vulnerabilities through public GitHub issues.

Use [GitHub's private vulnerability reporting](https://docs.github.com/en/code-security/security-advisories/guidance-on-reporting-and-writing/privately-reporting-a-security-vulnerability)
to submit a report. You can access this via the **Security** tab of this repository.

We will acknowledge receipt within 5 business days and aim to provide a resolution
timeline within 15 business days.

## Scope

This plugin communicates with an on-premises Citrix App Layering (ELM) appliance
over HTTPS (port 443) and with CE VMs over WinRM (port 5985). Relevant security
considerations include:

- **Credentials**: ELM credentials and CE VM credentials are passed via HCL
  variables. Always use `sensitive = true` on password variables and supply
  values through environment variables (`PKR_VAR_*`) or a secrets manager —
  never hardcode credentials in template files.
- **TLS verification**: `insecure_connection = true` disables TLS certificate
  verification. Use this only in lab environments with self-signed certificates,
  never in production.
