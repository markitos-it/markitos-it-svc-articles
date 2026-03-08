#!/bin/bash
#:[.'.']:>- ===================================================================================
#:[.'.']:>- Marco Antonio - markitos devsecops kulture
#:[.'.']:>- The Way of the Artisan
#:[.'.']:>- markitos.es.info@gmail.com
#:[.'.']:>- 🌍 https://github.com/orgs/markitos-it/repositories
#:[.'.']:>- 🌍 https://github.com/orgs/markitos-public/repositories
#:[.'.']:>- 📺 https://www.youtube.com/@markitos_devsecops
#:[.'.']:>- ===================================================================================
set -e

echo "🧹 Cleaning build artifacts..."

echo "🧹 Cleaning Docker containers, images, and volumes..."
docker rm markitos-it-svc-articles markitos-it-svc-articles-postgres 2>/dev/null || true
docker image rm postgres:17-alpine markitos-it-svc-article-markitos-it-svc-articles:latest 2>/dev/null || true
docker volume rm markitos-it-svc-article_markitos-it-svc-articles-postgres_data 2>/dev/null || true
echo "✅ Removed Docker containers, images, and volumes"

echo "🧹 Cleaning generated protobuf files..."
rm -f proto/*.pb.go
rm -f proto/*_grpc.pb.go
echo "✅ Removed generated protobuf files"

echo "🧹 Cleaning Go build cache..."
go clean -testcache
go clean -cache
go clean -modcache
echo "✅ Cleaned Go build cache"

echo "✨ Clean complete"