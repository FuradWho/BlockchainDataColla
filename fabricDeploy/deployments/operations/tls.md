openssl req -new -x509 -key ../keystore/client.key -out client.crt -days 3650

openssl genrsa -out client.key 2048

```
vim /etc/ssl/openssl.cnf
编辑openssl.cnf,在[v3_ca]下面添加：subjectAltName = IP:域名|IP地址
[ v3_ca ]
subjectAltName = IP:172.10.15.110

否则将会报错：
x509: cannot validate certificate for <ipaddress> because it doesn't contain any IP SANs
```

protoc --micro_out=. --go_out=. crs.proto 

openssl req -x509 -new -nodes -key ../keystore/ca.key -subj "/CN=colla.ca.com" -days 5000 -out ca.pem