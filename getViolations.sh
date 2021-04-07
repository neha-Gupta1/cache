#!/bin/bash

result=$(curl -X GET --header "Content-Type: application/json" -H "X-API-Key: smbJt0FTNzDTxB8jKXWeJODniwZAHE6w" "http://19f8754685e8.ngrok.io/api/v1/finding/project/cecbf7b4-076c-41e1-b4ae-69ea2e509271")
echo "Response from server"
echo $result  > vulnerabilities.json

result=$(curl -X GET --header "Content-Type: application/json" -H "X-API-Key: smbJt0FTNzDTxB8jKXWeJODniwZAHE6w" "http://19f8754685e8.ngrok.io/api/v1/violation/project/cecbf7b4-076c-41e1-b4ae-69ea2e509271")
echo "Response from server"
echo $result  > violation.json

