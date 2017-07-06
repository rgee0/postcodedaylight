docker service rm postcodedaylight

curl localhost:8080/system/functions -d '
{"service": "postcodedaylight", "image": "rgee0/postcodedaylight", "envProcess": "/go/bin/postcodedaylight", "network": "func_functions"}'