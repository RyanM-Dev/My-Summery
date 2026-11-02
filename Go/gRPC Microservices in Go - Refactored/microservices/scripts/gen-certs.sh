#!/usr/bin/env bash
# Creates a development CA plus a Payment server certificate and an Order
# client certificate for mTLS (Chapter 6). Output: ./certs
set -euo pipefail

OUT="${1:-certs}"
DAYS=60
SUBJ="/C=TR/ST=EURASIA/L=ISTANBUL/O=Microservices"
mkdir -p "$OUT"
cd "$OUT"

# 1. CA: private key + self-signed certificate
openssl req -x509 -sha256 -newkey rsa:4096 -nodes -days 365 \
  -keyout ca-key.pem -out ca-cert.pem \
  -subj "$SUBJ/OU=CA/CN=microservices-dev-ca"

# 2. Payment server certificate. The SAN lists every name clients dial.
openssl req -newkey rsa:4096 -nodes -sha256 \
  -keyout server-key.pem -out server-req.pem \
  -subj "$SUBJ/OU=PaymentService/CN=payment"
echo "subjectAltName=DNS:payment,DNS:localhost,DNS:*.microservices.dev,IP:127.0.0.1" > server-ext.cnf
openssl x509 -req -in server-req.pem -days "$DAYS" -sha256 \
  -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial \
  -out server-cert.pem -extfile server-ext.cnf

# 3. Order client certificate
openssl req -newkey rsa:4096 -nodes -sha256 \
  -keyout client-key.pem -out client-req.pem \
  -subj "$SUBJ/OU=OrderService/CN=order"
echo "extendedKeyUsage=clientAuth" > client-ext.cnf
openssl x509 -req -in client-req.pem -days "$DAYS" -sha256 \
  -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial \
  -out client-cert.pem -extfile client-ext.cnf

rm -f ./*-req.pem ./*.cnf ca-cert.srl
# The containers run as a non-root user, so the dev keys must be readable.
# Never do this with real keys.
chmod 644 ./*.pem
echo "certificates written to $(pwd)"
