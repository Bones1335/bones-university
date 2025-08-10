#!/bin/bash

curl -X POST http://localhost:8080/admin/reset

CREATE_UNSET_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"unset"}')

echo $CREATE_UNSET_ROLE | jq .

CREATE_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"enrollment_last_name":"Snow","enrollment_first_name":"John","enrollment_personal_email":"john.snow83@gmail.com","enrollment_password":"0123456"}')

echo $CREATE_USER | jq .

jsnowID=$(echo $CREATE_USER | jq -r .users_id)

CREATE_ADMIN_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"administrator"}')

echo $CREATE_ADMIN_ROLE | jq .

CREATE_STUDENT_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"student"}')

echo $CREATE_STUDENT_ROLE | jq .

CREATE_PROFESSOR_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"professor"}')

echo $CREATE_PROFESSOR_ROLE | jq .

LOGIN_JSNOW=$(curl -X POST http://localhost:8080/api/login -H "Content-Type:application/json" -d '{"login_username":"jsnow","login_password":"0123456"}')

echo $LOGIN_JSNOW | jq .

token=$(echo $LOGIN_JSNOW | jq -r .token)

UPDATE_JSNOW=$(curl -X PUT "http://localhost:8080/api/users/$jsnowID" -H "Authorization: Bearer $token" -d "{\"users_id\":\"$jsnowID\",\"last_name\":\"Snow\",\"first_name\":\"John\",\"personal_email\":\"john.snow82@gmail.com\",\"password\":\"0123456\"}")

echo $UPDATE_JSNOW | jq .
