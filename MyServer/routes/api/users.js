// @login & register
const express = require("express");
const router = express.Router();
var options = require("./../myhfc/org1Config");
// @route  GET api/users/test
// @desc   返回的请求的json数据
// @access public
router.get("/test", (req, res) => {
  res.json({ msg: "login works" });
});

router.post("/test", (req, res) => {
  var name = req.body.name;
  var pwd = req.body.password;
  res.json({
    name: name,
    pwd: pwd,
    status: "sucess",
    dahua: "你好"
  });
});
router.post("/addPatient", (req, res) => {
  var addPatientBasicInfo = require("./../myhfc/myhfcInvoke");
  var name = req.body.name;
  var region = req.body.region;
  var idCard = req.body.idCard;
  var age = req.body.age;
  var phone = req.body.phone;
  var sex = req.body.sex;
  var nationality = req.body.nationality;
  var bloodGroup = req.body.bloodGroup;
  var job = req.body.job;
  var address = req.body.address;
  var Name = req.body.Name;
  var Phone = req.body.Phone;
  var relation = req.body.relation;
  var symptoms = req.body.symptoms;
  var illHistory = req.body.illHistory;
  var geneticHistory = req.body.geneticHistory;
  var smokeHistory = req.body.smokeHistory;
  var doctor = req.body.doctor;
  var department = req.body.department;
  console.log("-----name-------" + name);
  var request = {
    fcn: "addPatientBasicInfo",
    args: [
      name,
      region,
      idCard,
      age,
      phone,
      sex,
      nationality,
      bloodGroup,
      job,
      address,
      Name,
      Phone,
      relation,
      symptoms,
      illHistory,
      geneticHistory,
      smokeHistory,
      doctor,
      department
    ]
  };
  var stt = JSON.stringify(request);
  console.log("-----stt-------" + stt);
  addPatientBasicInfo(request, function(str) {
    console.log("=====str==========" + str.status);
    console.log("=====payload==========" + str.payload);

    res.send({
      status: "OK",
      detail: str
    });
  });

  // res.json({
  //   status: "sucess",
  //   dahua: "你好"
  // });
});

module.exports = router;
