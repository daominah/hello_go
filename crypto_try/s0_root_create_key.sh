export ROOT_CA_KEY=rootCA.key
# if you want a non password protected key, just let this var empty
export ROOT_CA_KEY_PASS_PHRASE=123qwe

set -x

if [[ -n "${ROOT_CA_KEY_PASS_PHRASE}" ]]; then
    encryptOption="-aes256 -passout pass:${ROOT_CA_KEY_PASS_PHRASE}"
fi

openssl genrsa ${encryptOption} -out ${ROOT_CA_KEY} 4096

set +x

# source: https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309
