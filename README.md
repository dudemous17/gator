# Gator: RSS Feed Aggregator

A CLI-based RSS feed aggregator written in Go. This tool allows you to manage users, subscribe to RSS feeds, and fetch the latest posts directly to your terminal.

## Prerequisites

To run this project, you need the following installed on your system:

* **Go**: Version 1.22 or higher.
* **PostgreSQL**: Ensure you have a running Postgres instance and a database created for the project.

## Installation

You can install the `gator` CLI directly using the `go install` command:

```bash
go install -x github.com/dudemous17/gator@latest
```

## Setup & Configuration

Before running the program, you need to create a configuration file named `.gatorconfig.json` in your home directory. This file tells the application which user is currently logged in and how to connect to your database.

1. **Create the config file:**
   Create a file at `~/.gatorconfig.json` with the following structure:

   ```json
   {
     "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
     "current_user_name": "your_username"
   }
   ```

2. **Initialize the Database:**
   Make sure your database schema is up to date (using your preferred migration tool) before starting.
   The files in sql/schema are designed with goose for schema migration. Example:
   ```bash
   goose postgres "postgres://username:password@localhost:5432/gator" up
   ```

## Usage & Commands

Once installed and configured, you can interact with the program using various commands. Here are a few examples:

* **Usage:**
```bash
gator
```
**Register a user:**
  ```bash
  gator register <username>
  ```
* **Login:**
  ```bash
  gator login <username>
  ```
* **Add a feed:**
  ```bash
  gator addfeed <name> <url>
  ```
* **Follow a feed:**
  ```bash
  gator follow <url>
  ```
* **Fetch posts:**
  ```bash
  gator agg <time_between_reqs>
  ```
  *(e.g., `gator agg 1m` to fetch updates every minute)*
