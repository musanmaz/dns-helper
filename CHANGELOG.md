# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.1] - 2025-08-23

### Added
- Initial release of DNS Helper
- Cross-platform DNS switching (macOS, Linux, Windows)
- Built-in DNS profiles (Cloudflare, Google, Quad9, OpenDNS)
- DNS benchmarking with latency measurements
- Status checking for current DNS configuration
- Dry-run mode for safe testing
- Command-line interface with Cobra
- Version command to display build information
- Comprehensive English documentation
- Enhanced CI/CD pipeline with multi-platform testing
- Improved GoReleaser configuration for automated releases
- **Reset command** to restore DNS settings to DHCP defaults
- **Port number stripping** for all platforms (macOS, Linux, Windows)
- **Debug output** for better troubleshooting
- Comprehensive unit test suite
- Integration tests for real DNS operations
- Test helpers and utilities
- Environment-based test configuration

### Supported Platforms
- **macOS**: Uses `networksetup` and `dscacheutil`
- **Linux**: Supports `systemd-resolved`, NetworkManager, and direct `/etc/resolv.conf`
- **Windows**: PowerShell-based configuration using `Set-DnsClientServerAddress`

---

## Version Information

- **Go Version**: 1.22+
- **License**: Apache License 2.0
- **Author**: Mehmet Şirin Usanmaz
