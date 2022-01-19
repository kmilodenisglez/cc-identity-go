## how to use model-identity-go
The goal of model-identity-go repository is to share common data transfer object / structs definitions for other repos of the fuel HLF traceability project.

[follow the instructions](https://github.com/ic-matcom/model-identity-go/blob/main/README.md)



## DID Sample
```bash
# DID examples:
did:aa43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71
did:aa53bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71
did:aa63bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71
did:aa73bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71

# ID examples:
4a57a199-9850-45d2-a46b-f4b1d7fc9958
```
## WORLD-STATE

http://192.168.49.131:5984/_utils/

credentials: admin/adminpw

## chaincode execution samples (CLI)

### InitLedger
```bash
# populate the ledger with first data 
peer chaincode invoke  -c '{"function":"org.identity:InitLedger","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```

### populating with fake data
```bash
# populate issuer for dev (arg: none)
peer chaincode invoke  -c '{"function":"org.identity:OnlyDevIssuer","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# populate access and role for dev (arg: none)
peer chaincode invoke  -c '{"function":"org.identity:OnlyDevAccess","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# populate identity for dev (arg: none)
peer chaincode invoke  -c '{"function":"org.identity:OnlyDevParticipant","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

 
```
### interact with the access and role transactions
```bash
# before must invoke InitLedger

# GetAccesses (arg: none)
peer chaincode query -c '{"function":"org.identity:GetAccesses","Args":["{\"pageSize\":3,\"bookmark\":\"\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
peer chaincode query  -c  '{"function":"org.identity:QueryAssetsWithPagination","Args":["{\"queryString\":{\"selector\":{\"docType\":\"did.access\"}},\"pageSize\":3,\"bookmark\":\"\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetAccess (arg: model.GetRequest)
peer chaincode query -c '{"function":"org.identity:GetAccess","Args":["{\"id\":\"access-id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# CreateRole (arg: model.RoleCreateRequest)
# before must invoke InitLedger
peer chaincode invoke -c '{"function":"org.identity:CreateRole","Args":["{\"name\":\"Rol de prueba\",\"contractFunctions\":[\"GetAccess\"]}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetRoles (arg: none)
peer chaincode query -c '{"function":"org.identity:GetRoles","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetRole (arg: model.GetRequest)
peer chaincode query -c '{"function":"org.identity:GetRole","Args":["{\"id\":\"role-id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# DeleteRole (arg: model.GetRequest)
peer chaincode invoke -c '{"function":"org.identity:DeleteRole","Args":["{\"id\":\"role-id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```

### interact with the issuer transactions
```bash
# CreateIssuer (arg: model.IssuerCreateRequest)
peer chaincode invoke -c '{"function":"org.identity:CreateIssuer","Args":["{\"name\":\"Autoridad de Certificación Tecnomática\",\"certPem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlLNWpDQ0JzNmdBd0lCQWdJR0FJdXl5WEFCTUEwR0NTcUdTSWIzRFFFQkRRVUFNSUg0TVNVd0l3WUpLb1pJDQpodmNOQVFrQkZoWmhaRzF2Ym5CcmFVQnRZV2xzTG0xdUxtTnZMbU4xTVFzd0NRWURWUVFHRXdKRFZURVNNQkFHDQpBMVVFQ0F3SlRHRWdTR0ZpWVc1aE1SQXdEZ1lEVlFRSERBZENiM2xsY205ek1VTXdRUVlEVlFRS0REcEpibVp5DQpZV1Z6ZEhKMVkzUjFjbUVnWkdVZ1RHeGhkbVVnVU1PNllteHBZMkVnWkdVZ2JHRWdVbVZ3dzdwaWJHbGpZU0JrDQpaU0JEZFdKaE1SZ3dGZ1lEVlFRTERBOUJkWFJ2Y21sa1lXUWdVbUhEclhveFBUQTdCZ05WQkFNTU5FRjFkRzl5DQphV1JoWkNCa1pTQkRaWEowYVdacFkyRmphY096YmlCVFpYSjJhV05wYnlCRFpXNTBjbUZzSUVOcFpuSmhaRzh3DQpIaGNOTWpFd01qQXpNVFF3T1RBMldoY05Namt3TWpBeE1UUXdPVEEyV2pDQnBURUxNQWtHQTFVRUJoTUNRMVV4DQpFakFRQmdOVkJBZ01DVXhoSUVoaFltRnVZVEVXTUJRR0ExVUVCd3dOUTJWdWRISnZJRWhoWW1GdVlURW5NQ1VHDQpBMVVFQ2d3ZVRXbHVhWE4wWlhKcGJ5QmtaU0JGYm1WeVo4T3RZU0I1SUUxcGJtRnpNUTR3REFZRFZRUUxEQVZEDQpkWEJsZERFeE1DOEdBMVVFQXd3b1FYVjBiM0pwWkdGa0lHUmxJRU5sY25ScFptbGpZV05wdzdOdUlGUmxZMjV2DQpiY09oZEdsallUQ0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQQURDQ0Fnb0NnZ0lCQUpGWVV5Y253N015DQpTUENZMjMwTUhEYXRNek16YzEvK3NYcWhRbU1KQXg3T0kxL0ZUNzBzMmxZQzNrd3hLSnVha2Qzc1ZlQ0plVHptDQpyQUtFdDR5VTdUWkRDbCt0ckFYMjdhKytCNWc4a2h5OXJ1aFNOYUFqMXlOekRMRXZoSC9VNytHOHV1YlBDc1FxDQpWZ29nc01IaFFqd0hnR3lublRjNkVvWDNZZ2c0RHYycXp2c1lwa1pURlNMMzB1NC9RYTVETVlxNm0wSVlMUDJCDQpYbFJUL3FCSmQ0N0RkOUg3QWR5MHFvQWRkTkpuWWwrdnVhcC8vYWRuSDFQbHE3TGQ5UUw5R2NITGw5SUxkQkJ2DQpsVFJESWZTL1Vpc1l2cVV4Rm52K29aSHZBaDhNS0RyZzFBSnpEQWlSc041MDFtQVdaNlh1aFd1V2pReUp2ZnI0DQp6bnR3TEFTUUloaEJiZkVaSDdDd3FKSGZWa3g0U01ORWhkNlNqQjI0NElqelQ2NFpBNXZOcW8zRE03dnZweElrDQo2ZWtINXdKUzdLeXlvdkY4NGJHbStFT2I5TVBwTlA1NnlzSiszcnl1R2dua0EvUmE4SlBpd0dFTHZnZVovSmxRDQpVN0wxU0xxbHhaekdzVitzUkZFcWVaZGRLU1lhbEZaMFZmUFowcStrNXBnZ0xSYkRDL0oxeDJreUlDRHVtQk8zDQpHYlpGYldQQ3B0cjhwVjZHL1J0T0VZcnNzaEN0SlZzT3lXQ2dPNWRaT0dGSTUrVTN1NTZ3UHgxeUxIVHJuTEQ5DQppclcybm92YURKd0tFRVJRUHJYV0FCZFhoZjJJclBkcWdwcnVOTkhpbTJmVXgzamNNU3VwajFES3YyNVBjbG5HDQpZVDE4N0VjYkNCTWpHRmNUc3A5UWZOWHBQbm9qU2ZWNUFnTUJBQUdqZ2dMRk1JSUN3VEFQQmdOVkhSTUJBZjhFDQpCVEFEQVFIL01CMEdBMVVkRGdRV0JCUkpGQ3ZkYUxkd25kVytGWXFRRWhqaFJ4RWttVENDQVNzR0ExVWRJd1NDDQpBU0l3Z2dFZWdCUUttYUxtY1diZDZkSmhBY1BORitrOGgyTWVrYUdCL3FTQit6Q0IrREVsTUNNR0NTcUdTSWIzDQpEUUVKQVJZV1lXUnRiMjV3YTJsQWJXRnBiQzV0Ymk1amJ5NWpkVEVMTUFrR0ExVUVCaE1DUTFVeEVqQVFCZ05WDQpCQWdNQ1V4aElFaGhZbUZ1WVRFUU1BNEdBMVVFQnd3SFFtOTVaWEp2Y3pGRE1FRUdBMVVFQ2d3NlNXNW1jbUZsDQpjM1J5ZFdOMGRYSmhJR1JsSUV4c1lYWmxJRkREdW1Kc2FXTmhJR1JsSUd4aElGSmxjTU82WW14cFkyRWdaR1VnDQpRM1ZpWVRFWU1CWUdBMVVFQ3d3UFFYVjBiM0pwWkdGa0lGSmh3NjE2TVQwd093WURWUVFERERSQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybkRzMjRnVTJWeWRtbGphVzhnUTJWdWRISmhiQ0JEYVdaeVlXUnZnZ1VDDQpWQXZrQVRBT0JnTlZIUThCQWY4RUJBTUNBWVl3UXdZSUt3WUJCUVVIQVFFRU56QTFNRE1HQ0NzR0FRVUZCekFCDQpoaWRvZEhSd09pOHZiMk56Y0M1elpYSmpaVzVqYVdZdVkzVXZkbUV2YzNSaGRIVnpMMjlqYzNBd1J3WURWUjBmDQpCRUF3UGpBOG9EcWdPSVkyYUhSMGNEb3ZMMk55YkM1elpYSmpaVzVqYVdZdVkzVXZkbUV2WTNKc2N5OXpaV0Z5DQpZMmd1WTJkcFAyRnNhV0Z6UFVGRFUwTkRNRDRHQTFVZElBUTNNRFV3TXdZRFZSMGdNQ3d3S2dZSUt3WUJCUVVIDQpBZ0VXSG1oMGRIQTZMeTl6WlhKalpXNWphV1l1Ylc0dVkzVXZaSEJqTG1Sdll6Q0JnUVlKWUlaSUFZYjRRZ0VODQpCSFFXY2tObGNuUnBabWxqWVdSdklFUnBaMmwwWVd3Z1IyVnVaWEpoWkc4Z2NHRnlZU0JzWVNCQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybnpiaUJKYm5SbGNtMWxaR2xoT2lCQmRYUnZjbWxrWVdRZ1pHVWdRMlZ5DQpkR2xtYVdOaFkybnpiaUJVWldOdWIyM2hkR2xqWVRBTkJna3Foa2lHOXcwQkFRMEZBQU9DQkFFQVpIR0hjVnozDQpBRWZoUVVRK0loOXFkSVkzVTVET2wwYXB0SjI2U0F4bkE2MjhBNm15SGlxdlFKa2N4VVYrSXY1c1hqN1lpMnpRDQpvR1BSMVJMbHIvMU1weExJbitNdExkUGt3ZE94elVsVk5FT2x5SVJJb0lJdmNIdGc5ZTZYblNhc1hVM1E4OFBqDQpsMFhxQlR0Q0Q3dldFRWhTbDFhZkJHLzJsYkF1a1VLcEJPWllMb2RWRDF6MGhBM0lpcUord29rd0tmU3BTUWxMDQpldk11emFMYXpZUmU5Sk9xaURSTkhLN1ZndGRKa0haZmtDWlE3QkFrd3o3ZkthaG1JdUFUSCtqM2Iyc3czM0xLDQppbmx6Um5ITW5XL0hUZzhpRnczc2hKbFh3UGpXUGJOQjBycUFWSS9rdDRMa2pvL2lLK21tMzdDTWRkamFSMHE2DQpyUHF5emdkczZXdzlmdnFRYVVsMXcyazJmMkExckRhVHlsZjAxQlhPV0lsc0ErdGF5WVVUeGlpZ2cxcmE1RUdTDQpYRU9Fc1NHRkJ3VUs5WnFHT2ljbW9iek1LeWFUdU5QQ1JhV1NZcXA0dzFvYlkyWk11Vkg0Zm50QVlpR2ZhbmtSDQpZRk1IYjVIVm5Ud2UyaDNTbG9jVkJyS3NsWkZubFlvREFRbVVSN0YvR3pFRFdqanU0OC9CY0Mybm5CTGt4NTVtDQp1VUdVS0M4NTlodHpOVXBBRm5icUwxUmUyMXBLRUx2SlJjdDRkQy8vUlNIbzlCS3duUi9tMTkzSlZOeEJCelc1DQp6dkZ5S2w0UGtkWTJ5em1pTlFqc251VE9EZVZZQjYrV2pObmtLdnlyK0d5WFRPYUd0LzRIR3Fpb1p6Y0FWTjh1DQpvN3VFWjF4bmNFWnZLN2hnak02TmZPOEFMSGNlcy9oekR5TEc5MVR1eTlRbnJwYWFndHFUamNKSGJvUVZoVmJtDQpaOHdHc256bVpwNDVWdFU3L3FTZVhkTUR4M25VdDR1QlE4RHZBZ1l0TFEyakdvUUhyOUIybVI3TkFzbG1XWlJUDQo0cFJlbWdMZjUyM1hOb1pUVC96dWQ4YWNER09oN3Y4TTQyT0tBcEpDV0h0YzRpcHZYREx5anZkd2xZUkNWWElIDQpqUnp3UWZVSUZsK2V6Rm53ajdrUG1yWEtjT2dyRkk1NXlKSTN1TDB0TDNPek04WG4zMjdxYm1yMi9zeHh1ekl1DQpRaldyT3pNMWV0VVJKNGttU29oV2dQeFdHZmRQK3NIdVNyNGs0Qmkza3NjeVplNkw3d3BDWmsxVjhQV3JlNjZUDQpBZzRGbC8vRXN2cnZydFBXQjlrQ3JSN0VVYk03ZUFWVWJoaEhhTi90alBQQ1VSd3NXU2V6a3NhK3BDWUowOWUwDQpydjVTN05JMG5kUjZJV3VERm5IeEJuL3hrdGZwRG9PL0svcW5hOW1kSDVKQ0c4eVAvdjhkbkR3NmdFTzVWNnBPDQpnVGNiQ1M3R0J2TkRKMnpGWi9iaWx1L1JYNDJSaUM5MmlpSDczSkt2bDRGVkJ6T0NRWWZ0SDNwQlFDMnU5eUVqDQprSHVpakFBRlRjRFVRNFczL1g4VGhNWmZ1MGZQY2NMdW1la2x2VmVTNCtNY3Zaci9Uc2ZlK2haTXZtaCt3OGdmDQpSdE9PNDJYck9sSFMvR0FOOXRMUFBNOWxUWmEvQmdNZzZJY25KMnRGUjNuNElrOS9Tb2xJbEtaT2ZDZFdtRUFGDQpwSHdVVjRQNFoxdVgrcVFOTDFVaFA3SmRQQktrQ09JTEgvdHZJc1p2OEpzYXlRRVNGcWgzcHVMaTg5ZlNTMTV1DQpQMWk0YWhLSWQyY01pZz09DQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tDQo=\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetIssuerHistory (arg: model.GetRequest)
peer chaincode query  -c '{"function":"org.identity:GetIssuerHistory","Args":["{\"id\":\"1d8c4f4d-ce32-4263-83cc-a76739c29469\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetIssuers  (arg: none)
peer chaincode query -c '{"function":"org.identity:GetIssuers","Args":[]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetIssuer (arg: model.GetRequest)
peer chaincode query -c '{"function":"org.identity:GetIssuer","Args":["{\"id\":\"issuer-id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# DeleteIssuer (arg: model.GetRequest)
peer chaincode invoke -c '{"function":"org.identity:DeleteIssuer","Args":["{\"id\":\"issuer-id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# RenewIssuer (arg: model.IssuerUpdateRequest)
peer chaincode invoke  -c '{"function":"org.identity:RenewIssuer","Args":["{\"id\":\"1d8c4f4d-ce32-4263-83cc-a76739c29469\",\"certPem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlLNWpDQ0JzNmdBd0lCQWdJR0FJdXl5WEFCTUEwR0NTcUdTSWIzRFFFQkRRVUFNSUg0TVNVd0l3WUpLb1pJDQpodmNOQVFrQkZoWmhaRzF2Ym5CcmFVQnRZV2xzTG0xdUxtTnZMbU4xTVFzd0NRWURWUVFHRXdKRFZURVNNQkFHDQpBMVVFQ0F3SlRHRWdTR0ZpWVc1aE1SQXdEZ1lEVlFRSERBZENiM2xsY205ek1VTXdRUVlEVlFRS0REcEpibVp5DQpZV1Z6ZEhKMVkzUjFjbUVnWkdVZ1RHeGhkbVVnVU1PNllteHBZMkVnWkdVZ2JHRWdVbVZ3dzdwaWJHbGpZU0JrDQpaU0JEZFdKaE1SZ3dGZ1lEVlFRTERBOUJkWFJ2Y21sa1lXUWdVbUhEclhveFBUQTdCZ05WQkFNTU5FRjFkRzl5DQphV1JoWkNCa1pTQkRaWEowYVdacFkyRmphY096YmlCVFpYSjJhV05wYnlCRFpXNTBjbUZzSUVOcFpuSmhaRzh3DQpIaGNOTWpFd01qQXpNVFF3T1RBMldoY05Namt3TWpBeE1UUXdPVEEyV2pDQnBURUxNQWtHQTFVRUJoTUNRMVV4DQpFakFRQmdOVkJBZ01DVXhoSUVoaFltRnVZVEVXTUJRR0ExVUVCd3dOUTJWdWRISnZJRWhoWW1GdVlURW5NQ1VHDQpBMVVFQ2d3ZVRXbHVhWE4wWlhKcGJ5QmtaU0JGYm1WeVo4T3RZU0I1SUUxcGJtRnpNUTR3REFZRFZRUUxEQVZEDQpkWEJsZERFeE1DOEdBMVVFQXd3b1FYVjBiM0pwWkdGa0lHUmxJRU5sY25ScFptbGpZV05wdzdOdUlGUmxZMjV2DQpiY09oZEdsallUQ0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQQURDQ0Fnb0NnZ0lCQUpGWVV5Y253N015DQpTUENZMjMwTUhEYXRNek16YzEvK3NYcWhRbU1KQXg3T0kxL0ZUNzBzMmxZQzNrd3hLSnVha2Qzc1ZlQ0plVHptDQpyQUtFdDR5VTdUWkRDbCt0ckFYMjdhKytCNWc4a2h5OXJ1aFNOYUFqMXlOekRMRXZoSC9VNytHOHV1YlBDc1FxDQpWZ29nc01IaFFqd0hnR3lublRjNkVvWDNZZ2c0RHYycXp2c1lwa1pURlNMMzB1NC9RYTVETVlxNm0wSVlMUDJCDQpYbFJUL3FCSmQ0N0RkOUg3QWR5MHFvQWRkTkpuWWwrdnVhcC8vYWRuSDFQbHE3TGQ5UUw5R2NITGw5SUxkQkJ2DQpsVFJESWZTL1Vpc1l2cVV4Rm52K29aSHZBaDhNS0RyZzFBSnpEQWlSc041MDFtQVdaNlh1aFd1V2pReUp2ZnI0DQp6bnR3TEFTUUloaEJiZkVaSDdDd3FKSGZWa3g0U01ORWhkNlNqQjI0NElqelQ2NFpBNXZOcW8zRE03dnZweElrDQo2ZWtINXdKUzdLeXlvdkY4NGJHbStFT2I5TVBwTlA1NnlzSiszcnl1R2dua0EvUmE4SlBpd0dFTHZnZVovSmxRDQpVN0wxU0xxbHhaekdzVitzUkZFcWVaZGRLU1lhbEZaMFZmUFowcStrNXBnZ0xSYkRDL0oxeDJreUlDRHVtQk8zDQpHYlpGYldQQ3B0cjhwVjZHL1J0T0VZcnNzaEN0SlZzT3lXQ2dPNWRaT0dGSTUrVTN1NTZ3UHgxeUxIVHJuTEQ5DQppclcybm92YURKd0tFRVJRUHJYV0FCZFhoZjJJclBkcWdwcnVOTkhpbTJmVXgzamNNU3VwajFES3YyNVBjbG5HDQpZVDE4N0VjYkNCTWpHRmNUc3A5UWZOWHBQbm9qU2ZWNUFnTUJBQUdqZ2dMRk1JSUN3VEFQQmdOVkhSTUJBZjhFDQpCVEFEQVFIL01CMEdBMVVkRGdRV0JCUkpGQ3ZkYUxkd25kVytGWXFRRWhqaFJ4RWttVENDQVNzR0ExVWRJd1NDDQpBU0l3Z2dFZWdCUUttYUxtY1diZDZkSmhBY1BORitrOGgyTWVrYUdCL3FTQit6Q0IrREVsTUNNR0NTcUdTSWIzDQpEUUVKQVJZV1lXUnRiMjV3YTJsQWJXRnBiQzV0Ymk1amJ5NWpkVEVMTUFrR0ExVUVCaE1DUTFVeEVqQVFCZ05WDQpCQWdNQ1V4aElFaGhZbUZ1WVRFUU1BNEdBMVVFQnd3SFFtOTVaWEp2Y3pGRE1FRUdBMVVFQ2d3NlNXNW1jbUZsDQpjM1J5ZFdOMGRYSmhJR1JsSUV4c1lYWmxJRkREdW1Kc2FXTmhJR1JsSUd4aElGSmxjTU82WW14cFkyRWdaR1VnDQpRM1ZpWVRFWU1CWUdBMVVFQ3d3UFFYVjBiM0pwWkdGa0lGSmh3NjE2TVQwd093WURWUVFERERSQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybkRzMjRnVTJWeWRtbGphVzhnUTJWdWRISmhiQ0JEYVdaeVlXUnZnZ1VDDQpWQXZrQVRBT0JnTlZIUThCQWY4RUJBTUNBWVl3UXdZSUt3WUJCUVVIQVFFRU56QTFNRE1HQ0NzR0FRVUZCekFCDQpoaWRvZEhSd09pOHZiMk56Y0M1elpYSmpaVzVqYVdZdVkzVXZkbUV2YzNSaGRIVnpMMjlqYzNBd1J3WURWUjBmDQpCRUF3UGpBOG9EcWdPSVkyYUhSMGNEb3ZMMk55YkM1elpYSmpaVzVqYVdZdVkzVXZkbUV2WTNKc2N5OXpaV0Z5DQpZMmd1WTJkcFAyRnNhV0Z6UFVGRFUwTkRNRDRHQTFVZElBUTNNRFV3TXdZRFZSMGdNQ3d3S2dZSUt3WUJCUVVIDQpBZ0VXSG1oMGRIQTZMeTl6WlhKalpXNWphV1l1Ylc0dVkzVXZaSEJqTG1Sdll6Q0JnUVlKWUlaSUFZYjRRZ0VODQpCSFFXY2tObGNuUnBabWxqWVdSdklFUnBaMmwwWVd3Z1IyVnVaWEpoWkc4Z2NHRnlZU0JzWVNCQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybnpiaUJKYm5SbGNtMWxaR2xoT2lCQmRYUnZjbWxrWVdRZ1pHVWdRMlZ5DQpkR2xtYVdOaFkybnpiaUJVWldOdWIyM2hkR2xqWVRBTkJna3Foa2lHOXcwQkFRMEZBQU9DQkFFQVpIR0hjVnozDQpBRWZoUVVRK0loOXFkSVkzVTVET2wwYXB0SjI2U0F4bkE2MjhBNm15SGlxdlFKa2N4VVYrSXY1c1hqN1lpMnpRDQpvR1BSMVJMbHIvMU1weExJbitNdExkUGt3ZE94elVsVk5FT2x5SVJJb0lJdmNIdGc5ZTZYblNhc1hVM1E4OFBqDQpsMFhxQlR0Q0Q3dldFRWhTbDFhZkJHLzJsYkF1a1VLcEJPWllMb2RWRDF6MGhBM0lpcUord29rd0tmU3BTUWxMDQpldk11emFMYXpZUmU5Sk9xaURSTkhLN1ZndGRKa0haZmtDWlE3QkFrd3o3ZkthaG1JdUFUSCtqM2Iyc3czM0xLDQppbmx6Um5ITW5XL0hUZzhpRnczc2hKbFh3UGpXUGJOQjBycUFWSS9rdDRMa2pvL2lLK21tMzdDTWRkamFSMHE2DQpyUHF5emdkczZXdzlmdnFRYVVsMXcyazJmMkExckRhVHlsZjAxQlhPV0lsc0ErdGF5WVVUeGlpZ2cxcmE1RUdTDQpYRU9Fc1NHRkJ3VUs5WnFHT2ljbW9iek1LeWFUdU5QQ1JhV1NZcXA0dzFvYlkyWk11Vkg0Zm50QVlpR2ZhbmtSDQpZRk1IYjVIVm5Ud2UyaDNTbG9jVkJyS3NsWkZubFlvREFRbVVSN0YvR3pFRFdqanU0OC9CY0Mybm5CTGt4NTVtDQp1VUdVS0M4NTlodHpOVXBBRm5icUwxUmUyMXBLRUx2SlJjdDRkQy8vUlNIbzlCS3duUi9tMTkzSlZOeEJCelc1DQp6dkZ5S2w0UGtkWTJ5em1pTlFqc251VE9EZVZZQjYrV2pObmtLdnlyK0d5WFRPYUd0LzRIR3Fpb1p6Y0FWTjh1DQpvN3VFWjF4bmNFWnZLN2hnak02TmZPOEFMSGNlcy9oekR5TEc5MVR1eTlRbnJwYWFndHFUamNKSGJvUVZoVmJtDQpaOHdHc256bVpwNDVWdFU3L3FTZVhkTUR4M25VdDR1QlE4RHZBZ1l0TFEyakdvUUhyOUIybVI3TkFzbG1XWlJUDQo0cFJlbWdMZjUyM1hOb1pUVC96dWQ4YWNER09oN3Y4TTQyT0tBcEpDV0h0YzRpcHZYREx5anZkd2xZUkNWWElIDQpqUnp3UWZVSUZsK2V6Rm53ajdrUG1yWEtjT2dyRkk1NXlKSTN1TDB0TDNPek04WG4zMjdxYm1yMi9zeHh1ekl1DQpRaldyT3pNMWV0VVJKNGttU29oV2dQeFdHZmRQK3NIdVNyNGs0Qmkza3NjeVplNkw3d3BDWmsxVjhQV3JlNjZUDQpBZzRGbC8vRXN2cnZydFBXQjlrQ3JSN0VVYk03ZUFWVWJoaEhhTi90alBQQ1VSd3NXU2V6a3NhK3BDWUowOWUwDQpydjVTN05JMG5kUjZJV3VERm5IeEJuL3hrdGZwRG9PL0svcW5hOW1kSDVKQ0c4eVAvdjhkbkR3NmdFTzVWNnBPDQpnVGNiQ1M3R0J2TkRKMnpGWi9iaWx1L1JYNDJSaUM5MmlpSDczSkt2bDRGVkJ6T0NRWWZ0SDNwQlFDMnU5eUVqDQprSHVpakFBRlRjRFVRNFczL1g4VGhNWmZ1MGZQY2NMdW1la2x2VmVTNCtNY3Zaci9Uc2ZlK2haTXZtaCt3OGdmDQpSdE9PNDJYck9sSFMvR0FOOXRMUFBNOWxUWmEvQmdNZzZJY25KMnRGUjNuNElrOS9Tb2xJbEtaT2ZDZFdtRUFGDQpwSHdVVjRQNFoxdVgrcVFOTDFVaFA3SmRQQktrQ09JTEgvdHZJc1p2OEpzYXlRRVNGcWgzcHVMaTg5ZlNTMTV1DQpQMWk0YWhLSWQyY01pZz09DQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tDQo=\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```

### interact with the identity transactions
```bash

# CreateParticipant with role (arg: model.ParticipantCreateRequest)
peer chaincode invoke  -c '{"function":"org.identity:CreateParticipant","Args":["{\"roles\":[\"06658e98-e829-4a54-8377-8f62ac4de4b7\"],\"publicKey\":\"my public key\",\"certPem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlHRWpDQ0EvcWdBd0lCQWdJVWRyd2M0NFhvK2ZLQ1JzTXlIdHd6ZDBnQ01iVXdEUVlKS29aSWh2Y05BUUVODQpCUUF3Z2FVeEN6QUpCZ05WQkFZVEFrTlZNUkl3RUFZRFZRUUlEQWxNWVNCSVlXSmhibUV4RmpBVUJnTlZCQWNNDQpEVU5sYm5SeWJ5QklZV0poYm1FeEp6QWxCZ05WQkFvTUhrMXBibWx6ZEdWeWFXOGdaR1VnUlc1bGNtZkRyV0VnDQplU0JOYVc1aGN6RU9NQXdHQTFVRUN3d0ZRM1Z3WlhReE1UQXZCZ05WQkFNTUtFRjFkRzl5YVdSaFpDQmtaU0JEDQpaWEowYVdacFkyRmphY096YmlCVVpXTnViMjNEb1hScFkyRXdIaGNOTWpFd056TXdNVGN6TlRJM1doY05Nak13DQpOek13TVRjek5USTJXakNCcWpFYk1Ca0dDZ21TSm9tVDhpeGtBUUVNQ3pnM01EVXhNakV4TkRVM01SNHdIQVlEDQpWUVFEREJWWmFYTmxiQ0JCYzNScFlYcGhjbUZwYmlCRWFXNHhIVEFiQmdOVkJBd01GRVZ6Y0M0Z1FpQkRhV1Z1DQpZMmxoY3lCSmJtWXVNUlF3RWdZRFZRUUxEQXREZFhCbGRDMU5hVzVsYlRFVk1CTUdBMVVFQ2d3TVZHVmpibTl0DQp3NkYwYVdOaE1SSXdFQVlEVlFRSURBbE1ZU0JJWVdKaGJtRXhDekFKQmdOVkJBWVRBa05WTUZrd0V3WUhLb1pJDQp6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVDNVNwZTlRem12WldVcnBLMHo0bDJVYjVwcVczZEs4OXlzV1k3d0xHDQpUMldybjFwSHFLckpHM0NXdFl6QVllaW9aRlA1bENJTjdHUE5ycXNlWUY1S2thT0NBZnd3Z2dINE1Bd0dBMVVkDQpFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVU1JRcjNXaTNjSjNWdmhXS2tCSVk0VWNSSkprd1ZRWUlLd1lCDQpCUVVIQVFFRVNUQkhNQ01HQ0NzR0FRVUZCekFDaGhkb2RIUndjem92TDNCcmFTNWpkWEJsZEM1amRTOWpZVEFnDQpCZ2dyQmdFRkJRY3dBWVlVYUhSMGNEb3ZMMjlqYzNBdVkzVndaWFF1WTNVd0tRWURWUjB1QkNJd0lEQWVvQnlnDQpHb1lZYUhSMGNEb3ZMMlJsYkhSaFkzSnNMbU4xY0dWMExtTjFNRDRHQTFVZEpRUTNNRFVHQ0NzR0FRVUZCd01DDQpCZ2dyQmdFRkJRY0RBd1lJS3dZQkJRVUhBd1FHQ2lzR0FRUUJnamNLQXd3R0NTcUdTSWIzTHdFQkJUQ0IxUVlEDQpWUjBmQklITk1JSEtNSUhIb0JlZ0ZZWVRhSFIwY0RvdkwyTnliQzVqZFhCbGRDNWpkYUtCcTZTQnFEQ0JwVEV4DQpNQzhHQTFVRUF3d29RWFYwYjNKcFpHRmtJR1JsSUVObGNuUnBabWxqWVdOcHc3TnVJRlJsWTI1dmJjT2hkR2xqDQpZVEVPTUF3R0ExVUVDd3dGUTNWd1pYUXhKekFsQmdOVkJBb01IazFwYm1semRHVnlhVzhnWkdVZ1JXNWxjbWZEDQpyV0VnZVNCTmFXNWhjekVXTUJRR0ExVUVCd3dOUTJWdWRISnZJRWhoWW1GdVlURVNNQkFHQTFVRUNBd0pUR0VnDQpTR0ZpWVc1aE1Rc3dDUVlEVlFRR0V3SkRWVEFkQmdOVkhRNEVGZ1FVbkZXUFFRVkpDSGt3STNCUDVMSFVHMy9sDQozVjB3RGdZRFZSMFBBUUgvQkFRREFnWGdNQTBHQ1NxR1NJYjNEUUVCRFFVQUE0SUNBUUE5TTBGZlVMZVcySEkrDQpNNFZDNEZRczZybDZ0dmJ6NHFRalZ3MWZDbnNNVVNxZDBUejB0eTdWdGxCcGJteXFhNnRqSitFYmxncUVGOTBTDQpseHlnN1NyY1RkMGxWVFlIaExURUhhWFAzWENCVkhRUDVhVUIxQmZYR0pkNGJwaEZDUk5pQ1ExajhXa0xUTy8vDQo3bWNJSW9vSVh5Zk93K1N5d085VzFhZUt2amV5OEVrVTQ5bUtIakZXbC92Ritic0NPQUNwK1dCeGZQNFgrbm9yDQowdGVVR0MrZGZiYkY3OTM5c2tTdEU2SEwvTzRRT3RqeWZXWFZVLzhDNWRjMlpHMTJiSzZXOE9Hc294bTJyMFkrDQpSVkhOdWNPemhTTHJEL3B6ZDBLS1BVSE5ma1NmM3dBenlJYzM4amVreWx3ZFlHdnk5VWV0dER2NFJ2cFYrSE9BDQo2M1FUVDlFbmlvQkNQVHZTbTVMb2Q2eVpsd0xLbzBNakdJZXNmNG9tU3RsSWxIcUE0ZUppL2V5dzdxZmdZdTk4DQptY2t6dVNhVWlQVE1YaUgwWC84dDFlVEkzRHpjVEo0Uytmc251WEd0OGs5WEFWa2w4elNzM2xROERNVG1UVll3DQpTSG1ERm5hQ0psbDRlazJUNlpOWVB6dmFnOWdBVGxBQUMzSEw4ZGlycXdLN2FJeURveENPSHU3a0JJQU0xK202DQpDb0ZPcnVmalZvVVRtTzlUWUlocUlXZWJOMWY5Z3hXSkVBbnF6Zm4wUkVSa3NTMU0zV0JnaGZWY0xoeDY3WEd3DQpzTXltdHZLQ1hWRkpYSmNJaW5kRTUyYVAyQzVnQ1NVU1VyVVJOcTYvRjNqUVB1UE1ySW1ydHo1ZWoyV0tKVmhBDQpuS0hseFN2UkdYTXNuSDVoVmpibWZBU3B0cU9CQ3c9PQ0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ0K\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetParticipant (arg: model.ParticipantGetRequest)
peer chaincode query  -c '{"function":"org.identity:GetParticipant","Args":["{\"did\":\"did:ab43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec72\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetParticipants (arg: none)
peer chaincode query  -c  '{"function":"org.identity:QueryAssetsWithPagination","Args":["{\"queryString\":{\"selector\":{\"docType\":\"did.participant\"}},\"pageSize\":3,\"bookmark\":\"\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
peer chaincode query -c '{"function":"org.identity:GetParticipants","Args":["{\"pageSize\":3,\"bookmark\":\"\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetParticipantHistory (arg: model.ParticipantGetRequest)
peer chaincode query -c '{"function":"org.identity:GetParticipantHistory","Args":["{\"did\":\"did:fa3bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# DeleteParticipant (arg: model.ParticipantDeleteRequest)
peer chaincode invoke  -c '{"function":"org.identity:DeleteParticipant","Args":["{\"userDid\":\"did:ab43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec72\",\"callerDid\":\"did:ab43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec72\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# GetBatch
peer chaincode query -c '{"function":"org.tecnomatica.fuelbatch:GetBatch","Args":["{\"id\":\"did:aa43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71\",\"payload\":{\"id\":\"my id\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```

### Rich Queries without Pagination
```bash
# query "participant" by name in attrs field
peer chaincode query  -c  '{"function":"org.identity:QueryAssetsBy","Args":["{\"selector\":{\"docType\":\"did.participant\",\"attrs.name\":\"Yisel Astiazarain Din\"},\"use_index\":[\"indexParticipantAttrsDoc\",\"indexParticipantAttrs\"]}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# query "participant" by id (indexes)
peer chaincode query  -c  '{"function":"org.identity:QueryAssetsBy","Args":["{\"selector\":{\"docType\":\"did.participant\",\"id\":\"98a530e9-986a-437d-a8f8-ed0978374205\"},\"use_index\":[\"indexParticipantIDDoc\",\"indexParticipantID\"]}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

### Rich Queries with Pagination
# query "participant" by name in attrs field
# searching for all participants named "Yisel"
peer chaincode query  -c  '{"function":"org.identity:QueryAssetsWithPagination","Args":["{\"queryString\":{\"selector\":{\"docType\":\"did.participant\",\"attrs.name\":{\"$gt\":\"Yise\"}}},\"pageSize\":3,\"bookmark\":\"\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```


["{\"queryString\":{\"selector\":{\"docType\":\"did.participant\",\"attrs.name\":{\"$gt\":\"Yise\"}}},\"pageSize\":3,\"bookmark\":\"\"}"]

peer chaincode invoke  -c '{"function":"org.tecnomatica.fuelbatch:Instantiate","Args":["{\"id\":\"id\",\"payload\":\"{}\",\"signature\":\"sign\"}"]}' -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA -C $CHANNEL_NAME -n $CC_NAME --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

"Args":["{\"Id\":\"did:aa43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71\",\"Function\":\"\"}"]}'

"Args":["{\"id\":\"did:aa43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71\",\"payload\":\"\"}"]}'

'{"function":"org.tecnomatica.fuelbatch:GetBatch","Args":["{\"Id\":\"did:aa43bdf5b4bcfac88ce9093ec3f0d58290f11c7ef6d2a683a7ee56746b333ec71\",\"payload\":\"{}\"}"]}'




