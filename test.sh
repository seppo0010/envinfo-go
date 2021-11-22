#!/bin/bash
set -Eeux
docker build -t envinfo .

tests=${@:-tests/*/*/*}
for t in $tests; do
	pushd "./$t"
	expected_value=$(cat expected) || expected_value=$(basename $t)
	if [ -f search ]; then
		search=$(cat search)
	else
		name="$(basename $(dirname $t))"
		name=${name/_/ }
		search=" $name:"
	fi
	docker build -t envinfo/test .
	out="$(docker run envinfo/test 2>&1)"
	value=$(echo -n "$out" |grep -w "$search" |cut -d ':' -f 2)
	if [[ "${value:1}" != $expected_value* ]]; then
		exit 1;
	fi
	popd
done

echo "Test finished successfully!"
