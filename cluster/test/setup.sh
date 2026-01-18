#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROVIDER_DIR=$( readlink -f "$SCRIPT_DIR/../.." )

echo "Creating minio deployment"
${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/minio.yaml"

echo "Creating default provider"
${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/providerconfig.yaml"
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
