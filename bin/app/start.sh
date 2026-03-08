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

cd "$(dirname "$0")/../.."

make proto

export GRPC_PORT=${GRPC_PORT:-3004}
export GRPC_TLS_ENABLED=${GRPC_TLS_ENABLED:-false}
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-55434}
export DB_USER=${DB_USER:-markitos-it-svc-articles}
export DB_PASS=${DB_PASS:-markitos-it-svc-articles}
export DB_NAME=${DB_NAME:-markitos-it-svc-articles}

echo "🚀 Starting markitos-it-svc-articles (Go)..."
echo "📡 GRPC_PORT.......: $GRPC_PORT"
echo "📦 DB_HOST.........: $DB_HOST:$DB_PORT"
echo "📦 DB_USER.........: $DB_USER"
echo "📦 DB_NAME.........: $DB_NAME"
echo "📡 GRPC_TLS_ENABLED: $GRPC_TLS_ENABLED"
echo ""

docker compose up