var invoke = require("./myhfcInvoke");
var query = require("./myhfcQuery");
var request = {
  fcn: "addPatientBasicInfo",
  args: [
    "name",
    "region",
    "idCard",
    "age",
    "phone",
    "sex",
    "nationality",
    "bloodGroup",
    "job",
    "address",
    "Name",
    "Phone",
    "relation",
    "symptoms",
    "illHistory",
    "geneticHistory",
    "smokeHistory",
    "doctor",
    "department"
  ]
};
invoke(request, function(jsonStr) {
  console.log("------status---" + jsonStr.status);
  console.log("-----payload----" + jsonStr.payload);
});

// var requestqueyr = {
//   fcn: "queryPatientBasicInfo",
//   args: ["idCard"]
// };
// query(requestqueyr, function(jsonStr) {
//   console.log("------status---" + jsonStr.status);
//   console.log("-----payload----" + jsonStr.str);
// });
