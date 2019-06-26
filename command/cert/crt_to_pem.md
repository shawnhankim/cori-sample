# Convert CRT to PEM
openssl x509 -inform DER -in yourdownloaded.crt -out outcert.pem -text

# Testing The CA
openssl s_client -connect foo.whatever.com:443 -CApath /etc/ssl/certs

