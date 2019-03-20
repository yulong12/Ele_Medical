var crypto = require("crypto");
var fs = require("fs");

//读取一个Buffer
var buffer = fs.readFileSync("./users.js");
var fsHash = crypto.createHash("md5");

fsHash.update(buffer);
var md5 = fsHash.digest("hex");
console.log("文件的MD5是：%s", md5);
