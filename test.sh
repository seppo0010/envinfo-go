#!/bin/bash
set -Eeux
docker build -t envinfo .

tests=${@:-integration-test/*/*}
for t in $tests; do
	pushd "./$t"
	expected_version=$(basename $t)
	name=$(basename $(dirname $t))
	docker build -t envinfo/test .
	version=$(docker run envinfo/test 2>&1 |grep -w " $name:" |cut -d ':' -f 2 |awk '{print $1}')
	test "$version" == $expected_version
	popd
done

echo "Test finished successfully!"
