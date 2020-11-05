source .env
curl localhost:8080 |jq '.[].Title'