#!/bin/bash

set -eou pipefail

POSTGRES_GUFF_PASSWORD="./.secrets/postgres-guff-password"
if [[ ! -f "$POSTGRES_GUFF_PASSWORD" ]]; then
  (>&2 echo "Missing file $POSTGRES_GUFF_PASSWORD");
  exit 1;
fi;

testDb() {
    docker-compose exec db sh -c "psql postgres postgres -c \"SELECT 1;\"" > /dev/null;
}

waitForDb() {
    for i in {1..10}; do
      ret=0;
      testDb || ret=$?;
      if [[ "$ret" -eq 0 ]]; then
        break 2;
      fi;
    done;
}

docker-compose up -d db;
waitForDb;

CREATE="DO \$do\$ BEGIN IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname='guff') THEN CREATE ROLE guff CREATEDB INHERIT LOGIN PASSWORD '`cat "$POSTGRES_GUFF_PASSWORD"`'; END IF; END; \$do\$;"
UPDATE="ALTER ROLE guff CREATEDB INHERIT LOGIN PASSWORD '`cat "$POSTGRES_GUFF_PASSWORD"`';"

docker-compose exec db psql postgres postgres -c "$CREATE"
docker-compose exec db psql postgres postgres -c "$UPDATE"

docker-compose exec db createdb -U guff guff || true
docker-compose exec db createdb -U guff guff_dev || true
docker-compose exec db createdb -U guff guff_test || true
