# 采用dev模式检测chaincode


## 1.激活

```
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./medical

peer chaincode install -p chaincodedev/chaincode/medical -n mycc -v 0

peer chaincode instantiate -n mycc -v 0 -c '{"Args":[]}' -C myc
```

## 2，添加和查询患者基本信息
```

peer chaincode invoke -n mycc -c '{"Args":["addPatientBasicInfo", "101010101010", "Bob","130521199903090776","woman","19940308","han","O","","Policemen","110","Ada","120","no","no","no","no","no","Chronic pharyngitis","two years","no"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["queryPatientBasicInfo", "101010101010"]}' -C myc
```
## 3,添加和查询住院信息
```
peer chaincode invoke -n mycc -c '{"Args":["addResidentInfo", "11111111111111","Bob","101010101010","surgical","3B01","SuSan","Ali","20190101","20190108","Love Clean"]}' -C myc


peer chaincode invoke -n mycc -c '{"Args":["queryResidentInfo", "11111111111111"]}' -C myc
```

## 4,添加和查询病历信息
```

peer chaincode invoke -n mycc -c '{"Args":["addRecordInfo", "22222222222222","11111111111111","Medical Record","SuSan","20190108","/user/record","sdafasdfasdfadsfadsfas","no"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["queryRecordInfo", "22222222222222"]}' -C myc
```

## 5,添加和查询临床路径
```

peer chaincode invoke -n mycc -c '{"Args":["addClinicalPathway", "33333333333333","testClinicalPathway","testClinicalPathway"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["queryClinicalPathway", "33333333333333"]}' -C myc
```
## 6,添加和查询临床路径项
```
peer chaincode invoke -n mycc -c '{"Args":["addClinicalPathwayItem", "44444444444444","testClinicalPathwayItem","testClinicalPathwayItem"]}' -C myc


peer chaincode invoke -n mycc -c '{"Args":["queryClinicalPathwayItem", "44444444444444"]}' -C myc

```
## 7,添加和查询患者临床路径
```
peer chaincode invoke -n mycc -c '{"Args":["addPatientClinicalPathway", "55555555555555","33333333333333","20190101","11111111111111","SuSan","Rember"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["queryPatientClinicalPathway", "55555555555555"]}' -C myc
```
## 8,添加和查询临床和临床路径关系
```

peer chaincode invoke -n mycc -c '{"Args":["addClinicalRelation", "66666666666666","44444444444444","atttention"]}' -C myc


peer chaincode invoke -n mycc -c '{"Args":["queryClinicalRelation", "66666666666666"]}' -C myc
```
### 9,添加和查询临床路径执行情况

```

peer chaincode invoke -n mycc -c '{"Args":["addClinicalPathwayExecuStatus", "77777777777777","44444444444444","20190101","good","attention"]}' -C myc


peer chaincode invoke -n mycc -c '{"Args":["queryClinicalPathwayExecuStatus", "77777777777777"]}' -C myc

```