# organizations create private key and CSR

set -x
set -e

# an example organization

export ORG0_KEY=myorg0.key
# create a certificate signing request. CSR is where you specify the details for
# the certificate you want to generate. This request will be processed by the
# owner of the root CA key (you in this case since you create it earlier) to
# generate the certificate. While creating the CSR, it is important to specify
# the Common Name providing the IP address or domain name for the service
export ORG0_CSR=myorg0.csr

openssl genrsa -out ${ORG0_KEY} 4096

openssl req -new -sha256 \
    -key ${ORG0_KEY} -subj "/C=VN/O=DaominahChild0/CN=127.0.0.1" \
    -out ${ORG0_CSR}

# another organization

export ORG1_KEY=myorg1.key
export ORG1_PASS_PHRASE=123qwe
export ORG1_CSR=myorg1.csr

openssl genrsa -aes256 -passout pass:${ORG1_PASS_PHRASE} -out ${ORG1_KEY} 4096

openssl req -new -sha256 \
    -key ${ORG1_KEY} -passin pass:${ORG1_PASS_PHRASE} \
    -subj "/C=VN/O=DaominahChild1/CN=127.0.0.1" \
    -out ${ORG1_CSR}

set +x

echo "created files:"
echo "$(ls -l | grep myorg | grep -v crt)"
