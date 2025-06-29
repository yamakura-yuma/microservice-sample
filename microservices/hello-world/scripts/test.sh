#!/bin/bash

set -e

PORT=${PORT:-8080}
BASE_URL="http://localhost:$PORT"

echo "== / endpoint =="
curl -i "$BASE_URL/"
echo

echo "== /hello endpoint =="
curl -i "$BASE_URL/hello"
echo
