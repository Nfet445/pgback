#!/bin/bash
export PATH=$HOME/go-install:$HOME/go-install/go/bin:$PATH
export PBW_POSTGRES_CONN_STRING="postgresql://postgres:postgres@localhost:5433/pgbackweb?sslmode=disable"
export PBW_ENCRYPTION_KEY="dev-secret-key-change-in-production"
export PBW_LISTEN_HOST="0.0.0.0"
export PBW_LISTEN_PORT="8085"

# Check if goose is available
if which goose > /dev/null 2>&1; then
    echo "Running migrations..."
    goose -dir ./internal/database/migrations postgres "$PBW_POSTGRES_CONN_STRING" up
else
    echo "Goose not available, skipping migrations..."
fi

# Generate SQLC if needed
if which sqlc > /dev/null 2>&1; then
    echo "Generating SQLC code..."
    sqlc generate
fi

# Build static files
echo "Building static files..."
npm run tailwindcss -- --minify --config ./tailwind.config.ts --input ./internal/view/static/css/style.css --output ./internal/view/static/build/style.min.css
node ./scripts/build-js.ts

# Run the app
echo "Starting application..."
go run ./cmd/app/.
