#!/bin/bash

LARAVEL_APP_ROOT=$1
SPECIFIC_TABLE=$2

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
    docker compose -f pixelfed/docker-compose.yml exec -T db mysql -uroot -ppixelfed pixelfed -N -B -e "DESCRIBE $table;"
}

if [ -n "$SPECIFIC_TABLE" ]; then
    describe_table "$SPECIFIC_TABLE"
else
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