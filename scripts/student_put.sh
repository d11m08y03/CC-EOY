curl -X PUT http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmdhbmlzb3JfaWQiOjEsImVtYWlsIjoiam9obi5kb2VAZXhhbXBsZS5jb20iLCJleHAiOjE3MzMzMjUzNzR9.tHwe5hzeszIPzU29Zps3SIg3gDcxtMsPnYc4zJZkBBM" \
-d '{ 
    "student_id": "2314543"
}'
