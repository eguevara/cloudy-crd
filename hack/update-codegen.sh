#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z "${GOPATH}" ]]; then
  GOPATH=~/go
fi

if [[ ! -d "${GOPATH}/src/k8s.io/code-generator" ]]; then
  echo "k8s.io/code-generator missing from GOPATH"
  exit 1
fi

cd ${GOPATH}/src/k8s.io/code-generator

./generate-groups.sh \
  all \
  github.com/eguevara/cloudy-crd/pkg/generated \
  github.com/eguevara/cloudy-crd/pkg/apis \
  cloudy:v1\
  $@
