//
const express = require("express");
const router = express.Router();
router.get("/uploadRecord", (req, res) => {
  res.json({ msg: "login works" });
});
router.post("/uploadRecord", (req, res) => {
  var patientNo = req.body.patientNo;
  var patientByte = req.body.patientByte;
  res.json({
    name: name,
    pwd: pwd,
    status: "sucess",
    dahua: "你好"
  });
});
module.exports = router;