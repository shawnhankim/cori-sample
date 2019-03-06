Check a certificate
===================
Check a certificate and return information about it (signing authority, expiration date, etc.):
$ openssl x509 -in file.pem -text -noout


Check a certificate's end date
==============================
$ openssl x509 -enddate -noout -in file.pem


Check a certificate's start date
================================
$ openssl x509 -startdate -noout -in file.pem
notBefore=Mar  5 19:06:23 2019 GMT


Check a CSR
===================
Verify the CSR and print CSR data filled in when generating the CSR:
$ openssl req -text -noout -verify -in server.csr




