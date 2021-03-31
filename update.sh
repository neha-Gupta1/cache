#!/bin/bash

cyclonedx-go -o bom.xml

cat > payload.json <<__HERE__
{
  "projectName":"test1",
  "projectVersion":"1.1" ,
  "bom": "$(cat bom.xml |base64 -)"
}
__HERE__

curl -X "PUT" http://19f8754685e8.ngrok.io/api/v1/bom \
     -H 'Content-Type: application/json' \
     -H 'X-API-Key: smbJt0FTNzDTxB8jKXWeJODniwZAHE6w' \
     -d @payload.json