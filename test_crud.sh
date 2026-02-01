#!/bin/bash

BASE_URL="http://localhost:8080/api/produk"

echo "1. Testing Create Product..."
CREATE_RES=$(curl -s -X POST $BASE_URL \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Product", "price": 10000, "stock": 50, "category_id": 1}')
echo "Response: $CREATE_RES"

ID=$(echo $CREATE_RES | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
if [ -z "$ID" ]; then
  echo "Failed to create product or get ID"
  exit 1
fi
echo "Created Product ID: $ID"
echo ""

echo "2. Testing Get All Products..."
curl -s -X GET $BASE_URL
echo ""
echo ""

echo "3. Testing Get Product By ID ($ID)..."
curl -s -X GET "$BASE_URL/$ID"
echo ""
echo ""

echo "4. Testing Update Product ($ID)..."
curl -s -X PUT "$BASE_URL/$ID" \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Product", "price": 15000, "stock": 40, "category_id": 2}'
echo ""
echo ""

echo "5. Testing Get Product By ID ($ID) to verify update..."
curl -s -X GET "$BASE_URL/$ID"
echo ""
echo ""

echo "6. Testing Delete Product ($ID)..."
curl -s -X DELETE "$BASE_URL/$ID"
echo ""
echo ""

echo "7. Verification: Get Product By ID ($ID) should fail..."
curl -s -X GET "$BASE_URL/$ID"
echo ""
