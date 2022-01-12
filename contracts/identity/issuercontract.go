package identity

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	model "github.com/ic-matcom/model-identity-go/model"
	modeltools "github.com/ic-matcom/model-identity-go/tools"
	"log"
)

// OnlyDevIssuer [temporary] function to populate with test data
func (ic *ContractIdentity) OnlyDevIssuer(ctx contractapi.TransactionContextInterface) (string, error) {
	log.Printf("[%s][OnlyDevIssuer:Populating]", ctx.GetStub().GetChannelID())
	const b64RootCertTecnomatica = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlLNWpDQ0JzNmdBd0lCQWdJR0FJdXl5WEFCTUEwR0NTcUdTSWIzRFFFQkRRVUFNSUg0TVNVd0l3WUpLb1pJDQpodmNOQVFrQkZoWmhaRzF2Ym5CcmFVQ\nnRZV2xzTG0xdUxtTnZMbU4xTVFzd0NRWURWUVFHRXdKRFZURVNNQkFHDQpBMVVFQ0F3SlRHRWdTR0ZpWVc1aE1SQXdEZ1lEVlFRSERBZENiM2xsY205ek1VTXdRUVlEVlFRS0REcEpibVp5DQpZV1Z6ZEhKMVkzUjFj\nbUVnWkdVZ1RHeGhkbVVnVU1PNllteHBZMkVnWkdVZ2JHRWdVbVZ3dzdwaWJHbGpZU0JrDQpaU0JEZFdKaE1SZ3dGZ1lEVlFRTERBOUJkWFJ2Y21sa1lXUWdVbUhEclhveFBUQTdCZ05WQkFNTU5FRjFkRzl5DQphV1J\noWkNCa1pTQkRaWEowYVdacFkyRmphY096YmlCVFpYSjJhV05wYnlCRFpXNTBjbUZzSUVOcFpuSmhaRzh3DQpIaGNOTWpFd01qQXpNVFF3T1RBMldoY05Namt3TWpBeE1UUXdPVEEyV2pDQnBURUxNQWtHQTFVRUJoTU\nNRMVV4DQpFakFRQmdOVkJBZ01DVXhoSUVoaFltRnVZVEVXTUJRR0ExVUVCd3dOUTJWdWRISnZJRWhoWW1GdVlURW5NQ1VHDQpBMVVFQ2d3ZVRXbHVhWE4wWlhKcGJ5QmtaU0JGYm1WeVo4T3RZU0I1SUUxcGJtRnpNU\nTR3REFZRFZRUUxEQVZEDQpkWEJsZERFeE1DOEdBMVVFQXd3b1FYVjBiM0pwWkdGa0lHUmxJRU5sY25ScFptbGpZV05wdzdOdUlGUmxZMjV2DQpiY09oZEdsallUQ0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQ\nQURDQ0Fnb0NnZ0lCQUpGWVV5Y253N015DQpTUENZMjMwTUhEYXRNek16YzEvK3NYcWhRbU1KQXg3T0kxL0ZUNzBzMmxZQzNrd3hLSnVha2Qzc1ZlQ0plVHptDQpyQUtFdDR5VTdUWkRDbCt0ckFYMjdhKytCNWc4a2h\n5OXJ1aFNOYUFqMXlOekRMRXZoSC9VNytHOHV1YlBDc1FxDQpWZ29nc01IaFFqd0hnR3lublRjNkVvWDNZZ2c0RHYycXp2c1lwa1pURlNMMzB1NC9RYTVETVlxNm0wSVlMUDJCDQpYbFJUL3FCSmQ0N0RkOUg3QWR5MH\nFvQWRkTkpuWWwrdnVhcC8vYWRuSDFQbHE3TGQ5UUw5R2NITGw5SUxkQkJ2DQpsVFJESWZTL1Vpc1l2cVV4Rm52K29aSHZBaDhNS0RyZzFBSnpEQWlSc041MDFtQVdaNlh1aFd1V2pReUp2ZnI0DQp6bnR3TEFTUUloa\nEJiZkVaSDdDd3FKSGZWa3g0U01ORWhkNlNqQjI0NElqelQ2NFpBNXZOcW8zRE03dnZweElrDQo2ZWtINXdKUzdLeXlvdkY4NGJHbStFT2I5TVBwTlA1NnlzSiszcnl1R2dua0EvUmE4SlBpd0dFTHZnZVovSmxRDQpV\nN0wxU0xxbHhaekdzVitzUkZFcWVaZGRLU1lhbEZaMFZmUFowcStrNXBnZ0xSYkRDL0oxeDJreUlDRHVtQk8zDQpHYlpGYldQQ3B0cjhwVjZHL1J0T0VZcnNzaEN0SlZzT3lXQ2dPNWRaT0dGSTUrVTN1NTZ3UHgxeUx\nIVHJuTEQ5DQppclcybm92YURKd0tFRVJRUHJYV0FCZFhoZjJJclBkcWdwcnVOTkhpbTJmVXgzamNNU3VwajFES3YyNVBjbG5HDQpZVDE4N0VjYkNCTWpHRmNUc3A5UWZOWHBQbm9qU2ZWNUFnTUJBQUdqZ2dMRk1JSU\nN3VEFQQmdOVkhSTUJBZjhFDQpCVEFEQVFIL01CMEdBMVVkRGdRV0JCUkpGQ3ZkYUxkd25kVytGWXFRRWhqaFJ4RWttVENDQVNzR0ExVWRJd1NDDQpBU0l3Z2dFZWdCUUttYUxtY1diZDZkSmhBY1BORitrOGgyTWVrY\nUdCL3FTQit6Q0IrREVsTUNNR0NTcUdTSWIzDQpEUUVKQVJZV1lXUnRiMjV3YTJsQWJXRnBiQzV0Ymk1amJ5NWpkVEVMTUFrR0ExVUVCaE1DUTFVeEVqQVFCZ05WDQpCQWdNQ1V4aElFaGhZbUZ1WVRFUU1BNEdBMVVF\nQnd3SFFtOTVaWEp2Y3pGRE1FRUdBMVVFQ2d3NlNXNW1jbUZsDQpjM1J5ZFdOMGRYSmhJR1JsSUV4c1lYWmxJRkREdW1Kc2FXTmhJR1JsSUd4aElGSmxjTU82WW14cFkyRWdaR1VnDQpRM1ZpWVRFWU1CWUdBMVVFQ3d\n3UFFYVjBiM0pwWkdGa0lGSmh3NjE2TVQwd093WURWUVFERERSQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybkRzMjRnVTJWeWRtbGphVzhnUTJWdWRISmhiQ0JEYVdaeVlXUnZnZ1VDDQpWQXZrQVRBT0\nJnTlZIUThCQWY4RUJBTUNBWVl3UXdZSUt3WUJCUVVIQVFFRU56QTFNRE1HQ0NzR0FRVUZCekFCDQpoaWRvZEhSd09pOHZiMk56Y0M1elpYSmpaVzVqYVdZdVkzVXZkbUV2YzNSaGRIVnpMMjlqYzNBd1J3WURWUjBmD\nQpCRUF3UGpBOG9EcWdPSVkyYUhSMGNEb3ZMMk55YkM1elpYSmpaVzVqYVdZdVkzVXZkbUV2WTNKc2N5OXpaV0Z5DQpZMmd1WTJkcFAyRnNhV0Z6UFVGRFUwTkRNRDRHQTFVZElBUTNNRFV3TXdZRFZSMGdNQ3d3S2dZ\nSUt3WUJCUVVIDQpBZ0VXSG1oMGRIQTZMeTl6WlhKalpXNWphV1l1Ylc0dVkzVXZaSEJqTG1Sdll6Q0JnUVlKWUlaSUFZYjRRZ0VODQpCSFFXY2tObGNuUnBabWxqWVdSdklFUnBaMmwwWVd3Z1IyVnVaWEpoWkc4Z2N\nHRnlZU0JzWVNCQmRYUnZjbWxrDQpZV1FnWkdVZ1EyVnlkR2xtYVdOaFkybnpiaUJKYm5SbGNtMWxaR2xoT2lCQmRYUnZjbWxrWVdRZ1pHVWdRMlZ5DQpkR2xtYVdOaFkybnpiaUJVWldOdWIyM2hkR2xqWVRBTkJna3\nFoa2lHOXcwQkFRMEZBQU9DQkFFQVpIR0hjVnozDQpBRWZoUVVRK0loOXFkSVkzVTVET2wwYXB0SjI2U0F4bkE2MjhBNm15SGlxdlFKa2N4VVYrSXY1c1hqN1lpMnpRDQpvR1BSMVJMbHIvMU1weExJbitNdExkUGt3Z\nE94elVsVk5FT2x5SVJJb0lJdmNIdGc5ZTZYblNhc1hVM1E4OFBqDQpsMFhxQlR0Q0Q3dldFRWhTbDFhZkJHLzJsYkF1a1VLcEJPWllMb2RWRDF6MGhBM0lpcUord29rd0tmU3BTUWxMDQpldk11emFMYXpZUmU5Sk9x\naURSTkhLN1ZndGRKa0haZmtDWlE3QkFrd3o3ZkthaG1JdUFUSCtqM2Iyc3czM0xLDQppbmx6Um5ITW5XL0hUZzhpRnczc2hKbFh3UGpXUGJOQjBycUFWSS9rdDRMa2pvL2lLK21tMzdDTWRkamFSMHE2DQpyUHF5emd\nkczZXdzlmdnFRYVVsMXcyazJmMkExckRhVHlsZjAxQlhPV0lsc0ErdGF5WVVUeGlpZ2cxcmE1RUdTDQpYRU9Fc1NHRkJ3VUs5WnFHT2ljbW9iek1LeWFUdU5QQ1JhV1NZcXA0dzFvYlkyWk11Vkg0Zm50QVlpR2Zhbm\ntSDQpZRk1IYjVIVm5Ud2UyaDNTbG9jVkJyS3NsWkZubFlvREFRbVVSN0YvR3pFRFdqanU0OC9CY0Mybm5CTGt4NTVtDQp1VUdVS0M4NTlodHpOVXBBRm5icUwxUmUyMXBLRUx2SlJjdDRkQy8vUlNIbzlCS3duUi9tM\nTkzSlZOeEJCelc1DQp6dkZ5S2w0UGtkWTJ5em1pTlFqc251VE9EZVZZQjYrV2pObmtLdnlyK0d5WFRPYUd0LzRIR3Fpb1p6Y0FWTjh1DQpvN3VFWjF4bmNFWnZLN2hnak02TmZPOEFMSGNlcy9oekR5TEc5MVR1eTlR\nbnJwYWFndHFUamNKSGJvUVZoVmJtDQpaOHdHc256bVpwNDVWdFU3L3FTZVhkTUR4M25VdDR1QlE4RHZBZ1l0TFEyakdvUUhyOUIybVI3TkFzbG1XWlJUDQo0cFJlbWdMZjUyM1hOb1pUVC96dWQ4YWNER09oN3Y4TTQ\nyT0tBcEpDV0h0YzRpcHZYREx5anZkd2xZUkNWWElIDQpqUnp3UWZVSUZsK2V6Rm53ajdrUG1yWEtjT2dyRkk1NXlKSTN1TDB0TDNPek04WG4zMjdxYm1yMi9zeHh1ekl1DQpRaldyT3pNMWV0VVJKNGttU29oV2dQeF\ndHZmRQK3NIdVNyNGs0Qmkza3NjeVplNkw3d3BDWmsxVjhQV3JlNjZUDQpBZzRGbC8vRXN2cnZydFBXQjlrQ3JSN0VVYk03ZUFWVWJoaEhhTi90alBQQ1VSd3NXU2V6a3NhK3BDWUowOWUwDQpydjVTN05JMG5kUjZJV\n3VERm5IeEJuL3hrdGZwRG9PL0svcW5hOW1kSDVKQ0c4eVAvdjhkbkR3NmdFTzVWNnBPDQpnVGNiQ1M3R0J2TkRKMnpGWi9iaWx1L1JYNDJSaUM5MmlpSDczSkt2bDRGVkJ6T0NRWWZ0SDNwQlFDMnU5eUVqDQprSHVp\nakFBRlRjRFVRNFczL1g4VGhNWmZ1MGZQY2NMdW1la2x2VmVTNCtNY3Zaci9Uc2ZlK2haTXZtaCt3OGdmDQpSdE9PNDJYck9sSFMvR0FOOXRMUFBNOWxUWmEvQmdNZzZJY25KMnRGUjNuNElrOS9Tb2xJbEtaT2ZDZFd\ntRUFGDQpwSHdVVjRQNFoxdVgrcVFOTDFVaFA3SmRQQktrQ09JTEgvdHZJc1p2OEpzYXlRRVNGcWgzcHVMaTg5ZlNTMTV1DQpQMWk0YWhLSWQyY01pZz09DQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tDQo="
	// insert issuer
	issuerRequest := model.IssuerCreateRequest{
		CertPem: b64RootCertTecnomatica,
	}
	issuerActual, err := ic.CreateIssuer(ctx, issuerRequest)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return issuerActual.ID, nil
}

// CreateIssuer create in the ledger the issuer's certificate with its attributes
//
// Arguments:
//		0: IssuerCreateRequest
// Returns:
//		0: Issuer
//		1: error
func (ic *ContractIdentity) CreateIssuer(ctx contractapi.TransactionContextInterface, issuerRequest model.IssuerCreateRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][CreateIssuer]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	exist, err := lus.CertificateAlreadyExists(ctx, issuerRequest.CertPem, IssuerDocType, []string{})
	if err != nil {
		return nil, err
	} else if exist {
		return nil, fmt.Errorf("an issuer with the same certificate already exists")
	}

	issuerID := lus.GenerateUUID()

	// If the new issuer is to be the default, we must check if there is
	// another issuer registered and marked "by default"
	if err := ic.issuerMarkAsDefault(ctx, issuerID, &issuerRequest); err != nil {
		return nil, err
	}

	// decode certificate
	certX509, err := lus.GetX509CertFromPem(issuerRequest.CertPem)
	if err != nil {
		return nil, err
	}
	// we insert commonName if the Name field is empty
	commonName := certX509.Subject.CommonName
	if issuerRequest.Name == "" {
		issuerRequest.Name = commonName
	}

	attrs := modeltools.GetAttrsCert(certX509)
	// get dates
	dateCert := lus.GetDateCertificate(certX509)

	// Create Issuer
	issuer := &model.Issuer{
		DocType:     IssuerDocType,
		ID:          issuerID,
		Name:        issuerRequest.Name,
		CertPem:     issuerRequest.CertPem,
		Attrs:       attrs,
		AttrsExtras: make(map[string]string),
		IssuedTime:  dateCert["issuedTime"],
		ExpiresTime: dateCert["expiresTime"],
		ByDefault:   issuerRequest.ByDefault,
	}

	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerID})
	if err != nil {
		return nil, err
	}

	// JSON encoding of issuer
	issuerJE, _ := json.Marshal(issuer)
	if err := ctx.GetStub().PutState(key, issuerJE); err != nil {
		return nil, fmt.Errorf("issuer %s could not be created: %v", issuerRequest.Name, err)
	}

	return &model.IssuerQueryResponse{
		ID:          issuer.ID,
		Name:        issuer.Name,
		PublicKey:   issuerRequest.PublicKey, // TODO: obtener pubkey del certificado, si no existe el del request
		Attrs:       issuer.Attrs,
		AttrsExtras: issuer.AttrsExtras,
		IssuedTime:  issuer.IssuedTime,
		ExpiresTime: issuer.ExpiresTime,
		Active:      issuer.Active,
		ByDefault:   issuer.ByDefault,
	}, nil
}

// IssuerUpdateRequest
type IssuerUpdateRequest struct {
	ID      string `json:"id"`
	CertPem string `json:"certPem"` // b64 certificate PEM active
}

// RenewIssuer update issuer certificate in the ledger
//
// Arguments:
//		0: IssuerUpdateRequest
// Returns:
//		0: Issuer
//		1: error
func (ic *ContractIdentity) RenewIssuer(ctx contractapi.TransactionContextInterface, issuerRequest IssuerUpdateRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][UpdateIssuer]", ctx.GetStub().GetChannelID())

	// check if client-node connected as admin
	if err := lus.AssertAdmin(ctx); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// get issuer
	issuerToUpdate, err := ic.GetIssuer(ctx, model.GetRequest{ID: issuerRequest.ID})
	if err != nil {
		return nil, fmt.Errorf("failed to get User identity: %v", err)
	}
	if issuerToUpdate == nil {
		return nil, fmt.Errorf("%s does not exist", issuerToUpdate.ID)
	}

	// decode certificate
	certX509, err := lus.GetX509CertFromPem(issuerRequest.CertPem)
	if err != nil {
		return nil, err
	}
	// we insert commonName if the Name field is empty
	commonName := certX509.Subject.CommonName

	attrs := modeltools.GetAttrsCert(certX509)
	// get dates
	dateCert := lus.GetDateCertificate(certX509)

	// Create Issuer
	issuerToUpdate.Name = commonName
	issuerToUpdate.PublicKey = issuerRequest.CertPem
	issuerToUpdate.Attrs = model.Attrs(attrs)
	issuerToUpdate.IssuedTime = dateCert["issuedTime"]
	issuerToUpdate.ExpiresTime = dateCert["expiresTime"]

	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerToUpdate.ID})
	if err != nil {
		return nil, err
	}

	// JSON encoding of issuer
	issuerJE, _ := json.Marshal(issuerToUpdate)
	if err := ctx.GetStub().PutState(key, issuerJE); err != nil {
		return nil, fmt.Errorf("issuer %s could not be created: %v", commonName, err)
	}
	return issuerToUpdate, nil
}

// GetIssuer get an issuer from the ledger
//
// Arguments:
//		0: GetRequest
// Returns:
//		0: Issuer
//		1: error
func (ic *ContractIdentity) GetIssuer(ctx contractapi.TransactionContextInterface, request model.GetRequest) (*model.IssuerQueryResponse, error) {
	log.Printf("[%s][GetIssuer]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{request.ID})
	if err != nil {
		return nil, fmt.Errorf("error happened creating key for: %v", err)
	} else if key == "" {
		return nil, fmt.Errorf("no state found for %s", key)
	}
	issuer, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %v", err)
	} else if issuer == nil {
		return nil, fmt.Errorf("no state found for %s", key)
	}
	var issuerJD model.Issuer
	err = json.Unmarshal(issuer, &issuerJD)
	if err != nil {
		return nil, err
	}

	return &model.IssuerQueryResponse{
		ID:          issuerJD.ID,
		Name:        issuerJD.Name,
		PublicKey:   issuerJD.CertPem,
		Attrs:       issuerJD.Attrs,
		AttrsExtras: issuerJD.AttrsExtras,
		IssuedTime:  issuerJD.IssuedTime,
		ExpiresTime: issuerJD.ExpiresTime,
		Active:      issuerJD.Active,
		ByDefault:   issuerJD.ByDefault,
	}, nil
}

// DeleteIssuer delete an issuer from the ledger
//
// Arguments:
//		0: GetRequest
// Returns:
//		0: Issuer
//		1: error
func (ic *ContractIdentity) DeleteIssuer(ctx contractapi.TransactionContextInterface, issuerRequest model.GetRequest) error {
	log.Printf("[%s][DeleteIssuer]", ctx.GetStub().GetChannelID())
	key, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerRequest.ID})
	if err != nil {
		return fmt.Errorf("error happened creating key: %v", err)
	} else if key == "" {
		return fmt.Errorf("no state found for %s", key)
	}
	err = ctx.GetStub().DelState(key)
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}

	// validateDelState = false
	// We are not interested in validating error, it may be the case that the issuer does not have an index
	if err = lus.DeleteIndex(ctx.GetStub(), ObjectTypeIssuerByDefault, []string{issuerRequest.ID}, false); err != nil {
		return err
	}

	return nil
}

// GetIssuers get all issuer
//
// Arguments:
//		0: none
// Returns:
//		0: []Issuer
//		1: error
func (ic *ContractIdentity) GetIssuers(ctx contractapi.TransactionContextInterface) ([]model.IssuerQueryResponse, error) {
	log.Printf("[%s][GetIssuers]", ctx.GetStub().GetChannelID())

	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IssuerDocType, []string{})
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var items []model.IssuerQueryResponse
	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if responseRange == nil {
			return nil, err
		}

		var item model.IssuerQueryResponse
		err = json.Unmarshal(responseRange.Value, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil

	//return  (*model.IssuerQueryResponse)(unsafe.Pointer(&items)), nil
}

// GetIssuerHistory returns the chain of custody for a issuer since issuance
//
// Arguments:
//		0: model.GetRequest
// Returns:
//		0: []model.IssuerHistoryQueryResponse
//		1: error
func (ic *ContractIdentity) GetIssuerHistory(ctx contractapi.TransactionContextInterface, issuerRequest model.GetRequest) ([]model.IssuerHistoryQueryResponse, error) {
	log.Printf("GetIssuerHistory: ID %v", issuerRequest.ID)

	compositeKeyID, err := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{issuerRequest.ID})
	if err != nil {
		return nil, fmt.Errorf("error happened creating composite key for issuer: %v", err)
	} else if compositeKeyID == "" {
		return nil, fmt.Errorf("no state found for %s", compositeKeyID)
	}
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(compositeKeyID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []model.IssuerHistoryQueryResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		issuer := &model.IssuerQueryResponse{}
		issuer.ID = issuerRequest.ID
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &issuer)
			if err != nil {
				return nil, err
			}
		}

		timestamp := modeltools.GetTimestampRFC3339(response.Timestamp)

		record := model.IssuerHistoryQueryResponse{
			TxID:     response.TxId,
			Time:     timestamp,
			Record:   issuer,
			IsDelete: response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

// issuerMarkAsDefault modifies issuer  "byDefault" status to true only if request ByDefault field
// is true or if there are no issuers
func (ic *ContractIdentity) issuerMarkAsDefault(ctx contractapi.TransactionContextInterface, issuerID string, issuerRequest *model.IssuerCreateRequest) error {
	log.Printf("[inside][issuerMarkAsDefault]")

	issuerResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(ObjectTypeIssuerByDefault, []string{})
	if err != nil {
		return err
	}
	defer issuerResultsIterator.Close()
	thereIsAtLeastOne := issuerResultsIterator.HasNext()

	if thereIsAtLeastOne && issuerRequest.ByDefault {
		responseRange, err := issuerResultsIterator.Next()
		if responseRange == nil {
			return err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return err
		}

		if len(compositeKeyParts) == 1 {
			returnedIssuerID := compositeKeyParts[0]
			response, err := ic.GetIssuer(ctx, model.GetRequest{ID: returnedIssuerID})
			if err != nil {
				return err
			}
			response.ByDefault = false

			// struct to JSON
			responseJSON, err := json.Marshal(response)
			if err != nil {
				return fmt.Errorf("error happened marshalling the issuer: %v", err)
			}

			ledgerKey, _ := ctx.GetStub().CreateCompositeKey(IssuerDocType, []string{returnedIssuerID})
			// Put issuer to StateDB
			err = ctx.GetStub().PutState(ledgerKey, responseJSON)
			if err != nil {
				return fmt.Errorf("error happened updating the issuer: %v", err)
			}

			// remove index ObjectTypeIssuerByDefault
			err = ctx.GetStub().DelState(responseRange.Key)
			if err != nil {
				return fmt.Errorf("error happened delete the  old issuer index : %v", err)
			}
		}
	} else if thereIsAtLeastOne && !issuerRequest.ByDefault {
		return nil
	}

	if !thereIsAtLeastOne && !issuerRequest.ByDefault {
		issuerRequest.ByDefault = true
	}

	if err := lus.CreateIndex(ctx.GetStub(), ObjectTypeIssuerByDefault, []string{issuerID}); err != nil {
		return fmt.Errorf("could not create the byDefault index for the issuer %v: %v", issuerID, err)
	}

	return nil
}
