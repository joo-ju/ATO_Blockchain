"use strict";

const { FileSystemWallet, Gateway } = require("fabric-network");
const path = require("path");

const ccpPath = path.resolve(__dirname, "..", "connection.json");

async function main() {
  try {
    const walletPath = path.join(process.cwd(), "..", "wallet");
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    const userExists = await wallet.exists("user1");
    if (!userExists) {
      console.log(
        'An identity for the user "user1" does not exist in the wallet'
      );
      console.log("Run the registUser.js application before retrying");
      return;
    }
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet,
      identity: "user1",
      discovery: { enabled: true, asLocalhost: true },
    });

    const network = await gateway.getNetwork("channelseller");

    const contract = network.getContract("ato-cc-2");

    var title = process.argv[2];
    var state = process.argv[3];
    var seller = process.argv[4];
    var category = process.argv[5];
    var price = process.argv[6];
    var content = process.argv[7];
    var walletid = process.argv[8];

    const result = await contract.submitTransaction(
      "setGoods",
      title,
      state,
      seller,
      category,
      price,
      content,
      walletid
    );
    console.log("Transaction has been submitted, result is:", result);

    await gateway.disconnect();
  } catch (error) {
    console.error(`Failed to submit transaction: ${error}`);
    process.exit(1);
  }
}

main();
