#!/bin/bash

LARAVEL_APP_ROOT=$1
SPECIFIC_TABLE=$2

# MySQL configuration - adjust these or use environment variables
DB_HOST=${DB_HOST:-"127.0.0.1"}
DB_PORT=${DB_PORT:-"3306"}
DB_USER=${DB_USER:-"root"}
DB_PASS=${DB_PASS:-"pixelfed"}
DB_NAME=${DB_NAME:-"pixelfed"}

pluralize() {
    local word=$1
    echo "$word" | sed -r '
        /([sS][hH]|[cC][hH]|[xX]|[sS])$/s/$/es/I;
        /[^aeiouAEIOU]y$/s/y/ies/I;
        /([aeiouAEIOU]y|[sS]|[xX]|[cC][hH]|[sS][hH]|[yY])$/!s/$/s/I;
    '
}

describe_table() {
    local table=$1
    echo "--- Table: $table ---"
    mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" -N -B -e "DESCRIBE $table;"
}

if [ -n "$SPECIFIC_TABLE" ]; then
    describe_table "$SPECIFIC_TABLE"
else
    if [ -z "$LARAVEL_APP_ROOT" ]; then
        echo "Usage: $0 <laravel_app_root> [specific_table]"
        exit 1
    fi

    grep -lE 'extends[[:space:]]+(Model|Authenticatable)' $LARAVEL_APP_ROOT/app/*.php $LARAVEL_APP_ROOT/app/Models/*.php 2>/dev/null | while read f; do \
        m=$(basename "$f" .php); \
        t=$(grep "protected \$table =" "$f" | sed -E "s/.*['\"](.*)['\"].*/\1/"); \
        if [ -z "$t" ]; then \
            snake_m=$(echo "$m" | sed -E 's/([a-z0-9])([A-Z])/\1_\2/g' | tr '[:upper:]' '[:lower:]'); \
            t=$(pluralize "$snake_m"); \
        fi; \
        describe_table "$t"; \
    done
fi
