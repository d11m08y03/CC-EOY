#!/bin/bash
curl -X POST http://102.222.106.153:8080/register \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "password123"
}'
