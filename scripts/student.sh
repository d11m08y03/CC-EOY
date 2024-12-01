curl -X POST http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InRlc3RAZW1haWwuY29tIiwiZXhwIjoxNzMzMTI1MDU3fQ.wAoEHTeJvmldPU5etKqz9UAgx7Uaa-lrb9QeJOH0JaA" \
-d '{ 
    "student_number": "S1234567"
}'
