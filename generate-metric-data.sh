#!/bin/bash

echo "\n This script will be running in a infinite loop. To stop it use Ctrl+C"
echo "\n"

while :
do
    printf "."
    for i in {1 .. 10}
    do
        curl --silent -X POST \
          http://127.0.0.1:8080/user-event/ \
          -H 'Cache-Control: no-cache' \
          -H 'Content-Type: application/json' \
          -H 'Postman-Token: 86b5bbf9-01bc-454c-b4ea-a12159ca36b2' \
          -d "{ \
            \"apiKey\":\"TEST_API_KEY\",
            \"userId\":2,
            \"timestamp\": \"2018-07-22T05:41:10.003+06:00\"
        }" > /dev/null &
    done
	sleep 1
done