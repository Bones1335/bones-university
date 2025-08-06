#!/bin/bash

curl -X POST http://localhost:8080/admin/reset

CREATE_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"last_name":"Snow","first_name":"John","personal_email":"john.snow83@gmail.com","password":"0123456"}')

echo $CREATE_USER | jq .

CREATE_ADMIN_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"administrator"}')

echo $CREATE_ADMIN_ROLE | jq .

CREATE_STUDENT_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"student"}')

echo $CREATE_STUDENT_ROLE | jq .

CREATE_PROFESSOR_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"professor"}')

echo $CREATE_PROFESSOR_ROLE | jq .

CREATE_UNSET_ROLE=$(curl -X POST http://localhost:8080/admin/roles -H "Content-Type:application/json" -d '{"role_name":"unset"}')

echo $CREATE_UNSET_ROLE | jq .
