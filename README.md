<div align="center">

# Forge

**Cross-Platform Git Activity Tracker & Streak Monitor**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Git](https://img.shields.io/badge/Git-F05032?style=for-the-badge&logo=git&logoColor=white)](https://git-scm.com)
[![systemd](https://img.shields.io/badge/systemd-000000?style=for-the-badge&logo=linux&logoColor=white)](https://systemd.io)
[![Beeep](https://img.shields.io/badge/Beeep-6f42c1?style=for-the-badge&logo=go&logoColor=white)](https://github.com/gen2brain/beeep)
[![JSON](https://img.shields.io/badge/JSON-000000?style=for-the-badge&logo=json&logoColor=white)](https://www.json.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)

*Multi-Repo Tracking · Daily Streak Monitoring · Desktop Notifications · Background Watch Mode · CLI Diagnostics*

</div>

---

A lightweight cross-platform Git activity tracker that helps developers maintain their daily coding streak through repository tracking, reminders, and background monitoring.

---

## Problem

Developers often forget to commit code consistently, especially when working across multiple repositories. Missing even a single day can break personal coding habits or contribution streaks.

Forge solves this by tracking commits across multiple Git repositories, sending desktop reminders when no commits are detected, and continuously monitoring development activity in the background.

---

## Features

**Repository Management** — Register, remove, and list tracked repositories with duplicate protection.

**Activity Monitoring** — Count today's commits across all tracked repositories and display daily streak status.

**Notifications** — Send immediate desktop reminders and run a background watch mode to prevent missed days.

**Diagnostics** — Verify Git installation, configuration integrity, and notification availability.

---

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | CLI development |
| Git CLI | Commit tracking |
| Beeep | Desktop notifications |
| systemd | Background service (Linux) |
| JSON | Configuration storage |

---

## Installation

```bash
git clone https://github.com/brijnandan/forge.git
cd forge
go build -o forge ./cmd
sudo cp forge /usr/local/bin/
forge version
```

---

## Commands

| Command | Description |
|---------|-------------|
| `forge add` | Register current repository |
| `forge remove` | Remove current repository |
| `forge list` | List tracked repositories |
| `forge status` | Show today's commit activity |
| `forge remind` | Send a reminder immediately |
| `forge watch` | Start background monitoring |
| `forge doctor` | Diagnose installation |
| `forge version` | Display version |

---

## How It Works

Repositories are stored in `~/.config/forge/config.json`:

```json
{
  "repositories": [
    "/home/user/projects/forge",
    "/home/user/projects/api"
  ]
}
```

`forge status` reads this file, iterates through each repository, and runs:

```bash
git rev-list --count --since=midnight HEAD
```

`forge watch` polls periodically and triggers a desktop notification if no commits are detected.

---

## Project Structure

```
forge/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   └── git/
│       └── git.go
├── go.mod
├── go.sum
└── README.md
```

---

## Example Output

```
Forge Status

Tracked Repositories    3
Commits Today           7
Streak Status           SAFE
```

```
Forge Doctor

  Config file found
  Git installed
  3 repositories configured
  Notifications available
```

---

## Roadmap

- Windows Task Scheduler support
- macOS LaunchAgent support
- Configurable reminder intervals
- GitHub contribution graph integration
- Export commit statistics

---

## License

MIT License
