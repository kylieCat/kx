#!/usr/bin/env bash

PART=$1

get_message () {
    part=${1}
    prevVer=$(waypoint latest)
    newVer=$(waypoint bump --${part})
    clog=$(tail -n +3 CHANGELOG.md)
    clog=${clog%\#\#\# ${prevVer}*}
    printf "%s\n\n${clog}" "Incrementing ${part} version from *${prevVer}* -> *${newVer}*"
}

do_commit() {
   part=${1}
   git commit --allow-empty -m "$(get_message ${part})"
}

do_tag() {
	git tag $(waypoint latest kx)
	git push --tags
}

do_push () {
    git push origin master
}

do_commit ${PART}
do_push
do_tag