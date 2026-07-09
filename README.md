# gator

A CLI RSS feed aggregator written in Go. Follow your favorite feeds, fetch posts on an interval, and browse them from the terminal. Built as part of the [Boot.dev](https://www.boot.dev) backend course.

## Requirements

You'll need two things installed to run gator:

- **Go** (1.24+) — [install instructions](https://go.dev/doc/install)
- **PostgreSQL** — the aggregator stores users, feeds, and posts in a Postgres database

## Installation

Install the gator CLI with `go install`:

```bash
go install github.com/rcopra/gator@latest
```

This compiles and places the `gator` binary in your `$GOPATH/bin` (usually `~/go/bin`), so make sure that's on your `PATH`.

## Configuration

gator reads its config from `~/.gatorconfig.json`. Create it with your Postgres connection string:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

The `current_user_name` field is managed by gator itself when you log in — you don't need to set it manually.

## Usage

Run commands with `gator <command> [args]`. Start by registering a user:

```bash
gator register <name>    # create a user and log in as them
gator login <name>       # switch to an existing user
```

Then add and follow some feeds:

```bash
gator addfeed <name> <url>   # add a feed and follow it
gator feeds                  # list all feeds
gator follow <url>           # follow an existing feed
gator following              # list feeds you follow
gator unfollow <url>         # unfollow a feed
```

Finally, start the aggregator and browse posts:

```bash
gator agg <interval>     # fetch feeds on a loop, e.g. `gator agg 1m`
gator browse [limit]     # show recent posts from feeds you follow (default 2)
```

Other commands:

```bash
gator users    # list all users
gator reset    # delete all users (and their data)
```
