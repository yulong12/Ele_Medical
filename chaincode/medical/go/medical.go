package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("电子病历")

type SimpleChaincode struct {
}

//患者基本信息数据
type Patient_Basic_Info struct {
	Patient_Num                    string `json:"Patient_Num"`                    //患者编号
	Patient_Name                   string `json:"Patient_Name"`                   //患者姓名
	Patient_Identity_Number        string `json:"Patient_Identity_Number"`        //患者身份证号
	Patient_Sex                    string `json:"Patient_Sex"`                    //患者性别
	Patient_Birth_Date             string `json:"Patient_Birth_Date"`             //患者出生日期
	Patient_Nationality            string `json:"Patient_Nationality"`            //患者民族
	Patient_Blood_Group            string `json:"Patient_Blood_Group"`            //患者血型：
	Patient_Job                    string `json:"Patient_Job"`                    //患者职业
	Patient_Employer               string `json:"Patient_Employer"`               //患者工作单位
	Patient_Tel                    string `json:"Patient_Tel"`                    //患者电话
	Patient_Relative_Name          string `json:"Patient_Relative_Name"`          //患者亲友姓名
	Patient_Relative_Tel           string `json:"Patient_Relative_Tel"`           //患者亲友电话
	Patient_Marriage_Status        string `json:"Patient_Marriage_Status"`        //患者婚姻状况
	Patient_Fertility_Status       string `json:"Patient_Fertility_Status"`       //患者生育状况
	Patient_Fertility_Age          string `json:"Patient_Fertility_Age"`          //患者生育年龄
	Patient_Son_Num                string `json:"Patient_Son_Num"`                //患者儿子数量
	Patient_Girl_Num               string `json:"Patient_Girl_Num"`               //患者女儿数量
	Patient_Medical_History        string `json:"Patient_Medical_History"`        //患者病史
	Patient_Smoke_History          string `json:"Patient_Smoke_History"`          //患者吸烟史
	Patient_Family_Genetic_History string `json:"Patient_Family_Genetic_History"` //患者家族遗传史
}

//住院信息表
type Resident_Info struct {
	Resident_Num              string `json:"Resident_Num"`              // 住院流水号
	Resident_Name             string `json:"Resident_Name"`             // 住院者姓名
	Resident_Patient_Num      string `json:"Resident_Patient_Num"`      // 患者编号
	Resident_Department       string `json:"Resident_Department"`       // 就诊科室
	Resident_Bed_Num          string `json:"Resident_Bed_Num"`          // 床位号
	Resident_Supervisor       string `json:"Resident_Supervisor"`       // 主管医生：
	Resident_Supervisor_Nurse string `json:"Resident_Supervisor_Nurse"` //主管护士
	Resident_In_Date          string `json:"Resident_In_Date"`          //住院日期
	Resident_Out_Date         string `json:"Resident_Out_Date"`         //出院日期
	Resident_Attention        string `json:"Resident_Attention"`        //备注
}

//电子病历信息
type Record_Info struct {
	Record_Num           string `json:"Record_Num"`           //文档编号
	Record_Resident_Num  string `json:"Record_Resident_Num"`  //住院流水号：
	Record_Name          string `json:"Record_Name"`          //文档名称
	Record_Create_Doctor string `json:"Record_Create_Doctor"` //创建医生
	Record_Create_Time   string `json:"Record_Create_Time"`   //创建时间
	Record_Path          string `json:"Record_Path"`          //文档路径
	Record_Hash_Value    string `json:"Record_Hash_Value"`    //文档哈希值
	Record_Attention     string `json:"Record_Attention"`     //备注
}

//临床路径
type Clinical_Pathway struct {
	Clinical_Pathway_Num       string `json:"Clinical_Pathway_Num"`       //临床路径编号
	Clinical_Pathway_Name      string `json:"Clinical_Pathway_Name"`      //临床路径名称：
	Clinical_Pathway_Attention string `json:"Clinical_Pathway_Attention"` //备注
}

//临床路径项
type Clinical_Pathway_Item struct {
	Clinical_Pathway_Item_Num       string `json:"Clinical_Pathway_Item_Num"`       //临床路径项编号：
	Clinical_Pathway_Item_Name      string `json:"Clinical_Pathway_Item_Name"`      //临床路径项名称：
	Clinical_Pathway_Item_Attention string `json:"Clinical_Pathway_Item_Attention"` //备注
}

// 患者临床路径
type Patient_Clinical_Pathway struct {
	Patient_Clinical_Pathway_Num       string `json:"Patient_Clinical_Pathway_Num"`       //患者临床路径编号
	Clinical_Pathway_Num               string `json:"Clinical_Pathway_Num"`               //临床路径编号
	Patient_Clinical_Begin_Time        string `json :"Patient_Clinical_Begin_Time"`       //项目开始时间
	Resident_Num                       string `json:"Resident_Num"`                       //住院流水号
	Patient_Clinical_Pathway_Creator   string `json:"Patient_Clinical_Pathway_Creator"`   //创建医生
	Patient_Clinical_Pathway_Attention string `json:"Patient_Clinical_Pathway_Attention"` //备注
}

//临床路径与临床路径项关系：Clinical_Relation
type Clinical_Relation struct {
	Clinical_Pathway_Num        string `json:"Clinical_Pathway_Num"`        //临床路径编号
	Clinical_Pathway_Item_Num   string `json:"Clinical_Pathway_Item_Num"`   //临床路径项编号
	Clinical_Relation_Attention string `json:"Clinical_Relation_Attention"` //备注
}

//临床路径执行情况
type Clinical_Pathway_Execu_Status struct {
	Clinical_Pathway_Execu_Num       string `json:"Clinical_Pathway_Execu_Num"`       //临床路径执行编号
	Clinical_Pathway_Item_Num        string `json:"Clinical_Pathway_Item_Num"`        //临床路径项编号
	Clinical_Pathway_Execu_Time      string `json:"Clinical_Pathway_Execu_Time"`      //执行时间
	Clinical_Pathway_Execu_Status    string `json:"Clinical_Pathway_Execu_Status"`    //完成情况
	Clinical_Pathway_Execu_Attention string `json:"Clinical_Pathway_Execu_Attention"` //备注

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
	if fun == "add_Patient_Basic_Info" {
		return sc.add_Patient_Basic_Info(stub, args)
	} else if fun == "query_Patient_Basic_Info" {
		return sc.query_Patient_Basic_Info(stub, args)
	} else if fun == "add_Resident_Info" {
		return sc.add_Resident_Info(stub, args)
	} else if fun == "query_Resident_Info" {
		return sc.query_Resident_Info(stub, args)
	} else if fun == "add_Record_Info" {
		return sc.add_Record_Info(stub, args)
	} else if fun == "query_Record_Info" {
		return sc.query_Record_Info(stub, args)
	} else if fun == "add_Clinical_Pathway" {
		return sc.add_Clinical_Pathway(stub, args)
	} else if fun == "query_Clinical_Pathway" {
		return sc.query_Clinical_Pathway(stub, args)
	} else if fun == "add_Clinical_Pathway_Item" {
		return sc.add_Clinical_Pathway_Item(stub, args)
	} else if fun == "query_Clinical_Pathway_Item" {
		return sc.query_Clinical_Pathway_Item(stub, args)
	} else if fun == "add_Patient_Clinical_Pathway" {
		return sc.add_Patient_Clinical_Pathway(stub, args)
	} else if fun == "query_Patient_Clinical_Pathway" {
		return sc.query_Patient_Clinical_Pathway(stub, args)
	} else if fun == "add_Clinical_Relation" {
		return sc.add_Clinical_Relation(stub, args)
	} else if fun == "query_Clinical_Relation" {
		return sc.query_Clinical_Relation(stub, args)
	} else if fun == "add_Clinical_Pathway_Execu_Status" {
		return sc.add_Clinical_Pathway_Execu_Status(stub, args)
	} else if fun == "query_Clinical_Pathway_Execu_Status" {
		return sc.query_Clinical_Pathway_Execu_Status(stub, args)
	}

	return shim.Success(nil)

}

//添加病人基本信息
func (sc *SimpleChaincode) add_Patient_Basic_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//查询病人基本信息
func (sc *SimpleChaincode) query_Patient_Basic_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//添加住院信息Resident_Info
func (sc *SimpleChaincode) add_Resident_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询住院信息
func (sc *SimpleChaincode) query_Resident_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加电子病历信息
func (sc *SimpleChaincode) add_Record_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询电子病历信息
func (sc *SimpleChaincode) query_Record_Info(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加临床路
func (sc *SimpleChaincode) add_Clinical_Pathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询临床路径
func (sc *SimpleChaincode) query_Clinical_Pathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加临床路径项
func (sc *SimpleChaincode) add_Clinical_Pathway_Item(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询临床路径项
func (sc *SimpleChaincode) query_Clinical_Pathway_Item(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加患者临床路径
func (sc *SimpleChaincode) add_Patient_Clinical_Pathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询患者临床路径
func (sc *SimpleChaincode) query_Patient_Clinical_Pathway(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加临床与临床路径项的关系
func (sc *SimpleChaincode) add_Clinical_Relation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询临床与临床路径项的关系
func (sc *SimpleChaincode) query_Clinical_Relation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//添加临床路径执行情况
func (sc *SimpleChaincode) add_Clinical_Pathway_Execu_Status(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//查询临床路径执行情况
func (sc *SimpleChaincode) query_Clinical_Pathway_Execu_Status(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
