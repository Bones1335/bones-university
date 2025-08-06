#!/bin/bash

curl -X POST http://localhost:8080/admin/reset

CREATE_USER=$(curl -X POST http://localhost:8080/api/users -H "Content-Type:application/json" -d '{"last_name":"Snow","first_name":"John","personal_email":"john.snow83@gmail.com","password":"0123456"}')

echo $CREATE_USER | jq .
