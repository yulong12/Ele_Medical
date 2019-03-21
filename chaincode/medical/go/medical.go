package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("电子病历")

type SimpleChaincode struct {
}

// 患者姓名：PatientName
// 患者来源：PatientFrom
// 患者身份证号：PatientIdentityNumber

// 患者年龄：PatientAge
// 患者电话：PatientTel
// 患者性别:PatientSex

// 患者民族：PatientNationality
// 患者血型：PatientBloodGroup
// 患者职业：PatientJob

// 家庭地址：PatientAddress

// 患者亲友姓名：PatientRelativeName
// 患者亲友电话：PatientRelativeTel
// 与患者关系：PatientRelation
// 患者亲友身体状况：PatientRelationStatus

// 患者病史：PatientMedicalHistory
// 患者吸烟史:PatientSmokeHistory
// 患者家族遗传史：PatientFamilyGeneticHistory
// 患者主治医生姓名：PatientDoctor
// 患者就诊科室：PatientdePartment
type PatientBasicInfo struct {
	PatientName           string `json:"PatientName "`
	PatientFrom           string `json:"PatientFrom"`
	PatientIdentityNumber string `json:"PatientIdentityNumber "`

	PatientAge string `json:"PatientAge"`
	PatientTel string `json:"PatientTel "`
	PatientSex string `json:"PatientSex "`

	PatientNationality string `json:"PatientNationality "`
	PatientBloodGroup  string `json:"PatientBloodGroup "`
	PatientJob         string `json:"PatientJob "`

	PatientAddress string `json;"PatientAddress"`

	PatientRelativeName   string `json:"PatientRelativeName"`
	PatientRelativeTel    string `json:"PatientRelativeTel "`
	PatientRelation       string `json;"PatientRelation"`
	PatientRelationStatus string `json:"PatientRelationStatus"`

	PatientMedicalHistory       string `json:"PatientMedicalHistory "`
	PatientSmokeHistory         string `json:"PatientSmokeHistory "`
	PatientFamilyGeneticHistory string `json:"PatientFamilyGeneticHistory"`
	PatientDoctor               string `json:"PatientDoctor"`
	PatientdePartment           string `json:"PatientdePartment"`
}

//  targetNo: 目标医院编号
//   redicalNO: 病历编号
//   applier: 申请人
//   applierNo: 申请人编号
//   applierHisNo:申请人所在医院编号
type ApplyRemoteRecord struct {
	TargetNo     string `json:"TargetNo"`
	RedicalNO    string `json:"RedicalNO"`
	Applier      string `json:"Applier"`
	ApplierNo    string `json:"ApplierNo"`
	ApplierHisNo string `json:"ApplierHisNo"`
}

//patientNo:病人ID
//DocterNo:医生职工编号
//HisNo：医院编号
// recordNo: 病历编号,
// recordName: 病历名称,
// recordPath: 病历存储路径,
// recordSize: 病历大小,
// recordHash: 病历的hash值
type UploadData struct {
	PatientNo  string `json:"PatientNo"`
	DocterNo   string `json:"DocterNo"`
	HisNo      string `json:"HisNo"`
	RecordNo   string `json:"RecordNo"`
	RecordName string `json:"RecordName"`
	RecordPath string `json:"RecordPath"`
	RecordSize string `json:"RecordSize"`
	RecordHash string `json:"RecordHash"`
}

const APPLYINDEX = "TargetNo~RedicalNO~Applier~ApplierNo~ApplierHisNo"
const InvalidNumArgs = "参数数量错误"
const MarshalFailed = "json序列化错误"
const SaveStubFailed = "存入区块链失败"
const SaveBlockSuc = "成功存入区块链"
const CreateKey = "创建组合键失败"
const GetDataFBlock = "从区块链中取出数据失败"
const UnmarshlFailed = "json反序列化失败"
const INDEX = "PatientNo~DocterNo~HisNo~RecordNo~RecordName~RecordPath~RecordSize~RecordHash"

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
	} else if fun == "saveHospitalized" {
		return sc.saveHospitalized(stub, args)
	} else if fun == "queryHospitalized" {
		return sc.queryHospitalized(stub, args)
	} else if fun == "applyRemoteData" {
		return sc.applyRemoteData(stub, args)
	} else if fun == "uploadRecordData" {
		return sc.uploadRecordData(stub, args)
	} else if fun == "queryRecordData" {
		return sc.queryRecordData(stub, args)
	}
	return shim.Success(nil)

}

type HospitalData struct {
	Name       string `json:"Name"`
	Age        string `json:"Age"`
	Phone      string `json:"Phone"`
	IdCard     string `json:"IdCard"`
	Sex        string `json:"Sex"`
	Address    string `json:"Address"`
	Doctor     string `json:"Doctor"`
	Nurse      string `json:"Nurse"`
	Illness    string `json:"Illness"`
	Treatment  string `json:"Treatment"`
	Medication string `json:"Medication"`
	Attention  string `json:"Attention"`
	Room       string `json:"Room"`
	InTime     string `json:"InTime"`
	OutTime    string `json:"OutTime"`
	Cost       string `json:"Cost"`
}

//住院病历存储
// name: name,
// age: age,
// phone: phone,
// idCard: idCard,
// sex: sex,
// address: address,
// doctor: doctor,
// nurse: nurse,
// illness: illness,
// treatment: treatment,
// medication: medication,
// attention: attention,
// room: room,
// inTime: inTime,
// outTime: outTime,
// cost: cost
func (sc *SimpleChaincode) saveHospitalized(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 16 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var hisData HospitalData
	hisData.Name = args[0]
	hisData.Age = args[1]
	hisData.Phone = args[2]
	//计算身份证号的哈希值
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(args[3]))
	Result := Sha1Inst.Sum([]byte(""))

	hisData.IdCard = string(Result[:])
	logger.Infof("=====string(Result[:])====%s=====", string(Result[:]))
	hisData.Sex = args[4]
	hisData.Address = args[5]
	hisData.Doctor = args[6]
	hisData.Nurse = args[7]
	hisData.Illness = args[8]
	hisData.Treatment = args[9]
	hisData.Medication = args[10]
	hisData.Attention = args[11]
	hisData.Room = args[12]
	hisData.InTime = args[13]
	hisData.OutTime = args[14]
	hisData.Cost = args[15]

	hisDatab, err := json.Marshal(hisData)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	err = stub.PutState(hisData.IdCard, hisDatab)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//查询患者住院病历
// args:idcard

func (sc *SimpleChaincode) queryHospitalized(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error

	//计算身份证号的哈希值
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(args[0]))
	Result := Sha1Inst.Sum([]byte(""))
	logger.Infof("=====string(Result[:])====%s=====", string(Result[:]))
	hisDatab, err := stub.GetState(string(Result[:]))
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}

	return shim.Success(hisDatab)
}

//args:
// TargetNo
// RedicalNO
// Applier
// ApplierNo
// ApplierHisNo
func (sc *SimpleChaincode) applyRemoteData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	applyK, err := stub.CreateCompositeKey(APPLYINDEX, []string{args[0], args[1], args[2], args[3], args[4]})
	logger.Infof("=====sapplyK===%s=====", applyK)
	if err != nil {
		return shim.Error(getErrReason("CreateCompositeKey", "0"))
	}
	value := []byte{0x00}
	err = stub.PutState(applyK, value)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason("已发出申请", "1"))
}

//patientNo
//DocterNo:医生职工编号
//HisNo:医院编号
// recordNo: 病历编号,
// recordName: 病历名称,
// recordPath: 病历存储路径,
// recordSize: 病历大小,
// recordHash: 病历的hash值
func (sc *SimpleChaincode) uploadRecordData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var uploadData UploadData
	uploadData.PatientNo = args[0]
	uploadData.DocterNo = args[1]
	uploadData.HisNo = args[2]
	uploadData.RecordNo = args[3]
	uploadData.RecordName = args[4]
	uploadData.RecordPath = args[5]
	uploadData.RecordSize = args[6]
	uploadData.RecordHash = args[7]

	t1 := time.Now().Unix()
	strtime := strconv.FormatInt(t1, 10)
	value := []byte{0x00}
	k, err := stub.CreateCompositeKey(INDEX, []string{uploadData.PatientNo, uploadData.DocterNo, uploadData.HisNo, uploadData.RecordNo, uploadData.RecordName, uploadData.RecordPath, uploadData.RecordSize, uploadData.RecordHash, strtime})

	logger.Infof("=====uploadRecordData=k===%s=====", k)
	if err != nil {
		return shim.Error(getErrReason("CreateCompositeKey失败", "0"))
	}
	err = stub.PutState(k, value)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason("上传电子病历成功", "1"))
}

//args :patientNo
func (sc *SimpleChaincode) queryRecordData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}

	Iterator, err := stub.GetStateByPartialCompositeKey(INDEX, []string{args[0]})
	if err != nil {
		return shim.Error(getErrReason("GetStateByPartialCompositeKey failed", "0"))
	}
	defer Iterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for Iterator.HasNext() {
		responseRange, err := Iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		_, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)

		if err != nil {
			return shim.Error(err.Error())
		}
		PatientNo := compositeKeyParts[0]
		DocterNo := compositeKeyParts[1]
		HisNo := compositeKeyParts[2]
		RecordNo := compositeKeyParts[3]
		RecordName := compositeKeyParts[4]
		RecordPath := compositeKeyParts[5]
		RecordSize := compositeKeyParts[6]
		RecordHash := compositeKeyParts[7]
		time := compositeKeyParts[8]

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"PatientNo\":")
		buffer.WriteString("\"")
		buffer.WriteString(PatientNo)
		buffer.WriteString("\"")

		buffer.WriteString("{\"DocterNo\":")
		buffer.WriteString("\"")
		buffer.WriteString(DocterNo)
		buffer.WriteString("\"")

		buffer.WriteString("{\"HisNo\":")
		buffer.WriteString("\"")
		buffer.WriteString(HisNo)
		buffer.WriteString("\"")

		buffer.WriteString("{\"RecordNo\":")
		buffer.WriteString("\"")
		buffer.WriteString(RecordNo)
		buffer.WriteString("\"")

		buffer.WriteString("{\"RecordName\":")
		buffer.WriteString("\"")
		buffer.WriteString(RecordName)
		buffer.WriteString("\"")

		buffer.WriteString("{\"RecordPath\":")
		buffer.WriteString("\"")
		buffer.WriteString(RecordPath)
		buffer.WriteString("\"")

		buffer.WriteString("{\"RecordSize\":")
		buffer.WriteString("\"")
		buffer.WriteString(RecordSize)
		buffer.WriteString("\"")

		buffer.WriteString("{\"RecordHash\":")
		buffer.WriteString("\"")
		buffer.WriteString(RecordHash)
		buffer.WriteString("\"")

		buffer.WriteString("{\"time\":")
		buffer.WriteString("\"")
		buffer.WriteString(time)
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	// var err error
	// var uploadData UploadData

	return shim.Success(buffer.Bytes())
}

// 患者姓名：PatientName
// 患者来源：PatientFrom
// 患者身份证号：PatientIdentityNumber

// 患者年龄：PatientAge
// 患者电话：PatientTel
// 患者性别:PatientSex

// 患者民族：PatientNationality
// 患者血型：PatientBloodGroup
// 患者职业：PatientJob

// 家庭地址：PatientAddress

// 患者亲友姓名：PatientRelativeName
// 患者亲友电话：PatientRelativeTel
// 与患者关系：PatientRelation
// 患者亲友身体状况：PatientRelationStatus

// 患者病史：PatientMedicalHistory
// 患者吸烟史:PatientSmokeHistory
// 患者家族遗传史：PatientFamilyGeneticHistory
// 患者主治医生姓名：PatientDoctor
// 患者就诊科室：PatientdePartment
func (sc *SimpleChaincode) addPatientBasicInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 19 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var patientBasicInfo PatientBasicInfo
	patientBasicInfo.PatientName = args[0]
	patientBasicInfo.PatientFrom = args[1]
	patientBasicInfo.PatientIdentityNumber = args[2]

	patientBasicInfo.PatientAge = args[3]
	patientBasicInfo.PatientTel = args[4]
	patientBasicInfo.PatientSex = args[5]

	patientBasicInfo.PatientNationality = args[6]
	patientBasicInfo.PatientBloodGroup = args[7]
	patientBasicInfo.PatientJob = args[8]

	patientBasicInfo.PatientAddress = args[9]

	patientBasicInfo.PatientRelativeName = args[10]
	patientBasicInfo.PatientRelativeTel = args[11]
	patientBasicInfo.PatientRelation = args[12]
	patientBasicInfo.PatientRelationStatus = args[13]

	patientBasicInfo.PatientMedicalHistory = args[14]
	patientBasicInfo.PatientSmokeHistory = args[15]
	patientBasicInfo.PatientFamilyGeneticHistory = args[16]
	patientBasicInfo.PatientDoctor = args[17]
	patientBasicInfo.PatientdePartment = args[18]

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

	buffer.WriteString("{\"PatientAge  \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientAge)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelation   \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelation)
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

	buffer.WriteString("{\"PatientRelation        \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelation)
	buffer.WriteString("\"")

	buffer.WriteString("{\"PatientRelationStatus        \":")
	buffer.WriteString("\"")
	buffer.WriteString(patientBasicInfo.PatientRelationStatus)
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
