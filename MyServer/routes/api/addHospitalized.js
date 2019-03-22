//建立住院档案
//
const express = require("express");
const router = express.Router();
var addHospitalized = require("./../myhfc/myhfcInvoke");
router.get("/addHospitalized", (req, res) => {
  res.json({ msg: "login works" });
});
router.post("/addHospitalized", (req, res) => {
  var name = req.body.name;
  var age = req.body.age;
  var phone = req.body.phone;
  var idCard = req.body.idCard;
  var sex = req.body.sex;
  var address = req.body.address;
  var doctor = req.body.doctor;
  var nurse = req.body.nurse;
  var illness = req.body.illness;
  var treatment = req.body.treatment;
  var medication = req.body.medication;
  var attention = req.body.attention;
  var room = req.body.room;
  var inTime = req.body.inTime;
  var outTime = req.body.outTime;
  var cost = req.body.cost;

  var requestJson = {
    fnc: "saveHospitalized",
    args: [
      name,
      age,
      phone,
      idCard,
      sex,
      address,
      doctor,
      nurse,
      illness,
      treatment,
      medication,
      attention,
      room,
      inTime,
      outTime,
      cost
    ]
  };
  addHospitalized(requestJson, function(str) {
    console.log("=====str==========" + str.status);
    console.log("=====payload==========" + str.payload);

    res.send({
      status: "OK",
      detail: str
    });
  });
});
module.exports = router;
