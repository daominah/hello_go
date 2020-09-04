csrs=(${ORG0_CSR} ${ORG1_CSR})
crts=(myorg0.crt myorg1.crt)

for i in ${!csrs[@]}; do
    openssl x509 -req -sha256 -days 10000 \
        -in ${csrs[i]} -CAcreateserial -CA ${ROOT_CA} \
        -CAkey ${ROOT_CA_KEY} -passin pass:${ROOT_CA_KEY_PASS_PHRASE} \
        -out ${crts[i]}
done
