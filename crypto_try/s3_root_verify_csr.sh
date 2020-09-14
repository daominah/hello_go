# root CA verifies CSRs, result are certificates

set -x
set -e

export ROOT_CA_KEY=rootCA.key # defined in s0
export ROOT_CA_KEY_PASS_PHRASE=123qwe # defined in s0
export ROOT_CA=rootCA.crt # defined in s1
export ORG0_CSR=myorg0.csr # defined in s2
export ORG1_CSR=myorg1.csr # defined in s2

csrs=(${ORG0_CSR} ${ORG1_CSR})
crts=(myorg0.crt myorg1.crt)

for i in ${!csrs[@]}; do
    openssl x509 -req -sha256 -days 10000 \
        -in ${csrs[i]} -CAcreateserial -CA ${ROOT_CA} \
        -CAkey ${ROOT_CA_KEY} -passin pass:${ROOT_CA_KEY_PASS_PHRASE} \
        -out ${crts[i]}
done

set +x

echo "created files:"
echo "$(ls -l | grep crt | grep -v ${ROOT_CA})"
