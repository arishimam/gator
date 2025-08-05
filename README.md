
# gator : RSS Feed Aggregator

`gator` is a Go application that scrapes RSS feeds, stores posts in a PostgreSQL database, and supports continuous feed fetching in customizable intervals.

---

## Prerequisites

Before you begin, make sure you have the following installed:

- **Go** (v1.24.2 or newer)
    - *[Download Go](https://go.dev/dl/)*
- **PostgreSQL** (v15+ recommended)
    - *[Install PostgreSQL](https://www.postgresql.org/download/)*


## Setup

1. Clone the repository.

`git clone https://github.com/arishimam/gator.git`

`cd gator`

2. Create a PostgreSQL database. Connect to 'psql' and run:

`CREATE DATABASE gator;`

3. Create the config file

Crease a `.gatorconfig.config` file in your home directory. 

`~/.gatorconfig.json`

Get your connection string. This is just a URL with the information needed to connect to a database. The format is: 

`protocol://username:password@host:port/database`

Examples:
- macOS: `postgres:arish:@localhost:5432/gator`
- Linux: `postgres:postgres@localhost:5432/gator`

Edit the config file and add the connection string with ssl mode disabled.

`{
    "db_url": "postgres://arish@localhost:5432/gator?sslmode=disable"
}`

4. Migrations

Run database migrations using goose.

`go install github.com/pressly/goose/v3/cmd/goose@latest`

Navigate to the sql/schema directory and run:

`goose postgres "your_connection_string_here" up`

Example:

`goose postgres postgres://arish:@localhost:5432/gator up`


5. Running the Program

- Navigate to main repo directory. Build the binary

`go build -o gator .`

- Run with a command
`./gator <command> [args...]`


- Common commands with examples

- Auth & User Commands:
    - Register a new user

    `./gator register arish`

    - Login an existing user

    `./gator login arish`

    - List all registered user

    `./gator users` 

- Feed Aggregration & Viewing:
    - Continuously fetch feeds on interval

    `./gator agg 1m`

    - View all available feeds

    `./gator feeds` 

    - Add a new feed (must be logged in)

    `./gator addfeed <url>` 

    - View latest posts from followed feeds

    `./gator browse [limit]` 

- Following Feeds

    - Follow a feed

    `./gator follow <url>`

    - Unfollow a feed

    `./gator unfollow <url>`

    - List feeds you are currently following

    `./gator following`









