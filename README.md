
# gator : RSS Feed Aggregator

A Go application that scrapes RSS feeds, stores posts in a PostgreSQL database, and allows for continuous fetching in configurable intervals.


## Prerequisites
Before you begin, make sure you have the following installed:

- Go (v1.24.2 or newer)
    - *[Download Go](https://go.dev/dl/)*
- PostgreSQL (v15+ recommended)
    - *[Install PostgreSQL](https://www.postgresql.org/download/)*


## Setup
1. Clone the repository.
`git clone ...`

2. Create a PostgreSQL database. Connect to 'psql' and run:

`CREATE DATABASE gator;`

3. Create a '.gatorconfig.config file in your home directory. 

`~/.gatorconfig.json`

Get your connection string. This is just a URL with the information needed to connect to a database. The format is: 

`protocol://username:password@host:port/database`

Examples:
- macOS: `postgres:arish:@localhost:5432/gator`
- Linux: `postgres:postgres@localhost:5432/gator`

Edit the config file and add the connection string with ssl mode disabled. We need this to use the connection string with `goose`.

`{
    "db_url": "postgres://arish@localhost:5432/gator?sslmode=disable"
}`

4. Migrations

Run database migrations using goose.

`go install github.com/pressly/goose/v3/cmd/goose@latest`

Navigate to the sql/schema directory and run:

`goose postgres "db_connection_string" up`

Example:

`goose postgres postgres://arish:@localhost:5432/gator up`


TODO: explain how to run the program

TODO: List common commands with examples



