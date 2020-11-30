#!/usr/bin/env bash

export privatePath=myorg0.key
export publicPath=myorg0.crt

openssl rsa --noout --modulus --in ${privatePath} | openssl md5
openssl x509 --noout --modulus --in ${publicPath} | openssl md5
