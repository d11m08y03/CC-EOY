curl -X POST http://localhost:8080/auth/students \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmdhbmlzb3JfaWQiOjEsImVtYWlsIjoiam9obi5kb2VAZXhhbXBsZS5jb20iLCJleHAiOjE3MzM2NzI4NzN9.pn6OOs8Xy_JPIy7F26A8E3MvajinBE9fnhyL2EBLOg0" \
-d '{ 
    "student_id": "2ss010173222222222aa",
    "full_name": "Unga Bungaaa"
}'
