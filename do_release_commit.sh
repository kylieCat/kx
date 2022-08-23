#!/usr/bin/env bash

PART=$1
[[ "${DIRTY}" ]] && ALLOW_DIRTY="--allow-dirty" || ALLOW_DIRTY=""

get_version_info() {
    part=${1}
    if [[  "${part}" == "rebuild" ]]; then
        current_version=$(cat .bumpversion.cfg | grep "current_version =" | sed 's/current_version = //')
        printf "%s::%s" "${current_version}" "${current_version}"
    else
        out=$(.venv/bin/bumpversion "${part}" --list "${ALLOW_DIRTY}")
        new_version=$(printf "%s\n" "${out}" | grep new_version | sed 's/new_version=//')
        current_version=$(printf "%s\n" "${out}" | grep current_version | sed 's/current_version=//')
        printf "%s::%s\n" "${new_version}" "${current_version}"
    fi
}

get_message () {
    prevVer=${1}
    newVer=${2}
    clog=$(tail -n +3 CHANGELOG.md)
    clog=${clog%\#\#\# ${prevVer}*}
    printf "%s\n\n${clog}" "Incrementing ${part} version from *${prevVer}* -> *${newVer}*"
}

do_commit() {
    prevVer=${1}
    newVer=${2}
    git commit --allow-empty -m "$(get_message "${prevVer}" "${newVer}")"
}

do_tag() {
	git tag ${1}
	git push --tags
}

do_push () {
    git push origin master
}

version_info=$(get_version_info "${PART}")
VERSION=${version_info%::*}
PREVIOUS_VERSION=${version_info#*::}

do_commit "${PREVIOUS_VERSION}" "${VERSION}"
do_tag "${VERSION}"
do_push
