
# 1，患者基本信息：PatientBasicInfo
```

患者姓名：PatientName 
患者来源：PatientFrom
患者身份证号：PatientIdentityNumber  

患者年龄：PatientAge
患者电话：PatientTel 
患者性别:PatientSex 


患者民族：PatientNationality 
患者血型：PatientBloodGroup 
患者职业：PatientJob  

家庭地址：PatientAddress


患者亲友姓名：PatientRelativeName 
患者亲友电话：PatientRelativeTel 
与患者关系：PatientRelation
患者亲友身体状况：PatientRelationStatus

患者病史：PatientMedicalHistory  
患者吸烟史:PatientSmokeHistory 
患者家族遗传史：PatientFamilyGeneticHistory 
患者主治医生姓名：PatientDoctor
患者就诊科室：PatientdePartment

 ```


# 2, 住院患者表：ResidentInfo
```
 住院流水号：ResidentNum
 住院者姓名；ResidentName
 患者编号：ResidentPatientNum
 就诊科室：ResidentDepartment
 床位号：ResidentBedNum
 主管医生：ResidentSupervisor
主管护士：ResidentSupervisorNurse
住院日期：ResidentInDate
出院日期:ResidentOutDate
备注：ResidentAttention
```

# 3.电子病历信息：RecordInfo
```
文档编号：RecordNum
住院流水号：RecordResidentNum
文档名称:RecordName
创建医生:RecordCreateDoctor
创建时间：RecordCreateTime
文档路径:RecordPath
文档哈希值:RecordHashValue
备注：RecordAttention
```
# 4,临床路径：ClinicalPathway
```
临床路径编号：ClinicalPathwayNum
临床路径名称：ClinicalPathwayName
备注：ClinicalPathwayAttention
```

# 5, 临床路径项：ClinicalPathwayItem
```
临床路径项编号：ClinicalPathwayItemNum
临床路径项名称：ClinicalPathwayItemName
备注：ClinicalPathwayItemAttention

```
# 6,患者临床路径：PatientClinicalPathway
```

患者临床路径编号：PatientClinicalPathwayNum
临床路径编号：ClinicalPathwayNum
项目开始时间：PatientClinicalBeginTime
住院流水号：ResidentNum
创建医生：PatientClinicalPathwayCreator
备注：PatientClinicalPathwayAttention
```
# 7,临床路径与临床路径项关系：ClinicalRelation
```

临床路径编号：ClinicalPathwayNum
临床路径项编号：ClinicalPathwayItemNum
备注：ClinicalRelationAttention
```
# 8,临床路径执行情况：ClinicalPathwayExecuStatus
```

临床路径执行编号：ClinicalPathwayExecuNum
临床路径项编号：ClinicalPathwayItemNum
执行时间：ClinicalPathwayExecuTime
完成情况：ClinicalPathwayExecuStatus
备注：ClinicalPathwayExecuAttention

```






