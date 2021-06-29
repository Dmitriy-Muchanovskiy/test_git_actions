#/bin/bash

cover=$(go test -cover | grep -o " .*%" | cut -c 2,3)
if [ $cover -lt 50 ]; then exit 1;fi
