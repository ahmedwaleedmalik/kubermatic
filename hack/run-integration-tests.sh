#!/usr/bin/env bash

# Copyright 2021 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
# Required for signal propagation to work so
# the cleanup trap gets executed when the script
# receives a SIGINT
set -o monitor

cd $(dirname $0)/..
source hack/lib.sh

TEST_NAME="Pre-warm Go build cache"
echodate "Attempting to pre-warm Go build cache..."

beforeGocache=$(nowms)
make download-gocache
pushElapsed gocache_download_duration_milliseconds $beforeGocache

export CGO_ENABLED=1

LOCALSTACK_TAG="${LOCALSTACK_TAG:-3.0.1}"
LOCALSTACK_IMAGE="${LOCALSTACK_IMAGE:-localstack/localstack:$LOCALSTACK_TAG}"

if [[ ! -z "${JOB_NAME:-}" ]] && [[ ! -z "${PROW_JOB_ID:-}" ]]; then
  start_docker_daemon_ci
fi

echodate "Running integration tests..."

# Run integration tests and only integration tests by:
# * Finding all files that contain the build tag via grep
# * Extracting the dirname as the `go test` command doesn't play well with individual files as args
# * Prefixing them with `./` as that's needed by `go test` as well
for file in $(grep --files-with-matches --recursive --extended-regexp '//go:build.+integration' cmd/ pkg/ | xargs dirname | sort -u); do
  echodate "Testing package ${file}..."

  if [[ "$file" =~ .*vsphere.* ]] && provider_disabled vsphere; then
    continue
  fi

  extraArgs=()
  if [[ "$file" =~ .*helmclient.* ]] || [[ "$file" =~ .*providers/template.* ]]; then
    extraArgs+=(-registry-hostname "$HOSTNAME")
  fi

  go_test $(echo $file | sed 's/\//_/g') -tags "integration,${KUBERMATIC_EDITION:-ce}" -race ./${file} -v "${extraArgs[@]}"
done
