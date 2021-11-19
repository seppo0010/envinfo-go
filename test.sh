#!/bin/bash
set -Eeux
docker build -t envinfo .

for t in integration-test/*/*; do
	pushd "./$t"
	docker build -t envinfo/test .
	version=$(docker run envinfo/test 2>&1)
	test "$version" == $(basename $t)
	popd
done
