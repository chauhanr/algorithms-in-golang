#TLS and HTTPS 

The TLS (Transport Layer Security) wich enables the HTTPS protocol uses asymmetric certificate
encryption. We need to generate the certificate and key pem files that can be used with a golang TLS
server. 

We can generate a self signed certificate for our application using the go utility under the
crypto/tls/ package 

`go run $GOROOT/src/crypto/tls/generate_cert.go --host=localhost`  this will generate the keys and
cert pem files for the localhost usage. 

 
