#!/usr/bin/env bash
echo "Bundling Function"
cd lambda
zip -r ../lambda.zip *
cd ..

echo "Creating Function"
aws lambda create-function \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --runtime nodejs14.x \
    --role arn:aws:iam::000000000000:role/lambda \
    --handler index.handler \
    --zip-file fileb://lambda.zip
    
echo "Invoking Function"
aws lambda invoke \
    --endpoint-url http://localhost:4566 \
    --function-name lambda \
    --payload fileb://payload.json \
    output.json