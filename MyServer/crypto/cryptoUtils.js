var crypto = require("crypto");
const padding = crypto.constants.RSA_PKCS1_PADDING;
const NodeRSA = require("node-rsa");
const fs = require("fs");

module.exports = {
  //data的参数类型为buffer
  MD5Hash: function(data, callback) {
    //计算该文件的md5值
    var fsHash = crypto.createHash("md5");
    fsHash.update(data);
    var md5 = fsHash.digest("hex");
    console.log("文件的MD5是：%s", md5);
    callback(md5);
  },
  //使用私钥解密
  Decrypt: function(privateKeyPath, EnDataStr, callback) {
    fs.exists(privateKeyPath, function(exists) {
      if (exists) {
        var pem = fs.readFileSync(privateKeyPath, "utf8");
        var key = new NodeRSA(pem);
        var decrypted = key.decrypt(EnDataStr, "utf8");
        console.log("decrypted:" + decrypted);
        callback(decrypted);
      }
    });
  },

  //使用公钥加密
  Encrypt: function(publicKeyPath, dataStr, callback) {
    fs.exists(publicKeyPath, function(exists) {
      if (exists) {
        var pem = fs.readFileSync(publicKeyPath, "utf8");
        var key = new NodeRSA(pem);
        encrypted = key.encrypt(dataStr, "base64");
        console.log("encrypted:" + encrypted);

        callback(encrypted);
      }
    });
  },
  //生成公钥和私钥
  Generator: function generator(publicPemPath, privatePemPath) {
    // const publicKeyPath = "./pem/public.pem";
    // const privateKeyPath = "./pem/private.pem";
    var key = new NodeRSA({ b: 512 });
    key.setOptions({ encryptionScheme: "pkcs1" });

    const publicKey = key.exportKey("public");
    fs.writeFile(publicPemPath, publicKey, err => {
      if (err) throw err;
      console.log("公钥已保存");
    });
    const privateKey = key.exportKey("private");
    fs.writeFile(privatePemPath, privateKey, err => {
      if (err) throw err;
      console.log("私钥已保存");
    });
  },
  //使用私钥签名
  SignData: function SignData(privateKeyPath, dataStr, callback) {
    fs.exists(privateKeyPath, function(exists) {
      if (exists) {
        var pem = fs.readFileSync(privateKeyPath, "utf8");
        var key = new NodeRSA(pem);
        signOut = key.sign(dataStr, "base64");
        console.log("=====signOut=======" + signOut);
        callback(signOut);
      }
    });
  },
  //验证签名
  VerySign: function VerySign(publicKeyPath, dataStr, signStr, callback) {
    fs.exists(publicKeyPath, function(exists) {
      if (exists) {
        var pem = fs.readFileSync(publicKeyPath, "utf8");
        var key = new NodeRSA(pem);
        var verifySign = key.verify(
          Buffer.from(dataStr),
          signStr,
          "base64",
          "base64"
        );
        console.log("------verifySign------" + verifySign);
        callback(verifySign);
      }
    });
  },

  //  attrs = attrs || [{
  //     name: 'commonName',
  //     value: 'example.org'
  //   }, {
  //     name: 'countryName',
  //     value: 'US'
  //   }, {
  //     shortName: 'ST',
  //     value: 'Virginia'
  //   }, {
  //     name: 'localityName',
  //     value: 'Blacksburg'
  //   }, {
  //     name: 'organizationName',
  //     value: 'Test'
  //   }, {
  //     shortName: 'OU',
  //     value: 'Test'
  //   }];

  SelfSign: function SelfSign(attrs, savePath) {
    var selfsigned = require("selfsigned");
    // var attrs = [
    //   { name: "commonName", value: "contoso.com", card: "130521199203080776" }
    // ];
    var pems = selfsigned.generate(attrs, { days: 365 });
    console.log(pems);
    var privateCert = pems.private;
    var publicCert = pems.public;
    var cert = pems.cert;
    this.SaveData(privateCert, "private.pem", savePath);
    this.SaveData(publicCert, "public.pem", savePath);
    this.SaveData(cert, "cert.pem", savePath);
    console.log("----private-------" + pems.private);
  },
  SaveData: function SaveData(data, fileName, filePath) {
    fs.writeFile(
      filePath + "/" + fileName,
      data,
      { flag: "w", encoding: "utf-8", mode: "0666" },
      function(err) {
        if (err) {
          console.log("文件写入失败");
        } else {
          console.log("文件写入成功");
        }
      }
    );
  }
};
