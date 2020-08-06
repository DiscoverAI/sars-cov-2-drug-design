#!/usr/bin/env bash

set -e
set -o nounset
set -o pipefail

script_dir=$(cd "$(dirname "$0")" ; pwd -P)
app_name=dev-env

goal_assume-role() {
  pushd "${script_dir}/assume-role" > /dev/null
    base_profile="tortugas-base"
    role="${1:-}"
    if [ -z "${role}" ]; then
      echo "Role name not supplied. Usage: <ROLE:developer|administrator>"
      exit 1
    fi

    ./go assume-role ${base_profile} "tortugas/${role}"
  popd > /dev/null
}


TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    assume-role             - Assumes Role from argument and updates Shared Credentials with valid session
"
  exit 1
fi
