// police
var org3Config = {
  user_id: "Admin@org2.example.com",
  msp_id: "Org2MSP",
  channel_id: "mychannel",
  chaincode_id: "mycc",
  peer_url: "grpcs://localhost:11051",
  event_url: "grpcs://localhost:11053",
  orderer_url: "grpcs://localhost:7050",
  network_url: "grpcs://localhost:11051",
  privateKeyFolder:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/fabric-samplesv1.3/first-network/crypto-config/peerOrganizations/org3.example.com/users/Admin@org1.example.com/msp/keystore",
  signedCert:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/fabric-samplesv1.3/first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem",
  peer_tls_cacerts:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/fabric-samplesv1.3/first-network/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt",
  orderer_tls_cacerts:
    "/Users/zhangyulong/Documents/gopath/src/github.com/hyperledger/fabric-samplesv1.3/first-network/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt",
  server_hostname: "peer0.org1.example.com"
};
module.exports = org3Config;
