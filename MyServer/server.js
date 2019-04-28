const express = require("express");
const bodyParser = require("body-parser");

const app = express();

var users = require("./routes/api/users");
var uploadRecord = require("./routes/api/uploadRecord");
var applyRecord = require("./routes/api/applyRecord");
var addHospitalized = require("./routes/api/addHospitalized");
var queryRecord = require("./routes/api/queryRecord");
var queryHospitalized = require("./routes/api/queryHospitalized");
var sendCard = require("./routes/api/sendCard");
// 使用body-parser中间件
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use("/api/users", users);
app.use("/api/uploadRecord", uploadRecord);
app.use("/api/applyRecord", applyRecord);
app.use("/api/addHospitalized", addHospitalized);
app.use("/api/queryRecord", queryRecord);
app.use("/api/queryHospitalized", queryHospitalized);
app.use("/api/sendCard", sendCard);

const port = process.env.PORT || 5001;

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
