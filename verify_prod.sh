#!/bin/bash

PROD_URL="https://website-kasir-api.hv3ahr.easypanel.host/api/produk"

echo "Verifying Production Endpoint: $PROD_URL"
echo ""

# Get all products
RESPONSE=$(curl -s $PROD_URL)

echo "Response excerpt:"
echo $RESPONSE | grep -o 'category":{[^}]*}' | head -n 3

if [[ $RESPONSE == *"category"* ]]; then
    echo ""
    echo "SUCCESS: 'category' field found in response!"
else
    echo ""
    echo "FAILURE: 'category' field NOT found. Deployment might still be in progress."
fi
