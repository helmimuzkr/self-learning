#!/bin/bash

app="app1"

echo "=== starting to remove containerr $app ==="
docker container stop $app && docker container rm $app
echo "=== remove container succedeed ==="
echo ""

echo "=== starting to build image $app ==="
docker image rm $app && docker image build -t $app .
echo "=== build succedeed ==="
echo ""

echo "=== starting to compose $app ==="
docker compose down && docker compose up -d
echo "=== docker compse succedeed ==="
