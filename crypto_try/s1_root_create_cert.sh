# create public key from generated key in s0

set -x
set -e

export ROOT_CA_KEY=rootCA.key # defined in s0
export ROOT_CA_KEY_PASS_PHRASE=123qwe # defined in s0
export ROOT_CA=rootCA.crt

openssl req -x509 -new -nodes -sha256 -days 10000 \
    -subj "/C=VN/O=DaominahTrustServices/CN=daominah" \
    -key ${ROOT_CA_KEY} -passin pass:${ROOT_CA_KEY_PASS_PHRASE} \
    -out ${ROOT_CA}

set +x

echo "create file $(ls -l | grep ${ROOT_CA})"
