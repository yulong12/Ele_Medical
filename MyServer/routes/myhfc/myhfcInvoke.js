//invokeV2.js
var hfc = require("fabric-client");
var path = require("path");
var util = require("util");
var sdkUtils = require("fabric-client/lib/utils");
var fs = require("fs");
var options = require("./org1Config");
// var channel = {};
// var client = null;
// var targets = [];
// var tx_id = null;
// var peer = null;
var getKeyFilesInDir = function(dir) {
  var files = fs.readdirSync(dir);
  var keyFiles = [];
  files.forEach(function(file_name) {
    var filePath = path.join(dir, file_name);
    if (file_name.endsWith("_sk")) {
      keyFiles.push(filePath);
    }
  });

  return keyFiles;
};
function postInvokeRequest(requestJson, getResponse) {
  var channel = {};
  var client = null;
  var targets = [];
  var tx_id = null;
  var peer = null;
  var responseJston = {};

  var fcn_rquest = requestJson;
  Promise.resolve()
    .then(function() {
      console.log("Load privateKey and signedCert");
      client = new hfc();
      var createUserOpt = {
        username: options.user_id,
        mspid: options.msp_id,
        cryptoContent: {
          privateKey: getKeyFilesInDir(options.privateKeyFolder)[0],
          signedCert: options.signedCert
        }
      };

      return sdkUtils
        .newKeyValueStore({
          path: "/tmp/fabric-client-stateStore"
        })
        .then(function(store) {
          client.setStateStore(store);
          return client.createUser(createUserOpt);
        });
    })
    .then(function(user) {
      channel = client.newChannel(options.channel_id);
      var data = fs.readFileSync(options.peer_tls_cacerts);
      peer = client.newPeer(options.peer_url, {
        pem: Buffer.from(data).toString(),
        "ssl-target-name-override": options.server_hostname
      });

      channel.addPeer(peer);
      var odata = fs.readFileSync(options.orderer_tls_cacerts);
      var caroots = Buffer.from(odata).toString();
      var orderer = client.newOrderer(options.orderer_url, {
        pem: caroots,
        "ssl-target-name-override": "orderer.example.com"
      });

      channel.addOrderer(orderer);
      targets.push(peer);
      return;
    })
    .then(function() {
      tx_id = client.newTransactionID();
      console.log("Assigning transaction_id :", tx_id._transaction_id);
      fcn_rquest.targets = targets;
      fcn_rquest.chaincodeId = options.chaincode_id;
      fcn_rquest.chainId = options.channel_id;
      fcn_rquest.txId = tx_id;
      var str = JSON.stringify(fcn_rquest);

      return channel.sendTransactionProposal(fcn_rquest);
    })
    .then(function(results) {
      var proposalResponses = results[0];
      var proposal = results[1];
      let isProposalGood = false;
      if (
        proposalResponses &&
        proposalResponses[0].response &&
        proposalResponses[0].response.status === 200
      ) {
        isProposalGood = true;
        console.log("Transaction proposal was good");
      } else {
        console.log(
          "Transaction proposal was bad" + proposalResponses[0].status
        );
        console.log(
          "Transaction proposal was bad" + proposalResponses[0].message
        );
        responseJston.status = proposalResponses[0].status;
        responseJston.payload = proposalResponses[0].message.toString();
      }
      if (isProposalGood) {
        console.log(
          util.format(
            'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"',
            proposalResponses[0].response.status,
            proposalResponses[0].response.message
          )
        );

        responseJston.status = proposalResponses[0].response.status;
        responseJston.payload = proposalResponses[0].response.payload.toString();
        console.log(
          "-----proposalResponses[0].response.payload.toString()------" +
            proposalResponses[0].response.payload.toString()
        );
        var responseJstonstr = JSON.stringify(responseJston);
        // console.log("-----responseJstonstr------" + responseJstonstr);
        // build up the request for the orderer to have the transaction committed
        var request = {
          proposalResponses: proposalResponses,
          proposal: proposal
        };

        // set the transaction listener and set a timeout of 30 sec
        // if the transaction did not get committed within the timeout period,
        // report a TIMEOUT status
        var transaction_id_string = tx_id.getTransactionID(); //Get the transaction ID string to be used by the event processing
        var promises = [];

        var sendPromise = channel.sendTransaction(request);
        promises.push(sendPromise); //we want the send transaction first, so that we know where to check status

        // get an eventhub once the fabric client has a user assigned. The user
        // is required bacause the event registration must be signed
        let event_hub = channel.newChannelEventHub(peer);

        // using resolve the promise so that result status may be processed
        // under the then clause rather than having the catch clause process
        // the status
        let txPromise = new Promise((resolve, reject) => {
          let handle = setTimeout(() => {
            event_hub.unregisterTxEvent(transaction_id_string);
            event_hub.disconnect();
            resolve({ event_status: "TIMEOUT" }); //we could use reject(new Error('Trnasaction did not complete within 30 seconds'));
          }, 3000);
          event_hub.registerTxEvent(
            transaction_id_string,
            (tx, code) => {
              // this is the callback for transaction event status
              // first some clean up of event listener
              clearTimeout(handle);

              // now let the application know what happened
              var return_status = {
                event_status: code,
                tx_id: transaction_id_string
              };
              if (code !== "VALID") {
                console.log("The transaction was invalid, code = " + code);
                resolve(return_status); // we could use reject(new Error('Problem with the tranaction, event status ::'+code));
              } else {
                console.log(
                  "The transaction has been committed on peer " +
                    event_hub.getPeerAddr()
                );
                resolve(return_status);
              }
            },
            err => {
              //this is the callback if something goes wrong with the event registration or processing
              reject(
                new Error("There was a problem with the eventhub ::" + err)
              );
            },
            { disconnect: true } //disconnect when complete
          );
          event_hub.connect();
        });
        promises.push(txPromise);

        return Promise.all(promises);
      }
    })
    .then(value => {
      getResponse(responseJston);
    });
}
module.exports = postInvokeRequest;
