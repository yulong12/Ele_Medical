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

//患者基本信息数据
type PatientBasicInfo struct {
	PatientNum                  string `json:"PatientNum"`                  //患者编号
	PatientName                 string `json:"PatientName"`                 //患者姓名
	PatientIdentityNumber       string `json:"PatientIdentityNumber"`       //患者身份证号
	PatientSex                  string `json:"PatientSex"`                  //患者性别
	PatientBirthDate            string `json:"PatientBirthDate"`            //患者出生日期
	PatientNationality          string `json:"PatientNationality"`          //患者民族
	PatientBloodGroup           string `json:"PatientBloodGroup"`           //患者血型：
	PatientJob                  string `json:"PatientJob"`                  //患者职业
	PatientEmployer             string `json:"PatientEmployer"`             //患者工作单位
	PatientTel                  string `json:"PatientTel"`                  //患者电话
	PatientRelativeName         string `json:"PatientRelativeName"`         //患者亲友姓名
	PatientRelativeTel          string `json:"PatientRelativeTel"`          //患者亲友电话
	PatientMarriageStatus       string `json:"PatientMarriageStatus"`       //患者婚姻状况
	PatientFertilityStatus      string `json:"PatientFertilityStatus"`      //患者生育状况
	PatientFertilityAge         string `json:"PatientFertilityAge"`         //患者生育年龄
	PatientSonNum               string `json:"PatientSonNum"`               //患者儿子数量
	PatientGirlNum              string `json:"PatientGirlNum"`              //患者女儿数量
	PatientMedicalHistory       string `json:"PatientMedicalHistory"`       //患者病史
	PatientSmokeHistory         string `json:"PatientSmokeHistory"`         //患者吸烟史
	PatientFamilyGeneticHistory string `json:"PatientFamilyGeneticHistory"` //患者家族遗传史
}

//住院信息表
type ResidentInfo struct {
	ResidentNum             string `json:"ResidentNum"`             // 住院流水号
	ResidentName            string `json:"ResidentName"`            // 住院者姓名
	ResidentPatientNum      string `json:"ResidentPatientNum"`      // 患者编号
	ResidentDepartment      string `json:"ResidentDepartment"`      // 就诊科室
	ResidentBedNum          string `json:"ResidentBedNum"`          // 床位号
	ResidentSupervisor      string `json:"ResidentSupervisor"`      // 主管医生：
	ResidentSupervisorNurse string `json:"ResidentSupervisorNurse"` //主管护士
	ResidentInDate          string `json:"ResidentInDate"`          //住院日期
	ResidentOutDate         string `json:"ResidentOutDate"`         //出院日期
	ResidentAttention       string `json:"ResidentAttention"`       //备注
}

//电子病历信息
type RecordInfo struct {
	RecordNum          string `json:"RecordNum"`          //文档编号
	RecordResidentNum  string `json:"RecordResidentNum"`  //住院流水号：
	RecordName         string `json:"RecordName"`         //文档名称
	RecordCreateDoctor string `json:"RecordCreateDoctor"` //创建医生
	RecordCreateTime   string `json:"RecordCreateTime"`   //创建时间
	RecordPath         string `json:"RecordPath"`         //文档路径
	RecordHashValue    string `json:"RecordHashValue"`    //文档哈希值
	RecordAttention    string `json:"RecordAttention"`    //备注
}

//临床路径
type ClinicalPathway struct {
	ClinicalPathwayNum       string `json:"ClinicalPathwayNum"`       //临床路径编号
	ClinicalPathwayName      string `json:"ClinicalPathwayName"`      //临床路径名称：
	ClinicalPathwayAttention string `json:"ClinicalPathwayAttention"` //备注
}

//临床路径项
type ClinicalPathwayItem struct {
	ClinicalPathwayItemNum       string `json:"ClinicalPathwayItemNum"`       //临床路径项编号：
	ClinicalPathwayItemName      string `json:"ClinicalPathwayItemName"`      //临床路径项名称：
	ClinicalPathwayItemAttention string `json:"ClinicalPathwayItemAttention"` //备注
}

// 患者临床路径
type PatientClinicalPathway struct {
	PatientClinicalPathwayNum       string `json:"PatientClinicalPathwayNum"`       //患者临床路径编号
	ClinicalPathwayNum              string `json:"ClinicalPathwayNum"`              //临床路径编号
	PatientClinicalBeginTime        string `json :"PatientClinicalBeginTime"`       //项目开始时间
	ResidentNum                     string `json:"ResidentNum"`                     //住院流水号
	PatientClinicalPathwayCreator   string `json:"PatientClinicalPathwayCreator"`   //创建医生
	PatientClinicalPathwayAttention string `json:"PatientClinicalPathwayAttention"` //备注
}

//临床路径与临床路径项关系：Clinical_Relation
type ClinicalRelation struct {
	ClinicalPathwayNum        string `json:"ClinicalPathwayNum"`        //临床路径编号
	ClinicalPathwayItemNum    string `json:"ClinicalPathwayItemNum"`    //临床路径项编号
	ClinicalRelationAttention string `json:"ClinicalRelationAttention"` //备注
}

//临床路径执行情况
type ClinicalPathwayExecuStatus struct {
	ClinicalPathwayExecuNum       string `json:"ClinicalPathwayExecuNum"`       //临床路径执行编号
	ClinicalPathwayItemNum        string `json:"ClinicalPathwayItemNum"`        //临床路径项编号
	ClinicalPathwayExecuTime      string `json:"ClinicalPathwayExecuTime"`      //执行时间
	ClinicalPathwayExecuStatus    string `json:"ClinicalPathwayExecuStatus"`    //完成情况
	ClinicalPathwayExecuAttention string `json:"ClinicalPathwayExecuAttention"` //备注

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
	} else if fun == "addResidentInfo" {
		return sc.addResidentInfo(stub, args)
	} else if fun == "queryResidentInfo" {
		return sc.queryResidentInfo(stub, args)
	} else if fun == "addRecordInfo" {
		return sc.addRecordInfo(stub, args)
	} else if fun == "queryRecordInfo" {
		return sc.queryRecordInfo(stub, args)
	} else if fun == "addClinicalPathway" {
		return sc.addClinicalPathway(stub, args)
	} else if fun == "queryClinicalPathway" {
		return sc.queryClinicalPathway(stub, args)
	} else if fun == "addClinicalPathwayItem" {
		return sc.addClinicalPathwayItem(stub, args)
	} else if fun == "queryClinicalPathwayItem" {
		return sc.queryClinicalPathwayItem(stub, args)
	} else if fun == "addPatientClinicalPathway" {
		return sc.addPatientClinicalPathway(stub, args)
	} else if fun == "queryPatientClinicalPathway" {
		return sc.queryPatientClinicalPathway(stub, args)
	} else if fun == "addClinicalRelation" {
		return sc.addClinicalRelation(stub, args)
	} else if fun == "queryClinicalRelation" {
		return sc.queryClinicalRelation(stub, args)
	} else if fun == "addClinicalPathwayExecuStatus" {
		return sc.addClinicalPathwayExecuStatus(stub, args)
	} else if fun == "queryClinicalPathwayExecuStatus" {
		return sc.queryClinicalPathwayExecuStatus(stub, args)
	}

	return shim.Success(nil)

}

//添加病人基本信息

// 患者编号：PatientNum
// 患者姓名：PatientName
// 患者身份证号：PatientIdentityNumber
// 患者性别:PatientSex
// 患者出生日期：PatientBirthDate
// 患者民族：PatientNationality
// 患者血型：PatientBloodGroup
// 患者职业：PatientJob
// 患者工作单位：PatientEmployer
// 患者电话：PatientTel
// 患者亲友姓名：PatientRelativeName
// 患者亲友电话：PatientRelativeTel
// 患者婚姻状况：PatientMarriageStatus
// 患者生育状况：PatientFertilityStatus
// 患者生育年龄：PatientFertilityAge
// 患者儿子数量：PatientSonNum
// 患者女儿数量：PatientGirlNum
// 患者病史：PatientMedicalHistory
// 患者吸烟史:PatientSmokeHistory
// 患者家族遗传史：PatientFamilyGeneticHistory
func (sc *SimpleChaincode) addPatientBasicInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 20 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientBasicInfo PatientBasicInfo
	patientBasicInfo.PatientNum = args[0]
	patientBasicInfo.PatientName = args[1]
	patientBasicInfo.PatientIdentityNumber = args[2]
	patientBasicInfo.PatientSex = args[3]
	patientBasicInfo.PatientBirthDate = args[4]
	patientBasicInfo.PatientNationality = args[5]
	patientBasicInfo.PatientBloodGroup = args[6]
	patientBasicInfo.PatientJob = args[7]
	patientBasicInfo.PatientEmployer = args[8]
	patientBasicInfo.PatientTel = args[9]
	patientBasicInfo.PatientRelativeName = args[10]
	patientBasicInfo.PatientRelativeTel = args[11]
	patientBasicInfo.PatientMarriageStatus = args[12]
	patientBasicInfo.PatientFertilityStatus = args[13]
	patientBasicInfo.PatientFertilityAge = args[14]
	patientBasicInfo.PatientSonNum = args[15]
	patientBasicInfo.PatientGirlNum = args[16]
	patientBasicInfo.PatientMedicalHistory = args[17]
	patientBasicInfo.PatientSmokeHistory = args[18]
	patientBasicInfo.PatientFamilyGeneticHistory = args[19]
	PatientBasicInfob, err := json.Marshal(patientBasicInfo)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(patientBasicInfo.PatientNum, PatientBasicInfob)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询病人基本信息
//args: PatientNum
func (sc *SimpleChaincode) queryPatientBasicInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientBasicInfo PatientBasicInfo
	patientBasicInfob, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(patientBasicInfob, &patientBasicInfo)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}
	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"PatientNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientName\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientIdentityNumber\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientIdentityNumber)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientSex\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientSex)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientBirthDate\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientBirthDate)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientNationality\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientNationality)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientBloodGroup\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientBloodGroup)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientJob\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientJob)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientEmployer\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientEmployer)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientTel\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientTel)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelativeName\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelativeName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelativeTel\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelativeTel)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientMarriageStatus\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientMarriageStatus)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientFertilityStatus\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientFertilityStatus)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientFertilityAge\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientFertilityAge)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientSonNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientSonNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientGirlNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientGirlNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientMedicalHistory\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientMedicalHistory)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientSmokeHistory\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientSmokeHistory)
	buffer.WriteString("\"")

	buffer.WriteString(",\"PatientFamilyGeneticHistory\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientFamilyGeneticHistory)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加住院信息Resident_Info

// 住院流水号：ResidentNum
// 住院者姓名；ResidentName
// 患者编号：ResidentPatientNum
// 就诊科室：ResidentDepartment
// 床位号：ResidentBedNum
// 主管医生：ResidentSupervisor
// 主管护士：ResidentSupervisorNurse
// 住院日期：ResidentInDate
// 出院日期:ResidentOutDate
// 备注：ResidentAttention

func (sc *SimpleChaincode) addResidentInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var residentInfo ResidentInfo
	residentInfo.ResidentNum = args[0]
	residentInfo.ResidentName = args[1]
	residentInfo.ResidentPatientNum = args[2]
	residentInfo.ResidentDepartment = args[3]
	residentInfo.ResidentBedNum = args[4]
	residentInfo.ResidentSupervisor = args[5]
	residentInfo.ResidentSupervisorNurse = args[6]
	residentInfo.ResidentInDate = args[7]
	residentInfo.ResidentOutDate = args[8]
	residentInfo.ResidentAttention = args[9]
	residentInfob, err := json.Marshal(residentInfo)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], residentInfob)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询住院信息
//args:住院流水号
func (sc *SimpleChaincode) queryResidentInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(InvalidNumArgs)
	}
	var err error
	var residentInfo ResidentInfo
	residentInfob, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}

	err = json.Unmarshal(residentInfob, &residentInfo)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}
	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"ResidentNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentName\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentPatientNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentPatientNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentDepartment\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentDepartment)
	buffer.WriteString("\"")

	buffer.WriteString("{\"：ResidentBedNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentBedNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentSupervisor\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentSupervisor)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentSupervisorNurse\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentSupervisorNurse)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentInDate\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentInDate)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentOutDate\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentOutDate)
	buffer.WriteString("\"")

	buffer.WriteString(",\"ResidentAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(residentInfo.ResidentAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加电子病历信息

// 文档编号：RecordNum
// 住院流水号：RecordResidentNum
// 文档名称:RecordName
// 创建医生:RecordCreateDoctor
// 创建时间：RecordCreateTime
// 文档路径:RecordPath
// 文档哈希值:RecordHashValue
// 备注：RecordAttention
func (sc *SimpleChaincode) addRecordInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error(InvalidNumArgs)
	}
	var err error
	var recordInfo RecordInfo
	recordInfo.RecordNum = args[0]
	recordInfo.RecordResidentNum = args[1]
	recordInfo.RecordName = args[2]
	recordInfo.RecordCreateDoctor = args[3]
	recordInfo.RecordCreateTime = args[4]
	recordInfo.RecordPath = args[5]
	recordInfo.RecordHashValue = args[6]
	recordInfo.RecordAttention = args[7]

	recordInfob, err := json.Marshal(recordInfo)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], recordInfob)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询电子病历信息
//args:文档编号
func (sc *SimpleChaincode) queryRecordInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var recordInfo RecordInfo
	recordInfob, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(recordInfob, &recordInfo)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"RecordNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordResidentNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordResidentNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordName\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordCreateDoctor\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordCreateDoctor)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordCreateTime\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordCreateTime)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordPath\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordPath)
	buffer.WriteString("\"")

	buffer.WriteString("{\"RecordHashValue\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordHashValue)
	buffer.WriteString("\"")

	buffer.WriteString(",\"RecordAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(recordInfo.RecordAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加临床路

// 临床路径编号：ClinicalPathwayNum
// 临床路径名称：ClinicalPathwayName
// 备注：ClinicalPathwayAttention

func (sc *SimpleChaincode) addClinicalPathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalPathway ClinicalPathway
	clinicalPathway.ClinicalPathwayNum = args[0]
	clinicalPathway.ClinicalPathwayName = args[1]
	clinicalPathway.ClinicalPathwayAttention = args[2]
	clinicalPathwayb, err := json.Marshal(clinicalPathway)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], clinicalPathwayb)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询临床路径
//args :临床路径编号：ClinicalPathwayNum
func (sc *SimpleChaincode) queryClinicalPathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(InvalidNumArgs)
	}
	var err error
	var clinicalPathway ClinicalPathway
	clinicalPathwayb, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(clinicalPathwayb, &clinicalPathway)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"ClinicalPathwayNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathway.ClinicalPathwayNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayName\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathway.ClinicalPathwayName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathway.ClinicalPathwayAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加临床路径项
// 临床路径项编号：ClinicalPathwayItemNum
// 临床路径项名称：ClinicalPathwayItemName
// 备注：ClinicalPathwayItemAttention

func (sc *SimpleChaincode) addClinicalPathwayItem(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalPathwayItem ClinicalPathwayItem
	clinicalPathwayItem.ClinicalPathwayItemNum = args[0]
	clinicalPathwayItem.ClinicalPathwayItemName = args[1]
	clinicalPathwayItem.ClinicalPathwayItemAttention = args[2]
	clinicalPathwayItemb, err := json.Marshal(clinicalPathwayItem)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], clinicalPathwayItemb)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询临床路径项
//args:ClinicalPathwayItemNum
func (sc *SimpleChaincode) queryClinicalPathwayItem(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))

	}
	var err error
	var clinicalPathwayItem ClinicalPathwayItem
	clinicalPathwayItemb, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(clinicalPathwayItemb, clinicalPathwayItem)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"ClinicalPathwayItemNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayItem.ClinicalPathwayItemNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayItemName\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayItem.ClinicalPathwayItemName)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayItemAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayItem.ClinicalPathwayItemAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加患者临床路径
// 患者临床路径编号：PatientClinicalPathwayNum
// 临床路径编号：ClinicalPathwayNum
// 项目开始时间：PatientClinicalBeginTime
// 住院流水号：ResidentNum
// 创建医生：PatientClinicalPathwayCreator
// 备注：PatientClinicalPathwayAttention
func (sc *SimpleChaincode) addPatientClinicalPathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientClinicalPathway PatientClinicalPathway
	patientClinicalPathway.PatientClinicalPathwayNum = args[0]
	patientClinicalPathway.ClinicalPathwayNum = args[1]
	patientClinicalPathway.PatientClinicalBeginTime = args[2]
	patientClinicalPathway.ResidentNum = args[3]
	patientClinicalPathway.PatientClinicalPathwayCreator = args[4]
	patientClinicalPathway.PatientClinicalPathwayAttention = args[5]
	patientClinicalPathwayb, err := json.Marshal(patientClinicalPathway)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], patientClinicalPathwayb)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询患者临床路径
// args:患者临床路径编号：PatientClinicalPathwayNum
func (sc *SimpleChaincode) queryPatientClinicalPathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientClinicalPathway PatientClinicalPathway
	patientClinicalPathwayb, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(patientClinicalPathwayb, &patientClinicalPathway)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"PatientClinicalPathwayNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.PatientClinicalPathwayNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.ClinicalPathwayNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientClinicalBeginTime\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.PatientClinicalBeginTime)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ResidentNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.ResidentNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientClinicalPathwayCreator\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.PatientClinicalPathwayCreator)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientClinicalPathwayAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(patientClinicalPathway.PatientClinicalPathwayAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加临床与临床路径项的关系

// 临床路径编号：ClinicalPathwayNum
// 临床路径项编号：ClinicalPathwayItemNum
// 备注：ClinicalRelationAttention

func (sc *SimpleChaincode) addClinicalRelation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalRelation ClinicalRelation
	clinicalRelation.ClinicalPathwayNum = args[0]
	clinicalRelation.ClinicalPathwayItemNum = args[1]
	clinicalRelation.ClinicalRelationAttention = args[2]
	clinicalRelationb, err := json.Marshal(clinicalRelation)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], clinicalRelationb)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询临床与临床路径项的关系
// 临床路径编号：ClinicalPathwayNum
func (sc *SimpleChaincode) queryClinicalRelation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalRelation ClinicalRelation
	clinicalRelationb, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(clinicalRelationb, &clinicalRelation)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"ClinicalPathwayNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalRelation.ClinicalPathwayNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayItemNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalRelation.ClinicalPathwayItemNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalRelationAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalRelation.ClinicalRelationAttention)
	buffer.WriteString("}")
	// bArrayMemberAlreadyWritten = true
	buffer.WriteString("]")
	logger.Infof("===========buffer======buffer========k=%s", buffer.String())
	return shim.Success(buffer.Bytes())
}

//添加临床路径执行情况
// 临床路径执行编号：ClinicalPathwayExecuNum
// 临床路径项编号：ClinicalPathwayItemNum
// 执行时间：ClinicalPathwayExecuTime
// 完成情况：ClinicalPathwayExecuStatus
// 备注：ClinicalPathwayExecuAttention

func (sc *SimpleChaincode) addClinicalPathwayExecuStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalPathwayExecuStatus ClinicalPathwayExecuStatus
	clinicalPathwayExecuStatus.ClinicalPathwayExecuNum = args[0]
	clinicalPathwayExecuStatus.ClinicalPathwayItemNum = args[1]
	clinicalPathwayExecuStatus.ClinicalPathwayExecuTime = args[2]
	clinicalPathwayExecuStatus.ClinicalPathwayExecuStatus = args[3]
	clinicalPathwayExecuStatus.ClinicalPathwayExecuAttention = args[4]
	clinicalPathwayExecuStatusb, err := json.Marshal(clinicalPathwayExecuStatus)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(args[0], clinicalPathwayExecuStatusb)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询临床路径执行情况
//args:临床路径执行编号：ClinicalPathwayExecuNum
func (sc *SimpleChaincode) queryClinicalPathwayExecuStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var clinicalPathwayExecuStatus ClinicalPathwayExecuStatus
	clinicalPathwayExecuStatusb, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(clinicalPathwayExecuStatusb, clinicalPathwayExecuStatus)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	// bArrayMemberAlreadyWritten := false
	// if bArrayMemberAlreadyWritten == true {
	// 	buffer.WriteString(",")
	// }
	buffer.WriteString("{\"ClinicalPathwayExecuNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayExecuStatus.ClinicalPathwayExecuNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayItemNum\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayExecuStatus.ClinicalPathwayItemNum)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientClinicalBeginTime\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayExecuStatus.ClinicalPathwayExecuTime)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayExecuStatus\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayExecuStatus.ClinicalPathwayExecuStatus)
	buffer.WriteString("\"")

	buffer.WriteString("{\"ClinicalPathwayExecuAttention\":")
	buffer.WriteString("\"")
	buffer.WriteString(clinicalPathwayExecuStatus.ClinicalPathwayExecuAttention)
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
