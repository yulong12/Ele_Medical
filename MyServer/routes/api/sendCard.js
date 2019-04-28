//
const express = require("express");
const router = express.Router();
var sendCard = require("./../myhfc/myhfcInvoke");
var util = require("./../utils/cryptoUtils");
router.get("/sendCard", (req, res) => {
  res.json({ msg: "sendCard works" });
});
router.post("/sendCard", (req, res) => {
  var idCard = req.body.idCard;
  var sex = req.body.sex;
  var age = req.body.age;
  var address = req.body.address;
  var hisNo = req.body.hisNo;

  util.Generator(
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/MyServer/routes/pem/public.pem",
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/MyServer/routes/pem/private.pem",
    function(publickey, privatekey) {
      console.log("========publickey========" + publickey);
      console.log("=========privatekey=======" + privatekey);
      console.log("========idCard========" + idCard);
      console.log("=========sex=======" + sex);
      console.log("========age========" + age);
      console.log("=========address=======" + address);
      console.log("=========hisNo=======" + hisNo);
      util.MD5Hash(idCard, function(idCardH) {
        var requestJson = {
          fcn: "makeCard",
          args: [idCardH, sex, age, address, publickey, hisNo]
        };

        sendCard(requestJson, function(str) {
          console.log("=====str==========" + str.status);
          console.log("=====payload==========" + str.payload);

          res.send({
            detail: str
          });
        });
      });
    }
  );
});
module.exports = router;
