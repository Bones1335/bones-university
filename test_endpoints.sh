#!/bin/bash

curl -X POST http://localhost:8080/admin/reset

echo "Create 'unset' role"
CREATE_UNSET_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"unset"}')

echo $CREATE_UNSET_ROLE | jq .

echo "Create 'John Snow' user"
CREATE_JSNOW_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Snow","enrollment_first_name":"John","enrollment_personal_email":"john.snow83@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_JSNOW_USER | jq .

jsnowID=$(echo $CREATE_JSNOW_USER | jq -r .users_id)

echo "Create 'administrator' role"
CREATE_ADMIN_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"administrator"}')

echo $CREATE_ADMIN_ROLE | jq .

echo "Create 'student' role"
CREATE_STUDENT_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"student"}')

echo $CREATE_STUDENT_ROLE | jq .

echo "Create 'professor' role"
CREATE_PROFESSOR_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"professor"}')

echo $CREATE_PROFESSOR_ROLE | jq .

echo "Login 'John Snow'"
LOGIN_JSNOW=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"jsnow","login_password":"0123456"}')

echo $LOGIN_JSNOW | jq .

jsToken=$(echo $LOGIN_JSNOW | jq -r .token)

echo "Update 'John Snow'"
UPDATE_JSNOW=$(curl -X PUT "http://localhost:8080/api/users/$jsnowID" -H "Authorization: Bearer $jsToken" -d "{\"users_id\":\"$jsnowID\",\"last_name\":\"Snow\",\"first_name\":\"John\",\"personal_email\":\"john.snow82@gmail.com\",\"password\":\"0123456\"}")

echo $UPDATE_JSNOW | jq .

echo "Create 'John Doe' user"
CREATE_JDOE_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Doe","enrollment_first_name":"John","enrollment_personal_email":"john.doe23@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_JDOE_USER | jq .

jdoeID=$(echo $CREATE_JDOE_USER | jq -r .users_id)

sudo -i -u postgres psql -d bones_university -v jsnow_id="$jsnowID" <<'EOF'
UPDATE users_roles 
SET role_id = (SELECT roles_id FROM roles WHERE role_name = 'administrator')
WHERE user_id = :'jsnow_id'
RETURNING user_id, role_id;
EOF

echo "Update 'John Doe' role"
UPDATE_JDOE_ROLE=$(curl -X PUT "http://localhost:8080/admin/users_roles/$jdoeID" -H "Authorization: Bearer $jsToken" -d "{\"role_name\":\"student\",\"users_id\":\"$jdoeID\"}")

echo $UPDATE_JDOE_ROLE | jq .

echo "Login 'John Doe'"
LOGIN_JDOE=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"jdoe","login_password":"0123456"}')

echo $LOGIN_JDOE | jq .

jdToken=$(echo $LOGIN_JDOE | jq -r .token)

echo "Get 'John Doe' information"
GET_JDOE_INFO=$(curl -X GET "http://localhost:8080/api/users/$jdoeID" -H "Authorization: Bearer $jdToken")

echo $GET_JDOE_INFO | jq .

echo "Create 'physical therapy degree' as John Snow"
CREATE_DEGREE_JSNOW=$(curl -X POST http://localhost:8080/admin/degrees -H "Authorization: Bearer $jsToken" -H "Content-Type:application/json" -d '{"degree_name":"Physical Therapy","degree_level":"Masters","degree_department":"Rehabilitation Department","degree_duration":4}')

echo $CREATE_DEGREE_JSNOW | jq .

physicalTherapy=$(echo $CREATE_DEGREE_JSNOW | jq -r .degrees_id)

echo "Create 'speech language therapy degree' as John Doe (and fail)"
CREATE_DEGREE_JDOE=$(curl -X POST http://localhost:8080/admin/degrees -H "Authorization: Bearer $jdToken" -d '{"degree_name":"Speech Language Therapy","degree_level":"Masters","degree_department":"Rehabilitation Department","degree_duration":3}')

echo $CREATE_DEGREE_JDOE | jq .

echo "Create 'french and francophone studies degree' as John Snow"
CREATE_DEGREE_JSNOW=$(curl -X POST http://localhost:8080/admin/degrees -H "Authorization: Bearer $jsToken" -d '{"degree_name":"French and Francophone Studies","degree_level":"Masters","degree_department":"French and Italian Department","degree_duration":2}')

echo $CREATE_DEGREE_JSNOW | jq .

echo "Get 'all degrees'"
GET_DEGREES=$(curl http://localhost:8080/api/degrees -H "Content-Type:application/json")

echo $GET_DEGREES | jq .

echo "Create 'academic year 1' as John Snow"
CREATE_YEAR_1_JSNOW=$(curl -X POST http://localhost:8080/admin/years -H "Authorization: Bearer $jsToken" -d '{"year_in_school":1}')

echo $CREATE_YEAR_1_JSNOW | jq .

startingYear=$(echo $CREATE_YEAR_1_JSNOW | jq -r .years_id)

echo "Create 'academic year 2' as John Snow"
CREATE_YEAR_2_JSNOW=$(curl -X POST http://localhost:8080/admin/years -H "Authorization: Bearer $jsToken" -d '{"year_in_school":2}')

echo $CREATE_YEAR_2_JSNOW | jq .

echo "Create 'academic year 3' as John Snow"
CREATE_YEAR_3_JSNOW=$(curl -X POST http://localhost:8080/admin/years -H "Authorization: Bearer $jsToken" -d '{"year_in_school":3}')

echo $CREATE_YEAR_3_JSNOW | jq .

echo "Create 'academic year 4' as John Snow"
CREATE_YEAR_4_JSNOW=$(curl -X POST http://localhost:8080/admin/years -H "Authorization: Bearer $jsToken" -d '{"year_in_school":4}')

echo $CREATE_YEAR_4_JSNOW | jq .

echo "Create 'academic year 5' as John Snow"
CREATE_YEAR_5_JSNOW=$(curl -X POST http://localhost:8080/admin/years -H "Authorization: Bearer $jsToken" -d '{"year_in_school":5}')

echo $CREATE_YEAR_5_JSNOW | jq .

echo "Enroll John Doe in 'physical therapy degree program'"
ENROLL_JDOE=$(curl -X POST http://localhost:8080/api/students_programs -H "Content-Type:appplication/json" -d "{\"cohort_year\":2025,\"start_date\":\"2025-09-01T00:00:00Z\",\"student_id\":\"$jdoeID\",\"program_id\":\"$physicalTherapy\",\"academic_year_id\":\"$startingYear\"}")

echo $ENROLL_JDOE | jq .

echo "Create 'Robert Jordan' user"
CREATE_RJORDAN_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Jordan","enrollment_first_name":"Robert","enrollment_personal_email":"jordan.robert@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_RJORDAN_USER | jq .

rjordanID=$(echo $CREATE_RJORDAN_USER | jq -r .users_id)

echo "Update 'Robert Jordan' role as John Snow"
UPDATE_RJORDAN_ROLE=$(curl -X PUT "http://localhost:8080/admin/users_roles/$rjordanID" -H "Authorization: Bearer $jsToken" -d "{\"role_name\":\"professor\",\"users_id\":\"$rjordanID\"}")

echo $UPDATE_RJORDAN_ROLE | jq .

echo "Login 'Robert Jordan'"
LOGIN_RJORDAN=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"rjorda","login_password":"0123456"}')

echo $LOGIN_RJORDAN | jq .

rjToken=$(echo $LOGIN_RJORDAN | jq -r .token)

echo "Create 'anatomy course' as Robert Jordan"
CREATE_COURSE_RJORDAN=$(curl -X POST http://localhost:8080/api/courses -H "Authorization: Bearer $rjToken" -d "{\"course_code\":\"UE1\",\"course_name\":\"Spinal Anatomy\",\"course_description\":\"This course will go over the anatomy of the spine.\",\"course_professor_id\":\"$rjordanID\"}")

echo $CREATE_COURSE_RJORDAN | jq .

spinalCourse=$(echo $CREATE_COURSE_RJORDAN | jq -r .courses_id)

echo "Link 'anatomy course' to 'physical therapy degree'as Robert Jordan"
CREATE_DEGREES_COURSES_RJORDAN=$(curl -X POST http://localhost:8080/api/degrees_courses -H "Authorization: Bearer $rjToken" -d "{\"degree_id\":\"$physicalTherapy\",\"course_id\":\"$spinalCourse\"}")

echo $CREATE_DEGREES_COURSES_RJORDAN | jq .

echo "Create 'anatomy assignment' as Robert Jordan"
CREATE_ASSIGNMENT_RJORDAN=$(curl -X POST http://localhost:8080/api/assignments -H "Authorization: Bearer $rjToken" -d "{\"assignment_name\":\"Spinal Cord Labeling\",\"assignment_due_date\":\"2025-10-31T23:59:00Z\",\"assignment_description\":\"This assignment asks you to label each anatomical component of the spinal cord. Use the attached document that's a picture of the spine with the blank lines you need to fill in with each element. You have until midnight on Halloween to submit your completed worksheet.\"}")

echo $CREATE_ASSIGNMENT_RJORDAN | jq .
