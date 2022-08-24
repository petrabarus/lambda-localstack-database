#!/usr/bin/env bash
echo "Bundling Function"
cd lambda
go get ./...
GOARCH=amd64 GOOS=linux go build -o ../dist/main main.go
cd ..
cp -r migrations dist
cd dist
zip -rv ../lambda.zip main migrations
cd ..

echo "Creating Function"
aws lambda create-function \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --timeout 10 \
    --role arn:aws:iam::000000000000:role/lambda \
    --handler main --runtime go1.x \
    --zip-file fileb://lambda.zip
    
echo "Invoking Function"
aws lambda invoke \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --payload fileb://payload.json \
    output.json