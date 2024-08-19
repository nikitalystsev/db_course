#!/bin/bash

if [ -f .env ]; then
    # shellcheck disable=SC2046
    export $(grep -v '^#' .env | xargs)
fi

migrate_up() {
    migrate -database "$POSTGRES_CREATE_DB_URL"  -path "$POSTGRES_CREATE_DB_MIGRATION_PATH" up
    migrate -database "$POSTGRES_CREATE_SCHEMA_URL" -path "$POSTGRES_CREATE_SCHEMA_MIGRATION_PATH" up
    migrate -database "$POSTGRES_FILL_DB_URL" -path "$POSTGRES_FILL_DB_MIGRATION_PATH" up
}

migrate_down() {
    migrate -database "$POSTGRES_FILL_DB_URL" -path "$POSTGRES_FILL_DB_MIGRATION_PATH" down
    migrate -database "$POSTGRES_CREATE_SCHEMA_URL" -path "$POSTGRES_CREATE_SCHEMA_MIGRATION_PATH" down
    migrate -database "$POSTGRES_CREATE_DB_URL"  -path "$POSTGRES_CREATE_DB_MIGRATION_PATH" down
}

case "$1" in
    up)
        migrate_up
        ;;
    down)
        migrate_down
        ;;
    *)
        echo "Usage: $0 {up|down}"
        exit 1
        ;;
esac