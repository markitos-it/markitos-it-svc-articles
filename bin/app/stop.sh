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

echo "🛑 Stopping markitos-it-svc-articles PostgreSQL..."
docker compose down -v --remove-orphans markitos-it-svc-articles-postgres markitos-it-svc-articles > /dev/null 2>&1 || true
echo "✅ markitos-it-svc-articles PostgreSQL stopped."
echo
