#!/usr/bin/env sh

# carpeta donde estan los binarios de fabric (peer, orderer, etc.)
export PATH=/mnt/d/DEVELOPMENT/WORKSPACE_BCN/fabric-bin-tmp/bin:"$PATH"
export ORDERER_ADDRESS=127.0.0.1:6050
export ORDERER_TLS_CA=/mnt/d/DEVELOPMENT/WORKSPACE_BCN/fabric-bin-tmp/fabric-testnet-nano-without-syschannel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt
export CHANNEL_NAME=mychannel
# recuerda modificar el CHAINCODE_ID por el retornado por el peer install
export CHAINCODE_ID=ccidentity_1.0:9f0f53be089000e5d999b56fc0f5311807b1a8dc61067d0e9c214077a3cf6471
export CHAINCODE_SERVER_ADDRESS=127.0.0.1:9999
export CC_NAME=identity

export CORE_PEER_TLS_ROOTCERT_FILE=/mnt/d/DEVELOPMENT/WORKSPACE_BCN/fabric-bin-tmp/fabric-testnet-nano-without-syschannel/crypto-config/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
export CORE_CHAINCODE_ID_NAME=$CHAINCODE_ID