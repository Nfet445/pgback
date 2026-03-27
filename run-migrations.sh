#!/bin/bash
# Extract and run only -- +goose Up sections from migrations

for f in ./internal/database/migrations/*.sql; do
    echo "Processing: $f"
    # Extract content between -- +goose Up and -- +goose Down
    awk '/\-\- \+goose Up/{flag=1;next}/\-\- \+goose Down/{flag=0}flag' "$f" | \
    docker exec -i pgbackweb-db psql -U postgres -d pgbackweb 2>&1 | grep -E "(CREATE|TABLE|ERROR|INSERT|ALTER)" | head -5
done
