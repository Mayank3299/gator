# Gator - RSS Feed Aggregator CLI

Gator is a command-line application written in Go that lets you subscribe to RSS feeds, fetch posts, and manage your feed data stored in Postgres.

## Requirements

Before running Gator, make sure you have:

- [Go](https://go.dev/doc/install) (v1.21 or later recommended) installed
- [PostgreSQL](https://www.postgresql.org/download/) running locally or on a server

You should also be using [Git](https://git-scm.com/) to track your changes in this repo.

---

## Installation

Clone the repository:

```bash
git clone https://github.com/Mayank3299/gator.git
cd gator
```

Install the gator CLI:
```bash
go install ./...
```
This will install the gator binary into your ```$GOPATH/bin``` (make sure it’s on your **PATH**).


## Database Setup

Run the migrations using goose:
```bash
goose postgres <your-database-url> up
```
This will create the required tables (users, feeds, feed_follows, posts, etc.) in your database.

## Config File
Gator uses a config file (by default stored in your home directory: ~/.gatorconfig.json) to keep track of the current user and other settings.

To create and set a user:
```bash
gator register your_username
```
This will write the user into your config file.

## Usage
Here are some example commands you can run:

- Register a user
```
gator register myusername
```
-  Set the current user
```
gator login myusername
```
- Reset users and dependent tables
```
gator reset
```
- List Users
```
gator users
```
- Aggregate posts from feed
```
gator agg <time between reqs>
```
- Add a feed
```
gator addfeed "Feed Name" https://example.com/rss
```
- List feeds
```
gator feeds
```
- Follow a feed
```
gator follow https://example.com/rss
```
- Feeds followed by the current user
```
gator following
```
- Unfollow feed
```
gator unfollow https://example.com/rss
```
- Browse posts from followed feeds
```
gator browse
```

---

## Development
To build locally:
```
go build -o gator
```

To run tests:
```
go test ./...
```

### Notes
- Make sure your Postgres instance is running before using the CLI.
- If you update the schema, don’t forget to create a new goose migration.