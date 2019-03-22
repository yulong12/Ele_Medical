//建立住院档案
//
const express = require("express");
const router = express.Router();

router.get("/queryHospitalized", (req, res) => {
  res.json({ msg: "login works" });
});
router.post("/queryHospitalized", (req, res) => {
  var queryHospitalizedata = require("./../myhfc/myhfcInvoke");
  var idcard = req.body.idcard;

  var requestJson = {
    fcn: "queryHospitalized",
    args: [idcard]
  };
  queryHospitalizedata(requestJson, function(str) {
    console.log("=====str==========" + str.status);
    console.log("=====payload==========" + str.payload);

    res.send({
      status: "OK",
      detail: str
    });
  });
});
module.exports = router;
