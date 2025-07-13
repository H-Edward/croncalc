# CronCalc

A simple web application to calculate the next scheduled times for cron expressions.

## Features

- Parse cron expressions and get the next 5 scheduled times
- Support for different timezones
- Common cron expression examples
- Simple, easy-to-use interface (tbc)

## API Endpoints

The application provides the following API endpoints:

- `/api/parse` - Parse a cron expression and return the next 5 scheduled times
  - Query Parameters:
    - `expr`: The cron expression to parse
    - `tz` (optional): The timezone to use (defaults to UTC)

- `/api/timezones` - Get a list of available timezones

## Running the Application

1. Ensure you have Go installed on your system
2. Clone this repository
3. Run the application:

```bash
go run main.go
# Or build it first and run
go build -o croncalc
./croncalc
```

By default, the application runs on port 8010, but you can configure this in the .env

## Cron Expression Format

CronCalc supports standard cron expressions with five fields:

```text
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

It also supports special expressions like `@reboot`.