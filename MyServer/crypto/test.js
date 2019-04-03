var util = require("./cryptoUtils");
// util.Generator("./pem/public.pem", "./pem/private.pem");
util.Encrypt("./pem/public.pem", "asfafsdads", function(encrpts) {
  console.log("==============" + encrpts);
  util.Decrypt("./pem/private.pem", encrpts, function(str) {
    console.log("==========" + str);
  });
});
util.SignData("./pem/private.pem", "asfafsdads", function(str) {
  console.log("-----str-----" + str);
  util.VerySign("./pem/public.pem", "asfafsdads", function(str) {
    console.log("-----str-------" + str);
  });
});
// console.log("========enstr", enstr);
// var destr = util.Decrypt(enstr, "./pem/xiaogangprivate.pem");
// console.log("========destr", destr);
// // util.Generator("./pem/public.pem", "./pem/private.pem");
// util.DeTest("./pem/public.pem", "./pem/private.pem");
