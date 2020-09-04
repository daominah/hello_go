export ROOT_CA=rootCA.crt

set -x

openssl req -x509 -new -nodes -sha256 -days 10000 \
    -subj "/C=VN/O=DaominahTrustServices/CN=daominah" \
    -key ${ROOT_CA_KEY} -passin pass:${ROOT_CA_KEY_PASS_PHRASE} \
    -out ${ROOT_CA}

set +x
