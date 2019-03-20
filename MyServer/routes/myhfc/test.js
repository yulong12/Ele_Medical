var invoke = require("./myhfcInvoke.1");
var query = require("./myhfcQuery.1");
// var request = {
//   fcn: "invoke",
//   args: ["a", "b", "1"]
// };
// invoke(request, function(jsonStr) {
//   console.log("------status---" + jsonStr.status);
//   console.log("-----payload----" + jsonStr.payload);
// });

var requestqueyr = {
  fcn: "query",
  args: ["a"]
};
query(requestqueyr, function(jsonStr) {
  console.log("------status---" + jsonStr.status);
  console.log("-----payload----" + jsonStr.str);
});
