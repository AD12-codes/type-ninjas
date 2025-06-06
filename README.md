# Type Ninjas

Code faster. Type smarter. Be a Type-Ninja.

## Prerequisites

- Go 1.21+
- PostgreSQL 15+
- SQLC 1.24.0+
- Air (for live reloading)
- golang-migrate

## Installation

1. Clone the repository

```bash
git clone https://github.com/AD12-codes/type-ninjas.git
cd type-ninjas
```

2. Install dependencies

```bash
make deps
```

3. Set up environment variables

```bash
cp .env.example .env
# Update .env with your PostgreSQL credentials
```

## Running the Project

```bash
# Start development server with live reload
make dev

# Build production binary
make build

# Run tests
make test
```

## Project Structure

```
.
├── cmd/          # Main application entrypoint
├── internal/     # Internal application code
├── db/           # Database migrations and queries
├── utils/        # Development utilities
└── Makefile      # Common development tasks
```

## Database Management

```bash
# Create new migration
make migration name=create_users_table

# Run migrations
make migrate-up

# Generate SQLC models
make sqlc-generate
```
