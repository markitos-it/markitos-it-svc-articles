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

echo "== protoc version =="
protoc --version

echo "== include dirs check =="
if [ -f /usr/include/google/protobuf/timestamp.proto ]; then
  ls -la /usr/include/google/protobuf/timestamp.proto
else
  echo "== /usr/include/google/protobuf/timestamp.proto not found =="
fi

if [ -f /usr/local/include/google/protobuf/timestamp.proto ]; then
  ls -la /usr/local/include/google/protobuf/timestamp.proto
else
  echo "== /usr/local/include/google/protobuf/timestamp.proto not found =="
fi

echo "== generating protobuf go files =="
protoc -I. -I/usr/include -I/usr/local/include \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/article.proto

echo "== generated files =="
ls -la proto