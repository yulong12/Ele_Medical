//
const express = require("express");
const router = express.Router();
var applyRecord = require("./../myhfc/myhfcInvoke");
router.get("/applyRecord", (req, res) => {
  res.json({ msg: "login works" });
});

router.post("/applyRecord", (req, res) => {
  // TargetNo
  // RedicalNO
  // Applier
  // ApplierNo
  // ApplierHisNo

  var targetNo = req.body.targetNo;
  var redicalNO = req.body.redicalNO;
  var applier = req.body.applier;
  var applierNo = req.body.applierNo;
  var applierHisNo = req.body.applierHisNo;
  var requestJson = {
    fnc: "applyRemoteData",
    args: [
      targetNo, //请求医院编号
      redicalNO, //请求病历编号
      applier, //请求人
      applierNo, //请求人编号
      applierHisNo //请求人所在医院编号
    ]
  };
  applyRecord(requestJson, function(str) {
    console.log("=====str==========" + str.status);
    console.log("=====payload==========" + str.payload);

    res.send({
      status: "OK",
      detail: str
    });
  });
});
// res.json(requestJson);

module.exports = router;
