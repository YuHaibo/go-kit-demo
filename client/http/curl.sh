#!/usr/bin/env sh -v

#Sum Response
curl -d '{"a":11111, "b":22222}' http://127.0.0.1:8890/sum

#Concat Response
curl -d '{"a":"11111", "b":"22222"}' http://127.0.0.1:8890/concat
