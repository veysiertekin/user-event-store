#!/bin/bash

echo "\n This script will be running in an infinite loop. To stop it use Ctrl+C"

while :
do
    printf "."
    for i in {1 .. 10}
    do
        #
        # Sample payload:
        # {
        #	"apiKey":"TEST_API_KEY",
        #	"userId":2,
        #	"timestamp": "2018-07-22T05:41:10.003+06:00"
        # }

        datetime=$(python -c "from datetime import datetime; print(datetime.now().isoformat()+\"+00:00\")" )

        curl --silent -X POST \
          http://127.0.0.1:8080/user-event/ \
          -H 'Cache-Control: no-cache' \
          -H 'Content-Type: application/json' \
          -d "{ \
            \"apiKey\":\"TEST_API_KEY\",
            \"userId\":2,
            \"timestamp\": \"$datetime\"
        }" > /dev/null &
    done
	sleep 1
done
