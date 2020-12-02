#!/usr/bin/env bash

export privatePath=generatedCAWithGo.key
export publicPath=generatedPublicKeyWithGo.pub

openssl rsa --noout --modulus --in ${privatePath}

# one of 2 following command is good, depend on cert format
openssl rsa --inform pem --noout --modulus --pubin --in ${publicPath}
