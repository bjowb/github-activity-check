# GitHub Events CLI

A simple and fast command-line tool written in Go for fetching **recent GitHub activity/events** of any public GitHub user.
Sample solution for the [github-user-activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh).

This tool uses the GitHub REST API and displays events like:

- PushEvent
- PullRequestEvent
- IssuesEvent
- CreateEvent
- ForkEvent
- WatchEvent  
…and more.

---

## Features

- Fetches the **latest public GitHub events** for any user
- Clean, readable output
- Optional GitHub Token support for higher rate limits
- Lightweight (no external dependencies)
- Works on Linux, macOS, and Windows

---

##  Additions You may want to Do

- If you want to increase your limit of events you can use tokens upon authorization

## Installation

### Clone the repo:
```bash
git clone https://github.com/bjowb/github-events-cli.git
cd github-events-cli
