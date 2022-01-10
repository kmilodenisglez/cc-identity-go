package identity

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"github.com/hyperledger/fabric-chaincode-go/pkg/attrmgr"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/ic-matcom/cc-identity-go/contracts/identity"
	lus "github.com/ic-matcom/cc-identity-go/lib-utils"
	"github.com/ic-matcom/cc-identity-go/testing"
	"github.com/ic-matcom/cc-identity-go/testing/mocks"
	"github.com/ic-matcom/cc-identity-go/testing/testcerts"
	model "github.com/ic-matcom/model-identity-go"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"time"
)

var (
	err          error
	org1MSP      []byte
	org1MSPAdmin []byte
	sc           = identity.ContractIdentity{}
)

var _ = ginkgo.Describe("Identity Smart Contract", func() {
	chaincodeStub := &mocks.ChaincodeStub{}
	clientIdentity := &mocks.ClientIdentity{}
	ctx := &mocks.TransactionContext{}
	ctx.GetStubReturns(chaincodeStub)
	ctx.GetClientIdentityReturns(clientIdentity)
	chaincodeStub.CreateCompositeKeyStub = testing.CreateComposeKey
	chaincodeStub.GetChannelIDReturns("channel")

	// load Central DC Root Cert
	testing.CertByteRootCentralDC, err = testcerts.Certificates[9].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load Tecnomatica Root Cert
	testing.CertByteRootTecnomatica, err = testcerts.Certificates[8].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load Root Cert
	testing.CertByteRoot, err = testcerts.Certificates[0].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load operator admin cert certificate with OU=admin
	testing.CertByteAdmin, err = testcerts.Certificates[1].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load user1 cert with attributes Cargo=Director de Calidad,
	// 	Email=user1@matcom.uh.cu,Nombre=Pedro Perez,Edad=32, etc.
	testing.CertByteUserWithAttrs, err = testcerts.Certificates[2].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load client cert
	testing.CertByteClient, err = testcerts.Certificates[3].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load cert expired with attrs
	testing.CertByteAttrsExpired, err = testcerts.Certificates[4].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load peer0 cert
	testing.CertBytePeer0, err = testcerts.Certificates[5].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load user2 cert, cert valid issued by unknown root-cert
	testing.CertByteUser2, err = testcerts.Certificates[6].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// load user3 cert, cert valid issued by Tecnomatica root-cert
	testing.CertByteUserYisel, err = testcerts.Certificates[7].CertBytes()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	ginkgo.BeforeEach(func() {
		org1MSP = testing.MarshalProtoOrPanic(&msp.SerializedIdentity{Mspid: testing.MspID, IdBytes: testing.CertByteClient})
		org1MSPAdmin = testing.MarshalProtoOrPanic(&msp.SerializedIdentity{Mspid: testing.MspID, IdBytes: testing.CertByteAdmin})
		// certificates with attributes
		chaincodeStub.GetCreatorReturns(org1MSP, nil)
	})

	ginkgo.It("certificate has expired", func() {
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteAttrsExpired)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		err = lus.HasExpired(cert)
		gomega.Expect(err).To(gomega.HaveOccurred())
	})

	ginkgo.It("encode peer0 certificate to base64", func() {
		certPeer0B64 := base64.StdEncoding.EncodeToString(testing.CertBytePeer0)
		gomega.Expect(certPeer0B64).To(gomega.BeIdenticalTo(testing.CertBase64Peer0))
	})

	ginkgo.It("decode peer0 certificate to byte", func() {
		certPeer0Decoded, err := base64.StdEncoding.DecodeString(testing.CertBase64Peer0)
		gomega.Expect(err).To(gomega.BeNil())
		_, err = lus.GetX509CertFromPemByte(certPeer0Decoded)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		equals := bytes.Equal(certPeer0Decoded, testing.CertBytePeer0)
		gomega.Expect(equals).To(gomega.BeTrue())
	})

	ginkgo.It("decode Tecnomatica Yisel certificate to byte", func() {
		certYiselDecoded, err := base64.StdEncoding.DecodeString(testing.CertBase64Yisel)
		gomega.Expect(err).To(gomega.BeNil())
		_, err = lus.GetX509CertFromPemByte(certYiselDecoded)
		gomega.Expect(err).To(gomega.BeNil())

		equals := bytes.Equal(certYiselDecoded, testing.CertByteUserYisel)
		gomega.Expect(equals).To(gomega.BeTrue())
	})

	ginkgo.It("verify Tecnomatica Root Cert attributes - using standard", func() {
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteRootTecnomatica)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		attrs := lus.GetAttrsCert(cert)

		gomega.Expect(attrs.Name).To(gomega.BeIdenticalTo("Autoridad de Certificación Tecnomática"))
		gomega.Expect(attrs.DNI).To(gomega.BeIdenticalTo(""))
		gomega.Expect(attrs.Company).To(gomega.BeIdenticalTo("Ministerio de Energía y Minas"))
		gomega.Expect(attrs.Position).To(gomega.BeIdenticalTo(""))
		gomega.Expect(attrs.Country).To(gomega.BeIdenticalTo("CU"))
		gomega.Expect(attrs.Province).To(gomega.BeIdenticalTo("La Habana"))
		gomega.Expect(attrs.Locality).To(gomega.BeIdenticalTo("Centro Habana"))
		gomega.Expect(attrs.OrganizationalUnit).To(gomega.BeIdenticalTo("Cupet"))
	})

	ginkgo.It("verify Tecnomatica Certificate attributes - using standard", func() {
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteUserYisel)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		attrs := lus.GetAttrsCert(cert)

		gomega.Expect(attrs.Name).To(gomega.BeIdenticalTo("Yisel Astiazarain Din"))
		gomega.Expect(attrs.DNI).To(gomega.BeIdenticalTo("87051211457"))
		gomega.Expect(attrs.Company).To(gomega.BeIdenticalTo("Tecnomática"))
		gomega.Expect(attrs.Position).To(gomega.BeIdenticalTo("Esp. B Ciencias Inf."))
		gomega.Expect(attrs.Country).To(gomega.BeIdenticalTo("CU"))
		gomega.Expect(attrs.Province).To(gomega.BeIdenticalTo("La Habana"))
		gomega.Expect(attrs.Locality).To(gomega.BeIdenticalTo(""))
		gomega.Expect(attrs.OrganizationalUnit).To(gomega.BeIdenticalTo("Cupet-Minem"))
	})

	ginkgo.It("verify certificate attributes - using non-standard X.509 extension asn1 1.2.3.4.5.6.7.8.1", func() {
		attrs, err := lus.GetAttrsNonStandardCert(testing.CertByteUserWithAttrs)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		gomega.Expect(attrs["Cargo"]).To(gomega.BeIdenticalTo("Director de Calidad"))
		gomega.Expect(attrs["Nombre"]).To(gomega.BeIdenticalTo("Pedro Perez"))
		gomega.Expect(attrs["Email"]).To(gomega.BeIdenticalTo("user1@matcom.uh.cu"))
		gomega.Expect(attrs["Edad"]).To(gomega.BeIdenticalTo("32"))
		gomega.Expect(attrs["NotExist"]).To(gomega.BeIdenticalTo(""))
	})

	ginkgo.It("verify certificate individual attributes - using non-standard X.509 extension asn1 1.2.3.4.5.6.7.8.1", func() {
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteUserWithAttrs)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		attrs, err := attrmgr.New().GetAttributesFromCert(cert)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		ok, err := lus.CheckAttr("AttrNotFound", "", attrs)
		gomega.Expect(err).To(gomega.HaveOccurred())
		gomega.Expect(ok).To(gomega.BeFalse())

		ok, err = lus.CheckAttr("Cargo", "Director de Calidad", attrs)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(ok).To(gomega.BeTrue())
	})

	ginkgo.It("certificate signed by another which isn't marked as a CA certificate root-cert", func() {
		block, _ := pem.Decode(testing.CertByteUser2)
		cert, err := x509.ParseCertificate(block.Bytes)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		_, err = cert.Verify(x509.VerifyOptions{CurrentTime: time.Time{}})
		invalid, ok := err.(x509.CertificateInvalidError)
		gomega.Expect(ok).To(gomega.BeFalse())
		gomega.Expect(invalid.Reason).Should(gomega.BeNumerically("==", x509.NotAuthorizedToSign))
	})

	ginkgo.It("check a fabric-cert issued by Fabric-Root Cert", func() {
		err = lus.CheckSignatureFrom(testing.CertByteRoot, testing.CertByteClient)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("check a certificate issued by Tecnomatica PKI", func() {
		err = lus.CheckSignatureFrom(testing.CertByteRootTecnomatica, testing.CertByteUserYisel)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("check attributes of Tecnomatica Root Cert", func() {
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteRootTecnomatica)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		attrs := lus.GetAttrsCert(cert)
		gomega.Expect(attrs.Name).To(gomega.BeIdenticalTo("Autoridad de Certificación Tecnomática"))
	})

	ginkgo.It("check a certificate issued by Root Cert", func() {
		err = lus.VerifyIssuedByRootCert(testing.CertByteRootTecnomatica, testing.CertByteUserYisel)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("create issuer", func() {
		// base64 encode - Tecnomatica Root Certificate -
		b64RootCertTecnomatica := base64.StdEncoding.EncodeToString(testing.CertByteRootTecnomatica)
		cert, err := lus.GetX509CertFromPemByte(testing.CertByteRootTecnomatica)
		dateCert := lus.GetDateCertificate(cert)
		gomega.Expect(err).To(gomega.BeNil())
		issuerRequest := model.IssuerCreateRequest{
			Name:    "Autoridad de Certificación Tecnomática",
			CertPem: b64RootCertTecnomatica,
		}

		// using client with admin cert
		chaincodeStub.GetCreatorReturns(org1MSPAdmin, nil)

		iterator := &mocks.StateQueryIterator{}
		iterator.HasNextReturnsOnCall(0, false)
		chaincodeStub.GetStateByPartialCompositeKeyReturns(iterator, nil)

		actualIssuer, err := sc.CreateIssuer(ctx, issuerRequest)
		gomega.Expect(err).To(gomega.BeNil())
		actualIssuerJSON, err := json.Marshal(actualIssuer)
		expectedIssuer := model.IssuerQueryResponse{
			//DocType: identity.IssuerDocType,
			ID:      actualIssuer.ID,
			Name:    "Autoridad de Certificación Tecnomática",
			CertPem: b64RootCertTecnomatica,
			Attrs: model.Attrs{
				Name:               "Autoridad de Certificación Tecnomática",
				DNI:                "",
				Position:           "",
				Country:            "CU",
				Company:            "Ministerio de Energía y Minas",
				Locality:           "Centro Habana",
				Province:           "La Habana",
				OrganizationalUnit: "Cupet",
			},
			AttrsExtras: map[string]string{},
			IssuedTime:  dateCert["issuedTime"],
			ExpiresTime: dateCert["expiresTime"],
			ByDefault:   true,
		}
		expectedIssuerJSON, err := json.Marshal(expectedIssuer)
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(expectedIssuerJSON).Should(gomega.MatchJSON(actualIssuerJSON))
	})

	ginkgo.It("get issuer", func() {
		// base64 encode
		b64RootCertTecnomatica := base64.StdEncoding.EncodeToString(testing.CertByteRootTecnomatica)

		issuer := identity.Issuer{
			DocType: identity.IssuerDocType,
			ID:      testing.ID1,
			Name:    "Tecnomatica",
			CertPem: b64RootCertTecnomatica,
		}
		bytes, err := json.Marshal(issuer)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		chaincodeStub.GetStateReturns(bytes, nil)
		response, err := sc.GetIssuer(ctx, model.GetRequest{ID: issuer.ID})
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(response.ID).To(gomega.BeIdenticalTo(testing.ID1))
	})

	ginkgo.It("get identity", func() {
		expected := model.ParticipantResponse{
			Did:     testing.Did1,
			Roles:   []string{},
			Creator: nil,
		}
		expectedJSON := testing.MarshalJSONOrPanic(expected)

		chaincodeStub.GetTxTimestampReturns(testing.Timestamp, nil)
		iterator := &mocks.StateQueryIterator{}
		iterator.HasNextReturnsOnCall(0, true)
		iterator.NextReturnsOnCall(0, &queryresult.KV{
			Key:   testing.ID1,
			Value: expectedJSON,
		}, nil)
		chaincodeStub.GetStateByPartialCompositeKeyReturns(iterator, nil)
		chaincodeStub.SplitCompositeKeyReturns("", []string{testing.Did1, testing.ID1}, nil)

		chaincodeStub.GetStateReturns(expectedJSON, nil)
		actual, err := sc.GetParticipant(ctx, model.ParticipantGetRequest{Did: testing.Did1})
		gomega.Expect(err).To(gomega.BeNil())
		actualJSON, err := json.Marshal(actual)
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(expectedJSON).Should(gomega.MatchJSON(actualJSON))
	})

	ginkgo.It("delete identity", func() {
		// using client with admin cert
		chaincodeStub.GetCreatorReturns(org1MSPAdmin, nil)

		returned := identity.Participant{
			Did:     testing.Did1,
			Roles:   []string{},
			Creator: "",
			Active:  true,
			MspID:   testing.MspID,
		}
		returnedJSON := testing.MarshalJSONOrPanic(returned)

		chaincodeStub.GetTxTimestampReturns(testing.Timestamp, nil)
		iterator := &mocks.StateQueryIterator{}
		iterator.HasNextReturnsOnCall(0, true)
		iterator.NextReturnsOnCall(0, &queryresult.KV{
			Key:   testing.ID1,
			Value: returnedJSON,
		}, nil)
		chaincodeStub.GetStateByPartialCompositeKeyReturns(iterator, nil)
		chaincodeStub.SplitCompositeKeyReturns("", []string{testing.Did1, testing.ID1}, nil)
		chaincodeStub.GetStateReturns(returnedJSON, nil)
		clientIdentity.GetMSPIDReturnsOnCall(0, testing.MspID, nil)

		requestDeleteParticipant := model.ParticipantDeleteRequest{
			UserDid:   testing.Did1,
			CallerDid: testing.Did1,
		}
		err := sc.DeleteParticipant(ctx, requestDeleteParticipant)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.It("create identity", func() {
		// using client with admin cert
		chaincodeStub.GetCreatorReturns(org1MSPAdmin, nil)
		// base64 encode
		b64ClientCert := base64.StdEncoding.EncodeToString(testing.CertByteUserYisel)
		gomega.Expect(err).To(gomega.BeNil())

		identRequest := model.ParticipantCreateRequest{
			Did:        testing.Did1,
			IssuerID:   testing.ID1,
			CreatorDid: "",
			CertPem:    b64ClientCert,
			Roles:      nil,
		}

		iterator := &mocks.StateQueryIterator{}
		iterator.HasNextReturnsOnCall(0, false)
		chaincodeStub.GetTxTimestampReturns(testing.Timestamp, nil)
		chaincodeStub.GetStateByPartialCompositeKeyReturns(iterator, nil)

		actual, err := sc.CreateParticipant(ctx, identRequest)
		gomega.Expect(err).To(gomega.BeNil())
		actualJSON, err := json.Marshal(actual)
		expected := model.ParticipantResponse{
			Did:     testing.Did1,
			Roles:   []string{},
			Creator: nil,
		}
		expectedJSON := testing.MarshalJSONOrPanic(expected)
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(expectedJSON).Should(gomega.MatchJSON(actualJSON))
	})

	ginkgo.It("get issuer histories", func() {
		chaincodeStub.GetTxTimestampReturns(testing.Timestamp, nil)
		ts := time.Unix(testing.Timestamp.Seconds, int64(testing.Timestamp.Nanos))
		formattedTime := ts.Format(time.RFC3339)

		//GetIssuerHistory
		record := &model.IssuerQueryResponse{
			ID:      testing.ID1,
			Name:    "Autoridad de Certificación Tecnomática",
			CertPem: testing.GomegaString(),
		}

		history1 := model.IssuerHistoryQueryResponse{
			Record:   record,
			TxID:     "TxId0",
			Time:     formattedTime,
			IsDelete: false,
		}
		history1JSON := testing.MarshalJSONOrPanic(history1)

		history2 := model.IssuerHistoryQueryResponse{
			Record:   record,
			TxID:     "TxId1",
			Time:     formattedTime,
			IsDelete: false,
		}
		history2JSON := testing.MarshalJSONOrPanic(history2)

		history3 := model.IssuerHistoryQueryResponse{
			Record:   record,
			TxID:     "TxId2",
			Time:     formattedTime,
			IsDelete: false,
		}
		history3JSON := testing.MarshalJSONOrPanic(history3)

		// Case: Get all histories
		iterator := &mocks.HistoryQueryIteratorInterface{}
		iterator.HasNextReturnsOnCall(0, true)
		iterator.HasNextReturnsOnCall(1, true)
		iterator.HasNextReturnsOnCall(2, true)
		iterator.HasNextReturnsOnCall(3, false)
		iterator.NextReturnsOnCall(0, &queryresult.KeyModification{TxId: "TxId0", Timestamp: testing.Timestamp, IsDelete: false, Value: history1JSON}, nil)
		iterator.NextReturnsOnCall(1, &queryresult.KeyModification{TxId: "TxId1", Timestamp: testing.Timestamp, IsDelete: false, Value: history2JSON}, nil)
		iterator.NextReturnsOnCall(2, &queryresult.KeyModification{TxId: "TxId2", Timestamp: testing.Timestamp, IsDelete: false, Value: history3JSON}, nil)

		chaincodeStub.GetHistoryForKeyReturns(iterator, nil)
		issuerRequest := model.GetRequest{
			ID: testing.ID1,
		}
		histories, _ := sc.GetIssuerHistory(ctx, issuerRequest)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect([]model.IssuerHistoryQueryResponse{history1, history2, history3}).ShouldNot(gomega.BeIdenticalTo(histories))
	})

})
