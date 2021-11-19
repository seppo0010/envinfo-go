#!/bin/bash
set -Eeux
docker build -t envinfo .

tests=${@:-integration-test/*/*/*}
for t in $tests; do
	pushd "./$t"
	expected_value=$(cat expected) || expected_value=$(basename $t)
	name=$(basename $(dirname $t))
	docker build -t envinfo/test .
	out="$(docker run envinfo/test 2>&1)"
	value=$(echo -n "$out" |grep -w " $name:" |cut -d ':' -f 2)
	if [[ "${value:1}" != $expected_value* ]]; then
		exit 1;
	fi
	popd
done

echo "Test finished successfully!"
