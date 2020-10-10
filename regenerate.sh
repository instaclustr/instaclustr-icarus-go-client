#! /bin/bash

echo Generating Instaclustr Icarus Go client
export workspace_dir=$(pwd)
export tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_api_client.sh

cd ../../

echo Cleaning Cassandra Sidecar Go client
rm -fr ${tmp_dir}/instaclustr_icarus/.swagger-codegen \
  && rm -fr ${tmp_dir}/instaclustr_icarus/.gitignore \
  && rm -fr ${tmp_dir}/instaclustr_icarus/.swagger-codegen-ignore \
  && rm -fr ${tmp_dir}/instaclustr_icarus/.travis.yml \
  && rm -fr ${tmp_dir}/instaclustr_icarus.conf.json

echo Moving documentations files to root dirs
rm -fr ${tmp_dir}/docs && rm -rf ./docs && mv ${tmp_dir}/instaclustr_icarus/docs .
rm -fr ${tmp_dir}/README.md && mv ${tmp_dir}/instaclustr_icarus/README.md .
