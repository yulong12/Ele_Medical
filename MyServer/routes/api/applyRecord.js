//
const express = require("express");
const router = express.Router();
router.get("/applyRecord", (req, res) => {
  res.json({ msg: "login works" });
});
router.post("/applyRecord", (req, res) => {
  var hisNo = req.body.hisNo;
  var doctorNo = req.body.doctorNo;
  var targetNo = req.body.targetNo;
  var redicalNO = req.body.redicalNO;
  var applier = req.body.applier;
  var applierNo = req.body.applierNo;
  var applierHisNo = req.body.applierHisNo;
  var requestJson = {
    hisNo: hisNo, //医院编号
    doctorNo: doctorNo, //医生编号
    targetNo: targetNo, //请求医院编号
    redicalNO: redicalNO, //请求病历编号
    applier: applier, //请求人
    applierNo: applierNo, //请求人编号
    applierHisNo: applierHisNo //请求人所在医院编号
  };
  res.json(requestJson);
});
module.exports = router;
