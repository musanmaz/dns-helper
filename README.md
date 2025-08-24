# DNS Helper

**DNS Helper** is a cross-platform (macOS / Linux / Windows) tool that allows you to switch DNS servers with a single command and measure which DNS is faster.

> ⚠️ **Note**: Changing DNS settings requires administrator privileges on most systems.
> - **macOS**: `sudo` with `networksetup` calls
> - **Linux**: `resolvectl` / `nmcli` / or `/etc/resolv.conf` (sudo)
> - **Windows**: PowerShell `Set-DnsClientServerAddress` (Administrator)

**Version**: 0.0.3

## Features

- **Easy switching**: `dns-helper switch cloudflare`
- **Built-in profiles**: `cloudflare`, `google`, `quad9`, `opendns`
- **Benchmarking**: DNS latency comparison (avg/p50/p90, success rate)
- **Status viewing**: See your active DNS settings and interfaces
- **Dry-run mode**: Preview changes before applying them
- **Cross-platform**: Works on macOS, Linux, and Windows

## Installation

### Option 1: Go Installation (Recommended)
If you have Go 1.22+ installed:

```bash
go install github.com/musanmaz/dns-helper/cmd/dns-helper@latest
```

### Option 2: Download Binary
Download the latest release for your platform from the [Releases page](https://github.com/musanmaz/dns-helper/releases).

### Option 3: Build from Source
```bash
git clone https://github.com/musanmaz/dns-helper.git
cd dns-helper
make build
```

## Quick Start

### List Available Profiles
```bash
dns-helper list
```

### Switch DNS Server
```bash
# Use a preset profile
dns-helper switch cloudflare

# Use custom DNS servers
dns-helper switch custom 1.1.1.1 1.0.0.1

# Preview changes without applying
dns-helper switch cloudflare --dry-run
```

### Reset DNS to DHCP Defaults
```bash
# Reset all interfaces to DHCP DNS
dns-helper reset

# Preview what would be reset
dns-helper reset --dry-run
```

### Check Current DNS Status
```bash
dns-helper status
```

### Benchmark DNS Performance
```bash
# Benchmark a specific profile
dns-helper benchmark cloudflare

# Benchmark all profiles
dns-helper benchmark all

# Custom benchmark settings
dns-helper benchmark all --domains google.com,github.com --runs 10 --timeout 2s
```

## Available DNS Profiles

| Profile | Primary DNS | Secondary DNS | Description |
|---------|-------------|---------------|-------------|
| `cloudflare` | 1.1.1.1 | 1.0.0.1 | Fast, privacy-focused DNS |
| `google` | 8.8.8.8 | 8.8.4.4 | Google's public DNS |
| `quad9` | 9.9.9.9 | 149.112.112.112 | Security-focused DNS |
| `opendns` | 208.67.222.222 | 208.67.220.220 | Cisco's OpenDNS |

## Command Reference

### `dns-helper switch [profile|custom] [ip1 ip2 ...]`
Switch to a preset DNS profile or custom IP addresses.

**Flags:**
- `--dry-run`: Show what would happen without making changes

**Examples:**
```bash
dns-helper switch cloudflare
dns-helper switch custom 8.8.8.8 8.8.4.4
dns-helper switch google --dry-run
```

### `dns-helper reset`
Reset DNS settings to DHCP defaults for all network interfaces.

**Flags:**
- `--dry-run`: Show what would happen without making changes

**Examples:**
```bash
dns-helper reset
dns-helper reset --dry-run
```

### `dns-helper status`
Display current DNS settings for all network interfaces.

### `dns-helper list`
Show all available DNS profiles with their IP addresses.

### `dns-helper benchmark [profile|all]`
Measure DNS resolver performance and latency.

**Flags:**
- `--domains`: Comma-separated list of domains to test (default: turk.net,google.com,cloudflare.com)
- `--runs`: Number of queries per domain (default: 5)
- `--timeout`: Single query timeout (default: 1.2s)

**Examples:**
```bash
dns-helper benchmark cloudflare
dns-helper benchmark all --domains example.com,test.com --runs 10
```

## Platform-Specific Details

### macOS
- Uses `networksetup` to configure DNS
- Automatically flushes DNS cache
- Restarts mDNSResponder service
- Requires `sudo` privileges

### Linux
- Prioritizes `systemd-resolved` with `resolvectl`
- Falls back to NetworkManager with `nmcli`
- Last resort: direct `/etc/resolv.conf` modification
- Requires appropriate privileges

### Windows
- Uses PowerShell `Set-DnsClientServerAddress`
- Applies to all active network adapters
- Requires Administrator privileges
- Runs PowerShell with execution policy bypass

## Development

### Prerequisites
- Go 1.22+
- Make (optional, for build automation)

### Build
```bash
make build
```

### Test
```bash
make test
```

### Run
```bash
make run
```

### Release
```bash
make release
```

## Project Structure

```
dns-helper/
├── cmd/dns-helper/     # Main application entry point
├── internal/
│   ├── bench/          # DNS benchmarking logic
│   ├── cli/            # Command-line interface
│   ├── platform/       # Platform-specific DNS operations
│   ├── resolvers/      # DNS profile definitions
│   └── util/           # Utility functions
├── Makefile            # Build automation
└── README.md           # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Security

⚠️ **Important**: This tool modifies system DNS settings, which can affect your internet connectivity. Always use `--dry-run` first to preview changes, and ensure you have a backup plan if something goes wrong.

## Troubleshooting

### Permission Denied
- **macOS/Linux**: Use `sudo` before the command
- **Windows**: Run PowerShell as Administrator

### DNS Changes Not Applied
- Try flushing DNS cache manually
- Restart network services
- Check if your system uses a managed DNS configuration

### Benchmark Fails
- Ensure internet connectivity
- Check firewall settings
- Try different domains or increase timeout

## Support

If you encounter issues or have questions:
1. Check the [Issues](https://github.com/musanmaz/dns-helper/issues) page
2. Create a new issue with detailed information
3. Include your operating system and Go version

---

**Made with ❤️ by [Mehmet Şirin Usanmaz](https://github.com/musanmaz)**
