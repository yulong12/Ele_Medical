//建立住院档案
//
const express = require("express");
const router = express.Router();
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
    name: name,
    age: age,
    phone: phone,
    idCard: idCard,
    sex: sex,
    address: address,
    doctor: doctor,
    nurse: nurse,
    illness: illness,
    treatment: treatment,
    medication: medication,
    attention: attention,
    room: room,
    inTime: inTime,
    outTime: outTime,
    cost: cost
  };
  res.json(requestJson);
});
module.exports = router;
