#!/bin/bash
mkdir -p /root/local/bin/
wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
chmod +x cfssl_linux-amd64
wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
chmod +x cfssljson_linux-amd64
wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
chmod +x cfssl-certinfo_linux-amd64
mv cfssl_linux-amd64 /root/local/bin/cfssl/cfssl
mv cfssl-certinfo_linux-amd64 /root/local/bin/cfssl-certinfo
mv cfssljson_linux-amd64 /root/local/bin/cfssljson
export PATH=/root/local/bin:$PATH
# export PATH=/home/liang/local/bin:$PATH