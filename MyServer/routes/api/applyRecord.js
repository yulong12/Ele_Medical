//
const express = require("express");
const router = express.Router();
router.get("/applyRecord", (req, res) => {
  res.json({ msg: "login works" });
});
router.post("/applyRecord", (req, res) => {
  var targetNo = req.body.targetNo;
  var redicalNO = req.body.redicalNO;
  var applier = req.body.applier;
  var applierNo = req.body.applierNo;
  res.json({
    name: name,
    pwd: pwd,
    status: "sucess",
    dahua: "你好"
  });
});
module.exports = router;
