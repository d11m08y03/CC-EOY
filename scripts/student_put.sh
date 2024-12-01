curl -X PUT http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImpvaG4uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMzMTYyNzk3fQ.jzyxrihbp519fDA60c0n02mT6QHHOUlvFj3kWHWkJF8" \
-d '{ 
    "student_id": "2010173"
}'
