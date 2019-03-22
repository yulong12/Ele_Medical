//
const express = require("express");
const router = express.Router();
var multer = require("multer");
var crypto = require("crypto");
var fs = require("fs");
var uploadRecord = require("./../myhfc/myhfcInvoke");
router.get("/uploadRecord", (req, res) => {
  res.json({ msg: "login works" });
});
//获取时间
function getNowFormatDate() {
  var date = new Date();
  var seperator1 = "-";
  var month = date.getMonth() + 1;
  var strDate = date.getDate();
  if (month >= 1 && month <= 9) {
    month = "0" + month;
  }
  if (strDate >= 0 && strDate <= 9) {
    strDate = "0" + strDate;
  }
  var currentdate =
    date.getFullYear() + seperator1 + month + seperator1 + strDate;
  return currentdate.toString();
}
var datatime = "public/images/" + getNowFormatDate();
//将图片放到服务器
var storage = multer.diskStorage({
  // 如果你提供的 destination 是一个函数，你需要负责创建文件夹
  destination: datatime,
  //给上传文件重命名，获取添加后缀名
  filename: function(req, file, cb) {
    cb(null, file.originalname);
  }
});
var upload = multer({
  storage: storage
});

router.post("/uploadRecord", upload.single("picUrl"), function(req, res, next) {
  console.log("----picTitle-------" + req.body.picTitle); //console.log(req.query.picTitle);//get
  console.log("-------picType-----" + req.body.picType);
  console.log("-------name-----" + req.body.name);
  console.log(req.file.recordNo); //req.file文件的具体信息

  //计算该文件的md5值
  var buffer = fs.readFileSync(datatime + "/" + req.file.filename);
  var fsHash = crypto.createHash("md5");
  fsHash.update(buffer);
  var md5 = fsHash.digest("hex");
  console.log("文件的MD5是：%s", md5);

  //向区块链中存储
  var recordNo = req.body.recordNo; //病历编号
  var patientNo = req.body.patientNo; //病人ID
  var doctorNo = req.body.doctorNo;
  var hisNO = req.body.hisNO;
  console.log("-------recordNo-----" + recordNo);
  var recordName = req.file.originalname; //病历名字，e:X照片
  var recordPath = req.file.path; //存储位置
  var recordSize = req.file.size + ""; //病历大小
  //patientNo
  //DocterNo:医生职工编号
  //HisNo:医院编号
  // recordNo: 病历编号,
  // recordName: 病历名称,
  // recordPath: 病历存储路径,
  // recordSize: 病历大小,
  // recordHash: 病历的hash值
  var requestJson = {
    fcn: "uploadRecordData",
    args: [
      patientNo,
      doctorNo,
      hisNO,
      recordNo,
      recordName,
      recordPath,
      recordSize,
      md5
    ]
  };
  uploadRecord(requestJson, function(str) {
    console.log("=====str==========" + str.status);
    console.log("=====payload==========" + str.payload);

    res.send({
      status: "OK",
      detail: str
    });
  });
  // var str = JSON.stringify(requestJson);
  // console.log("-------str-----" + str);
  // res.send({ ret_code: datatime, md5: md5 });
});

module.exports = router;
