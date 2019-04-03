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
        if (key.verify(dataStr, signStr, "bash64")) {
          callback("true");
        }
      }
    });
  }
};
