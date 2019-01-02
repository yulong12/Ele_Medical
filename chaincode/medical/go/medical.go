package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("电子病历")

type SimpleChaincode struct {
}

// 姓名	性别	年龄	婚姻状况	   职业	  工作单位	 普通数据
// name	sex	   age	 marry_status	job	  employer	common_message
// 身份证号	          出生地	   民族	        入院日期	    陈述者	      现住址
// identity_number	birth_place	nationality	admission_date	presenter	Current_address

type Privace_data struct {
	Name            string `json:"Name"`
	Sex             string `json:"Sex"`
	Age             string `json:"Age"`
	Marry_status    string `json:"Marry_status"`
	Job             string `json:"Job"`
	Employer        string `json:"Employer"`
	Identity_number string `json:"Identity_number"`
	Birth_place     string `json:"Birth_place"`
	Nationality     string `json:"Nationality"`
	Admission_date  string `json:"Admission_date"`
	Presenter       string `json:"Presenter"`
	Common_message  string `json:"Common_message"`
}

// 主述	            现病史	                  既往史
// main_statement	current_medical_history	past_history
// 个人史	          生育史	      婚姻史
// personal_history	birth_history	marriage_history

type Common_data struct {
	Identity_number         string `json:"Identity_number"`
	Main_statement          string `json:"Main_statement"`
	Current_medical_history string `json:"Current_medical_history"`
	Past_history            string `json:"Past_history"`
	Personal_history        string `json:"Personal_history"`
	Birth_history           string `json:"Birth_history"`
	Marriage_history        string `json:"Marriage_history"`
}

// 体温	脉搏	心率	血压	发育	营养
// T	P	   R	  BP	growth	nutrition
// 面容	 表情	      体位	      神志	               其他
// face	expression	position	consciousness		others

type Examination_table struct {
	Identity_number string `json:"Identity_number"`
	T               string `json:"T"`
	P               string `json:"P"`
	R               string `json:"R"`
	BP              string `json:"BP"`
	Growth          string `json:"Growth"`
	Nutrition       string `json:"Nutrition"`
	Face            string `json:"Face"`
	Expression      string `json:"Expression"`
	Position        string `json:"position"`
	Consciousness   string `json:"Consciousness"`
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
	if fun == "AddPrivateData" {
		return sc.AddPrivateData(stub, args)
	} else if fun == "AddCommomData" {
		return sc.AddCommomData(stub, args)
	} else if fun == "AddExaminationData" {
		return sc.AddExaminationData(stub, args)
	} else if fun == "QueryPrivateData" {
		return sc.QueryPrivateData(stub, args)
	} else if fun == "QueryExaminationData" {
		return sc.QueryExaminationData(stub, args)
	} else if fun == "QueryCommonData" {
		return sc.QueryCommonData(stub, args)

	}

	return shim.Success(nil)

}

//将隐私数据存入到区块中
// Name Sex Age  Marry_status  Job   Employer  Identity_number Birth_place  Nationality   Admission_date  Presenter  Common_message
func (sc *SimpleChaincode) AddPrivateData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 12 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error

	var private_data Privace_data
	private_data.Name = args[0]
	private_data.Sex = args[1]
	private_data.Age = args[2]
	private_data.Marry_status = args[3]
	private_data.Job = args[4]
	private_data.Employer = args[5]
	private_data.Identity_number = args[6]
	private_data.Birth_place = args[7]
	private_data.Nationality = args[8]
	private_data.Admission_date = args[9]
	private_data.Presenter = args[10]
	private_data.Common_message = args[11]
	private_b, err := json.Marshal(private_data)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	private_key, err := stub.CreateCompositeKey("PrivateData", []string{private_data.Identity_number})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}
	err = stub.PutPrivateData("collectionRecordPrivateDetails", private_key, private_b)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}
	return shim.Success(getRetReason(SaveBlockSuc, "6"))
}

//将普通数据存入区块和mysql数据库中

// Identity_number  Main_statement  Current_medical_history Past_history   Personal_history    Birth_history     Marriage_history
func (sc *SimpleChaincode) AddCommomData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	var err error
	var common_data Common_data
	common_data.Identity_number = args[0]
	common_data.Main_statement = args[1]
	common_data.Current_medical_history = args[2]
	common_data.Past_history = args[3]
	common_data.Personal_history = args[4]
	common_data.Birth_history = args[5]
	common_data.Marriage_history = args[6]
	commom_b, err := json.Marshal(common_data)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}

	h := sha256.New()
	h.Write(commom_b)
	bs := h.Sum(nil)

	logger.Infof("===========commom_b======hashvalue========k=%s", bs)
	common_key, err := stub.CreateCompositeKey("CommomData", []string{common_data.Identity_number})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}
	err = stub.PutPrivateData("collectionRecord", common_key, commom_b)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))

}

//将体检数据存入到区块和数据库中

// Identity_number T  P   R  BP  Growth Nutrition  Face   Expression  Position  Consciousness
func (sc *SimpleChaincode) AddExaminationData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var examination_data Examination_table
	if len(args) != 11 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	examination_data.Identity_number = args[0]
	examination_data.T = args[1]
	examination_data.P = args[2]
	examination_data.R = args[3]
	examination_data.BP = args[4]
	examination_data.Growth = args[5]
	examination_data.Nutrition = args[6]
	examination_data.Face = args[7]
	examination_data.Expression = args[8]
	examination_data.Position = args[9]
	examination_data.Consciousness = args[10]
	examination_b, err := json.Marshal(examination_data)
	if err != nil {
		return shim.Error(getErrReason(MarshalFailed, "0"))
	}
	h := sha256.New()
	h.Write(examination_b)
	bs := h.Sum(nil)

	logger.Infof("===========examination_b======hashvalue========k=%s", bs)

	examination_key, err := stub.CreateCompositeKey("ExaminationData", []string{examination_data.Identity_number})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}
	err = stub.PutPrivateData("collectionRecord", examination_key, examination_b)
	if err != nil {
		return shim.Error(getErrReason(SaveStubFailed, "0"))
	}

	return shim.Success(getRetReason(SaveBlockSuc, "1"))
}

//args:身份证号
func (sc *SimpleChaincode) QueryPrivateData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var privateData Privace_data

	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	logger.Infof("===========args========k=%s", args[0])

	//从区块链中取出隐私数据
	private_k, err := stub.CreateCompositeKey("PrivateData", []string{args[0]})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}

	logger.Infof("===========private_k========k=%s", private_k)
	private_b, err := stub.GetPrivateData("collectionRecordPrivateDetails", private_k)
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}

	logger.Infof("===========private_k 1========k=%s", string(private_b[:]))
	err = json.Unmarshal(private_b, &privateData)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}

	return shim.Success(private_b)
}

//查询普通数据 args:身份证号
func (sc *SimpleChaincode) QueryCommonData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var commonData Common_data
	var err error
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	commom_k, err := stub.CreateCompositeKey("CommomData", []string{args[0]})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}
	logger.Infof("===========commom_k========k=%s", commom_k)
	common_b, err := stub.GetPrivateData("collectionRecord", commom_k)

	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}

	err = json.Unmarshal(common_b, &commonData)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}
	return shim.Success(common_b)
}

//查询体健数据 args:身份证号
func (sc *SimpleChaincode) QueryExaminationData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var examinationData Examination_table
	var err error
	if len(args) != 1 {
		return shim.Error(getErrReason(InvalidNumArgs, "0"))
	}
	examination_k, err := stub.CreateCompositeKey("ExaminationData", []string{args[0]})
	if err != nil {
		return shim.Error(getErrReason(CreateKey, "0"))
	}
	logger.Infof("===========examination_k========k=%s", examination_k)

	examination_b, err := stub.GetPrivateData("collectionRecord", examination_k)
	if err != nil {
		return shim.Error(getErrReason(GetDataFBlock, "0"))
	}
	err = json.Unmarshal(examination_b, &examinationData)
	if err != nil {
		return shim.Error(getErrReason(UnmarshlFailed, "0"))
	}
	return shim.Success(examination_b)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
