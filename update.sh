#!/bin/bash

cyclonedx-go -o bom.xml

cat > payload.json <<__HERE__
{
  "projectName":"cache- golang",
  "projectVersion":"v1" ,
  "bom": "$(cat bom.xml |base64 -)"
}
__HERE__

curl -X "PUT" http://164127935d6f.ngrok.io/api/v1/bom \
     -H 'Content-Type: application/json' \
     -H 'X-API-Key: smbJt0FTNzDTxB8jKXWeJODniwZAHE6w' \
     -d @payload.json