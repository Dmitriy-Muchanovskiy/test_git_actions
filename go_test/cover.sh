#/bin/bash

cover=$(go test -cover | grep -o " .*%" | cut -c 2,3)
if [ $cover -gt 0 ]
then 
    echo "coverage: $cover% of statements"
else 
    echo "coverage 0% exit job" 
    exit 1
fi
