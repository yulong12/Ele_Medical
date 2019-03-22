# 采用 dev 模式测试 chincode

#### 使用 fabric-sample 中的 chaincode-docker-devmode 来测试

#### 详见：https://hyperledger-fabric.readthedocs.io/en/release-1.3/chaincode4ade.html

- 1 配置环境

```
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./medical
```

- 2 安装 chaincode

```
peer chaincode install -p chaincodedev/chaincode/medical -n mycc -v 0
```

- 3 激活 chaincode

```
  peer chaincode instantiate -n mycc -v 0 -c '{"Args":[""]}' -C myc
```

- 4 测试函数 uploadRecordData，上传病历

```
  peer chaincode invoke -n mycc -c '{"Args":["uploadRecordData", "patientNo","DocterNo", "HisNo","recordNo","recordName","recordPath","recordSize","recordHash"]}' -C myc
```

- 5 测试函数 saveHospitalized，保存住院病历

```
peer chaincode invoke -n mycc -c '{"Args":["saveHospitalized", "name", "age","phone","idCard","sex","address","doctor","nurse","illness","treatment","medication","attention","room","inTime","outTime","cost"]}' -C myc
```

- 6 测试函数 queryHospitalized，查询住院病历

```
peer chaincode invoke -n mycc -c '{"Args":["queryHospitalized", "idCard"]}' -C myc
```

- 7 测试函数 applyRemoteData，远程请求病历

```
  peer chaincode invoke -n mycc -c '{"Args":["applyRemoteData", "TargetNo","RedicalNO","Applier","ApplierNo","ApplierHisNo"]}' -C myc
```

- 8 测试函数 queryRecordData，查询病历

```
  peer chaincode invoke -n mycc -c '{"Args":["queryRecordData", "patientNo"]}' -C myc
```

##使用本地环境测试

```
peer chaincode instantiate -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -l golang -v 1.0 -c '{"Args":[""]}' -P 'OR ('\''Org1MSP.peer'\'','\''Org2MSP.peer'\'','\''Org3MSP.peer'\'')'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["uploadRecordData", "patientNo","DocterNo", "HisNo","recordNo","recordName","recordPath","recordSize","recordHash"]}'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["applyRemoteData","TargetNo","RedicalNO","Applier","ApplierNo","ApplierHisNo"]}'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["saveHospitalized", "name", "age","phone","idCard","sex","address","doctor","nurse","illness","treatment","medication","attention","room","inTime","outTime","cost"]}'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["queryHospitalized","idCard"]}'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["queryRecordData", "patientNo"]}'
```
