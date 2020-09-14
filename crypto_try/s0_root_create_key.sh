# Create a RSA private key and save to file ROOT_CA_KEY (may be encrypted or not)

set -x
set -e

export ROOT_CA_KEY=rootCA.key
export ROOT_CA_KEY_PASS_PHRASE=123qwe

if [[ -n "${ROOT_CA_KEY_PASS_PHRASE}" ]]; then
    encryptOption="-aes256 -passout pass:${ROOT_CA_KEY_PASS_PHRASE}"
fi

openssl genrsa ${encryptOption} -out ${ROOT_CA_KEY} 4096

set +x

echo "created file: $(ls -l | grep ${ROOT_CA_KEY})"

# source: https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309
