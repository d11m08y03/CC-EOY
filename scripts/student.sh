curl -X POST http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJlbWFpbCI6ImpvaG4uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyOTA5NTY4fQ.zXVLq2WCk1Z8N7YEVrarCjrZACe98y1ru_d_xCIduCs" \
-d '{
    "student_number": "S1234567"
}'
