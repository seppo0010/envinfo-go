#!/bin/bash
set -Eeux
docker build -t envinfo .

for t in integration-test/*/*; do
	pushd "./$t"
	expected_version=$(basename $t)
	name=$(basename $(dirname $t))
	docker build -t envinfo/test .
	version=$(docker run envinfo/test 2>&1 |grep -w " $name:" |cut -d ':' -f 2 |awk '{print $1}')
	test "$version" == $expected_version
	popd
done
