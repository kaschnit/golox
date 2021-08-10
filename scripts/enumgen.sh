#!/bin/bash
HERE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_ROOT=$(dirname ${HERE})
ENUM_CONFIGS_PATH="${PROJECT_ROOT}/configs/enum"
GENERATOR="${PROJECT_ROOT}/tools/codegen/enumgen_golang"

function generate() {
	echo "Generating enum from ${ENUM_CONFIGS_PATH}/$1"
	go run ${GENERATOR} -json ${ENUM_CONFIGS_PATH}/$1 -output $2
}

generate "TokenType.json" "${PROJECT_ROOT}/pkg/token/tokentype"