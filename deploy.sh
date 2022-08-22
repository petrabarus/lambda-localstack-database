#!/usr/bin/env bash
echo "Bundling Function"
cd lambda
GOARCH=amd64 GOOS=linux go build -o main main.go
zip -r ../lambda.zip main
cd ..

echo "Creating Function"
aws lambda create-function \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --role arn:aws:iam::000000000000:role/lambda \
    --handler main --runtime go1.x \
    --zip-file fileb://lambda.zip
    
echo "Invoking Function"
aws lambda invoke \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --payload fileb://payload.json \
    output.json