// ip config
var ipConfig = {
  // org1
  org1_peer_network: "grpcs://localhost:7051", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  org1_event: "grpcs://localhost:7053", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  // org2
  org2_peer_network: "grpcs://localhost:8051", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  org2_event: "grpcs://localhost:8053", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  // org3
  org3_peer_network: "grpcs://localhost:9051", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  org3_event: "grpcs://localhost:9053", // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
  orderer_url: "grpcs://localhost:7050" // 因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc
};
module.exports = ipConfig;
