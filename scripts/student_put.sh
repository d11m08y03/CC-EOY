curl -X PUT http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmdhbmlzb3JfaWQiOjIsImVtYWlsIjoiam9obi5kb2VAZXhhbXBsZS5jb20iLCJpc19hZG1pbiI6dHJ1ZSwiZXhwIjoxNzMzOTI5Nzg4fQ.djZmQrb-KIzY04XWiq6SvKXv2zlGjd3oGOs1gNeCTCs" \
-d '{ 
    "student_id": "2315007"
}'
