#!/usr/bin/env bash

# Instructions
# You will need Java installed on the local machine

# Params
echo Exporting Params
export client_name=${client_name:-instaclustr_icarus}

export codegen_filename=${codegen_filename:-swagger-codegen-cli-3.0.20.jar}
export tmp_dir=${tmp_dir:-${HOME}/Projects/tmp}
export client_dir=${tmp_dir}/${client_name}
export mustache_dir=./swagger_templates
export api_def_dir=./api_defs

export codegen_url=https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.20/${codegen_filename}

export swagger_def=$(ls ${api_def_dir} | grep ${client_name} | sort -V | tail -1)

echo Prepping Workspace
mkdir -p ${tmp_dir}
echo "{ \"packageName\": \"${client_name}\" }" > ${tmp_dir}/${client_name}.conf.json

if [ ! -f ${tmp_dir}/${codegen_filename} ];
then
  echo Downloading ${codegen_filename}
  wget -N ${codegen_url} -P ${tmp_dir}
fi

echo Generating into ${client_dir}
java -jar ${tmp_dir}/${codegen_filename} generate \
    --lang go \
    --config ${tmp_dir}/${client_name}.conf.json \
    --api-package apis \
    --model-package models \
    --input-spec ${api_def_dir}/${swagger_def} \
    --output ${client_dir}
