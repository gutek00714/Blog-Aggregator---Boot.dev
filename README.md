# Gator

Gator is a CLI tool that aggregates RSS feeds from blogs you follow and stores their posts in a PostgreSQL database.

## Prerequisites

- **Go** 1.22 or higher
- **PostgreSQL** — a running instance of a Postgres database

## Installation
```bash
go install github.com/gutek00714/Blog-Aggregator---Boot.dev@latest
```

## Configuration

Gator requires a `.gatorconfig.json` file in your home directory. Create it with the following structure:
```json
{
  "db_url": "postgres://username:password@localhost:5432/gator",
  "current_user_name": "your_name"
}
```

## Usage

| Command | Description |
|---|---|
| `gator register <name>` | Register a new user |
| `gator addfeed <name> <url>` | Add an RSS feed to follow |
| `gator agg <interval>` | Start the aggregator (e.g. `1m`, `30s`) |
| `gator browse <limit>` | Browse the latest posts |
| `gator feeds` | List all available feeds |
| `gator follow <url>` | Follow a feed |
| `gator unfollow <url>` | Unfollow a feed |
| `gator following` | List feeds you follow |
