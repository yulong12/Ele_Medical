// police
var org1Config = {
  user_id: "Admin@org1.example.com",
  msp_id: "Org1MSP",
  channel_id: "mychannel",
  chaincode_id: "mycc",
  peer_url: "grpcs://localhost:7051",
  event_url: "grpcs://localhost:7053",
  orderer_url: "grpcs://localhost:7050",
  network_url: "grpcs://localhost:7051",
  privateKeyFolder:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore",
  signedCert:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem",
  peer_tls_cacerts:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/first-network/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt",
  orderer_tls_cacerts:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/Ele_Medical/first-network/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt",
  server_hostname: "peer0.org1.example.com"
};
module.exports = org1Config;
