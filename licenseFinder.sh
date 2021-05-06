#!/bin/bash
# license_finder report --format json --columns=name version licenses approved --enabled-package-managers gomodules > license.json
# sed -i '1d' license.json
# variables
host="https://secops-core.multicloud-ibm.com"
token="jBmBPGgtilvJ76-GfBgZgCBXr1lnS8xbdeMScUSpMHMpXVMprZR2nbCn8-5zH9zx"
allowed=("MIT" "ISC")
denied=("New BSD" "\"Apache 2.0,MIT\"")
branch=$(echo "\"$TRAVIS_BRANCH\"")
repo=$(echo "\"$TRAVIS_REPO_SLUG\"")
# TRAVIS_COMMIT=12344
postToDevopsIntelligence() {
#   echo $branch "\"$branch\"" "\"$repo\""s
#   echo ${all[@]}
  CODE=$(curl --location --request POST -sSL -w '%{http_code}' ''"$1"'/dash/api/dev_secops/v1/services/newTestLicense/licenses?scannedBy=license_finder' \
  --header 'Authorization: Token '"$2"'' \
  --header 'Content-Type: application/json' \
  --data-raw "${all[@]}" -k)
    if [[ "$CODE" == *"200"* ]]; then
    # server return 2xx response
        echo "Successfully Posted Open Source License Scan Data to IBM DevOps Intelligence :: " $CODE
    else
        echo "Error While Posting data to IBM DevOps Intelligence :: $CODE"
    fi
}
# list=$(license_finder --prepare-no-fail --format=json --no-recursive --no-debug | awk '!/^[A-Z]/' | jq '.dependencies[].licenses[]')
listcsv=$(license_finder --prepare-no-fail --format=csv --no-recursive --no-debug | awk '!/^[A-Z]/' > ll.csv)
arr=()
while IFS="," read -r rec_column1 rec_column2 rec_column3 
do
  arr+=("$rec_column3")
done < <(tail -n +2 ll.csv)
# echo "Done"
# echo "Allowed Licenses : "${allowed[@]}
# echo "Denied Denied : "${denied[@]}
# uniques=($(for v in "${arr[@]}"; do echo "'$v'";done| sort| uniq| xargs))
IFS=$'\n' uniques=(`printf "%s\n" "${arr[@]}" |sort -u`)
uniques=($(printf "%s\n" "${arr[@]}" | sort -u))
echo "Done" 
for v in "${uniques[@]}";do
  echo "$v"
done
echo "------------------"
# --arg ad "$(date --utc +%FT%T.%3NZ)" \
all=('[]')
for i in "${uniques[@]}"
do

  for a in "${allowed[@]}" 
  do
    if [ "$i" == "$a" ]; then
        payload=$( jq -n \
                  --arg ln "$i" \
                  --arg ad "$(date --utc +%FT%T.%3NZ)" \
                  --arg st "allowed" \
                  '{license_name: $ln, analysis_date: $ad, status: $st}' )
        all=$(echo $all | jq ".+=[$payload]")
        continue
    fi
  done
done
for i in "${uniques[@]}"
do
  for b in "${denied[@]}" 
  do
    if [ "$i" == "$b" ]; then
        payload=$( jq -n \
                  --arg ln "$i" \
                  --arg ad "$(date --utc +%FT%T.%3NZ)" \
                  --arg st "denied" \
                  '{license_name: $ln, analysis_date: $ad, status: $st}' )
        all=$(echo $all | jq ".+=[$payload]")
        continue
    fi
  done
done
# echo payload
echo ${all[@]}
postToDevopsIntelligence $host $token $all[@]