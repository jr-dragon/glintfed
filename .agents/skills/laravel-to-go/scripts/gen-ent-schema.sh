#!/bin/bash

LARAVEL_APP_ROOT=$1

grep -lE 'extends[[:space:]]+(Model|Authenticatable)' $LARAVEL_APP_ROOT/app/*.php $LARAVEL_APP_ROOT/app/Models/*.php 2>/dev/null | \
   sed 's|.*/||; s/\.php$//' | \
   xargs -I {} go run -mod=mod entgo.io/ent/cmd/ent new {}