
#### description
The identity certificate should be issue by just one Certificate Authority
Each company should be able to check/query attributes from a participant
A single DID should be enough to query all the information related to one identity, and his attributes
Only the company that certifies an attribute can edit it


The flow:
- Can register new identities
- Can delete identities
- Can disarm identities


#### import it as a module
```bash
go get github.com/ic-matcom/cc-identity-go
```

## 1. Running the chaincode as a service

Package and install the external chaincode on peer with the following simple commands:

```
tar cfz code.tar.gz connection.json
tar cfz external-chaincode.tgz metadata.json code.tar.gz

peer lifecycle chaincode install external-chaincode.tgz
```
Run the following command to query the package ID of the chaincode that you just installed:
```
peer lifecycle chaincode queryinstalled -o $ORDERER_ADDRESS --tlsRootCertFiles $ORDERER_TLS_CA
```

The command will return output similar to the following:
```bash
Chaincode code package identifier: ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5

# Installed chaincodes on peer:
Package ID: ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5
Label: ccidentity_1.0
```

Copy the returned chaincode package ID into an environment variable for use in subsequent commands (your ID may be different):

```bash
# in linux or darwin use export, ex:
export CHAINCODE_ID=ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5

# in windows use set, ex:
set CHAINCODE_ID=ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5
```

Build the chaincode:


```
# linux or darwin
go build -o ccass_identity_binary

# windows
go build -o ccass_identity_binary.exe
```

Set the chaincode server address:

```
# linux
export CHAINCODE_SERVER_ADDRESS=127.0.0.1:9999

# windows
set CHAINCODE_SERVER_ADDRESS=127.0.0.1:9999
```

And start the chaincode service:

```
# linux
./ccass_identity_binary

# windows
ccass_identity_binary.exe
```
## Activate the chaincode

approve and commit the chaincode:

```
peer lifecycle chaincode approveformyorg -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA --channelID $CHANNEL_NAME --name $CC_NAME --version 1 --package-id $CHAINCODE_ID --sequence 1

peer lifecycle chaincode commit -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA --channelID $CHANNEL_NAME --name $CC_NAME --version 1 --sequence 1
```
