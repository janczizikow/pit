# pit

API for pit submissions website.

## Installation

```sh
go mod download
```

## Usage

```sh
# Runs database migrations
make migrate

# Creates a new migration with a given name
make migration name=create_table_name

# Starts the HTTP server
go run ./cmd/api
```
