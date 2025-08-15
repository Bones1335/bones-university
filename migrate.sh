#!/bin/bash

cd sql/schema

goose postgres postgres://postgres:postgres@localhost:5432/bones_university?sslmode=disable down

goose postgres postgres://postgres:postgres@localhost:5432/bones_university?sslmode=disable up

cd ../..

