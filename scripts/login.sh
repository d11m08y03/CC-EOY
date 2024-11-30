#!/bin/bash

curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "john.doe@example.com",
    "password": "password123"
}'
