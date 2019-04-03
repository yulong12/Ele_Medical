#使用 openssl 生成用户证书

### 1 生成私钥

```

genrsa -out rsa_xiaoming_private_key.pem 1024
```

### 2 生成公钥

```

rsa -in rsa_xiaoming_private_key.pem -pubout -out rsa_public_key.pem
```
