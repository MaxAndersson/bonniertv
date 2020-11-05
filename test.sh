source .env
curl -H "X-Secret: 1234" localhost:8080  |jq '.[].Title'