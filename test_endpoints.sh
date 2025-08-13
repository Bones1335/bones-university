#!/bin/bash

curl -X POST http://localhost:8080/admin/reset

CREATE_UNSET_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"unset"}')

echo $CREATE_UNSET_ROLE | jq .

CREATE_JSNOW_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Snow","enrollment_first_name":"John","enrollment_personal_email":"john.snow83@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_JSNOW_USER | jq .

jsnowID=$(echo $CREATE_JSNOW_USER | jq -r .users_id)

CREATE_ADMIN_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"administrator"}')

echo $CREATE_ADMIN_ROLE | jq .

CREATE_STUDENT_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"student"}')

echo $CREATE_STUDENT_ROLE | jq .

CREATE_PROFESSOR_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"professor"}')

echo $CREATE_PROFESSOR_ROLE | jq .

LOGIN_JSNOW=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"jsnow","login_password":"0123456"}')

echo $LOGIN_JSNOW | jq .

jsToken=$(echo $LOGIN_JSNOW | jq -r .token)

UPDATE_JSNOW=$(curl -X PUT "http://localhost:8080/api/users/$jsnowID" -H "Authorization: Bearer $jsToken" -d "{\"users_id\":\"$jsnowID\",\"last_name\":\"Snow\",\"first_name\":\"John\",\"personal_email\":\"john.snow82@gmail.com\",\"password\":\"0123456\"}")

echo $UPDATE_JSNOW | jq .

CREATE_JDOE_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Doe","enrollment_first_name":"John","enrollment_personal_email":"john.doe23@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_JDOE_USER | jq .

jdoeID=$(echo $CREATE_JDOE_USER | jq -r .users_id)

sudo -i -u postgres psql -d bones_university -v jsnow_id="$jsnowID" <<'EOF'
UPDATE users_roles 
SET role_id = (SELECT roles_id FROM roles WHERE role_name = 'administrator')
WHERE user_id = :'jsnow_id'
RETURNING user_id, role_id;
EOF

UPDATE_JDOE_ROLE=$(curl -X PUT "http://localhost:8080/admin/users_roles/$jdoeID" -H "Authorization: Bearer $jsToken" -d "{\"role_name\":\"student\",\"users_id\":\"$jdoeID\"}")

echo $UPDATE_JDOE_ROLE | jq .

LOGIN_JDOE=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"jdoe","login_password":"0123456"}')

echo $LOGIN_JDOE | jq .

jdToken=$(echo $LOGIN_JDOE | jq -r .token)

GET_JDOE_INFO=$(curl -X GET "http://localhost:8080/api/users/$jdoeID" -H "Authorization: Bearer $jdToken")

echo $GET_JDOE_INFO | jq .

CREATE_DEGREE_JSNOW=$(curl -X POST http://localhost:8080/admin/degrees -H "Authorization: Cearer $jsToken" -d '{"degree_name":"Physical Therapy","degree_level":"Masters","degree_department":"Rehabilitation Department","degree_duration":4}')

echo $CREATE_DEGREE_JSNOW | jq .

CREATE_DEGREE_JDOE=$(curl -X POST http://localhost:8080/admin/degrees -H "Authorization: Cearer $jdToken" -d "{}")

echo $CREATE_DEGREE_JDOE | jq .
