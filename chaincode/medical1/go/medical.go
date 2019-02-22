package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("电子病历")

type SimpleChaincode struct {
}

type PatientBasicInfo struct {
	PatientName                 string `json:"PatientName "`
	PatientFrom                 string `json:"PatientFrom"`
	PatientIdentityNumber       string `json:"PatientIdentityNumber "`
	PatientSex                  string `json:"PatientSex "`
	PatientBirthDate            string `json:"PatientBirthDate "`
	PatientNationality          string `json:"PatientNationality "`
	PatientBloodGroup           string `json:"PatientBloodGroup "`
	PatientJob                  string `json:"PatientJob "`
	PatientTel                  string `json:"PatientTel "`
	PatientRelativeName         string `json:"PatientRelativeName"`
	PatientRelativeTel          string `json:"PatientRelativeTel "`
	PatientMedicalHistory       string `json:"PatientMedicalHistory "`
	PatientSmokeHistory         string `json:"PatientSmokeHistory "`
	PatientFamilyGeneticHistory string `json:"PatientFamilyGeneticHistory"`
	PatientDoctor               string `json:"PatientDoctor"`
	PatientdePartment           string `json:"PatientdePartment"`
}

const InvalidNumArgs = "参数数量错误"
const MarshalFailed = "json序列化错误"
const SaveStubFailed = "存入区块链失败"
const SaveBlockSuc = "成功存入区块链"
const CreateKey = "创建组合键失败"
const GetDataFBlock = "从区块链中取出数据失败"
const UnmarshlFailed = "json反序列化失败"

type ErrReason struct {
	Statue string `json:"Statue"`
	Reason string `json:"Reason"`
}

func getErrReason(des string, jud string) string {
	var reason ErrReason
	reason.Reason = des
	reason.Statue = jud
	b, _ := json.Marshal(reason)
	return string(b[:])
}
func getRetReason(des string, jud string) []byte {
	var reason ErrReason
	reason.Reason = des
	reason.Statue = jud
	b, _ := json.Marshal(reason)
	return b
}
func (sc *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Success(nil)
}

func (sc *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fun, args := stub.GetFunctionAndParameters()
	if fun == "addPatientBasicInfo" {
		return sc.addPatientBasicInfo(stub, args)

	} else if fun == "queryPatientBasicInfo" {
		return sc.queryPatientBasicInfo(stub, args)
	}
	return shim.Success(nil)

}

// 患者姓名：PatientName
// 患者来源：PatientFrom
// 患者身份证号：PatientIdentityNumber
// 患者性别:PatientSex
// 患者出生日期：PatientBirthDate
// 患者民族：PatientNationality
// 患者血型：PatientBloodGroup
// 患者职业：PatientJob
// 患者电话：PatientTel
// 患者亲友姓名：PatientRelativeName
// 患者亲友电话：PatientRelativeTel
// 患者病史：PatientMedicalHistory
// 患者吸烟史:PatientSmokeHistory
// 患者家族遗传史：PatientFamilyGeneticHistory
// 患者主治医生姓名：PatientDoctor
// 患者就诊科室：PatientdePartment
func (sc *SimpleChaincode) addPatientBasicInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 16 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientBasicInfo PatientBasicInfo
	patientBasicInfo.PatientName = args[0]
	patientBasicInfo.PatientFrom = args[1]
	patientBasicInfo.PatientIdentityNumber = args[2]
	patientBasicInfo.PatientSex = args[3]
	patientBasicInfo.PatientBirthDate = args[4]
	patientBasicInfo.PatientNationality = args[5]
	patientBasicInfo.PatientBloodGroup = args[6]
	patientBasicInfo.PatientJob = args[7]
	patientBasicInfo.PatientTel = args[8]
	patientBasicInfo.PatientRelativeName = args[9]
	patientBasicInfo.PatientRelativeTel = args[10]
	patientBasicInfo.PatientMedicalHistory = args[11]
	patientBasicInfo.PatientSmokeHistory = args[12]
	patientBasicInfo.PatientFamilyGeneticHistory = args[13]
	patientBasicInfo.PatientDoctor = args[14]
	patientBasicInfo.PatientdePartment = args[15]
	patientBasicInfob, err := json.Marshal(patientBasicInfo)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(patientBasicInfo.PatientIdentityNumber, patientBasicInfob)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

func (sc *SimpleChaincode) queryPatientBasicInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientBasicInfo PatientBasicInfo

	PatientBasicInfob, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(PatientBasicInfob, &patientBasicInfo)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}
	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"PatientName\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientFrom\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientFrom)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientIdentityNumber \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientIdentityNumber)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientSex  \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientSex)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientBirthDate   \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientBirthDate)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientNationality    \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientNationality)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientBloodGroup     \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientBloodGroup)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientJob      \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientJob)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientTel       \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientTel)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelativeName        \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelativeName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelativeTel        \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelativeTel)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientMedicalHistory         \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientMedicalHistory)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientSmokeHistory          \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientSmokeHistory)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientFamilyGeneticHistory           \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientFamilyGeneticHistory)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientDoctor\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientDoctor)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientdePartment\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientdePartment)
	buffer.WriteString("}")

	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
