#! /bin/bash

echo Generating Cassandra Sidecar Go client
export workspace_dir=$(pwd)
export tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_api_client.sh

cd ../../

echo Cleaning Cassandra Sidecar Go client
rm -fr ${tmp_dir}/cassandra_sidecar/.swagger-codegen \
  && rm -fr ${tmp_dir}/cassandra_sidecar/.gitignore \
  && rm -fr ${tmp_dir}/cassandra_sidecar/.swagger-codegen-ignore \
  && rm -fr ${tmp_dir}/cassandra_sidecar/.travis.yml \
  && rm -fr ${tmp_dir}/cassandra_sidecar.conf.json

echo Moving documentations files to root dirs
rm -fr ${tmp_dir}/docs && rm -rf ./docs && mv ${tmp_dir}/cassandra_sidecar/docs .
rm -fr ${tmp_dir}/README.md && mv ${tmp_dir}/cassandra_sidecar/README.md .
