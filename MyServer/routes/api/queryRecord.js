//
const express = require("express");
const router = express.Router();
var queryRecordData = require("./../myhfc/myhfcInvoke");
router.get("/queryRecord", (req, res) => {
  res.json({ msg: "login works" });
});

router.post("/queryRecord", (req, res) => {
  // TargetNo
  // RedicalNO
  // Applier
  // ApplierNo
  // ApplierHisNo

  var patientNo = req.body.patientNo;

  var requestJson = {
    fcn: "queryRecordData",
    args: [patientNo]
  };
  queryRecordData(requestJson, function(str) {
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
