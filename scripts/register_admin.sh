#!/bin/bash
curl -X POST http://localhost:8080/create-admin \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "password123"
}'
