# Health Checker (healthCli)

Minimal TUI-based health monitoring tool written in Go. It periodically checks a list of services (URLs), measures latency, and displays a terminal dashboard with statuses and a latency graph.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [Configuration](#configuration)
- [How it works](#how-it-works)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- Periodic HTTP GET checks with timeout and latency measurement.
- TUI dashboard built with `bubbletea` and styled with `lipgloss`.
- Latency graph using `asciigraph`.
- Simple, easy-to-extend code structure under `cmd/` and `internal/`.

## Prerequisites

- Go 1.24 or newer
- Git (optional)
- Make (optional)

## Quick Start

Clone the repository and run:

```sh
git clone https://github.com/Arya911/healthCli.git
cd "Health Checker"
make run
```

Or run directly with Go:

```sh
go run cmd/healthcmd/main.go
```

Build a standalone binary:

```sh
go build -o health cmd/healthcmd/main.go
./health
```

## Usage

- When the app starts it will display a dashboard with the list of monitored services, their current status (UP/DOWN/SLOW), the latest latency, and a small ASCII graph for the selected service.
- Use the arrow keys to change the selected service. Press `q` to quit.

## Configuration

Services and behavior are currently configured in code. Edit the `services` slice in:

```
internal/checker/monitor.go
```

The default services are defined in `internal/checker/monitor.go`. You can add or remove URLs there. The check interval is controlled by the ticker (default 5s), and the per-request timeout is set in `internal/checker/runner.go` (default 2s).

Example: open `internal/checker/monitor.go` and update the `services` variable.

## How it works

- `cmd/healthcmd/main.go` wires together the checker and the TUI.
- `internal/checker` runs periodic goroutines that perform HTTP GET requests, measure latency, and send `model.Result` messages on a channel.
- `internal/tui` consumes results and renders a dashboard using `bubbletea`, `lipgloss`, and `asciigraph`.

## Development

- The project uses Go modules; dependencies are declared in `go.mod`.
- To build locally: `go build -o health cmd/healthcmd/main.go`

- Useful files:
  - `cmd/healthcmd/main.go` — program entrypoint
  - `internal/checker` — checking logic and scheduler
  - `internal/tui` — terminal UI model, view and styles
  - `internal/model` — shared data structures

## Contributing

- Pull requests and improvements are welcome. If you add features, consider making the service list configurable via a file or environment variables.

## License

No license is included in this repository. Add a `LICENSE` file if you want to publish the project under an open-source license.

----
