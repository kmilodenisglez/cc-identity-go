package testcerts

import (
	"io/ioutil"
	"path"
	"runtime"
)

type (
	FileReader func(filename string) ([]byte, error)
	// Cert certificate data for testing
	Cert struct {
		CertFilename string
		PKeyFilename string
		readFile     FileReader
	}

	Certs []*Cert
)

func (cc Certs) UseReadFile(readFile FileReader) Certs {
	for _, c := range cc {
		c.readFile = readFile
	}
	return cc
}

var (
	Certificates = Certs{
		{CertFilename: `rootcert.pem`, PKeyFilename: `rootcert.key.pem`},                           // 0
		{CertFilename: `opsadmin_issued_by_rootcert.pem`, PKeyFilename: `opsadmin.key.pem`},        // 1    admin cert, with OU=admin
		{CertFilename: `user1_with_attrs_issued_by_rootcert.pem`, PKeyFilename: `user1.key.pem`},   // 2    attributes: Cargo=Director de Calidad,Email=user1@matcom.uh.cu,Nombre=Pedro Perez,Edad=32, etc.
		{CertFilename: `client_with_attrs_issued_by_rootcert.pem`, PKeyFilename: `client.key.pem`}, // 3    attributes hf.EnrollmentID=user1, hf.Type=client
		{CertFilename: `cert-revoked-with-attrs.pem`, PKeyFilename: ``},                            // 4
		{CertFilename: `peer0_issued_by_rootcert.pem`, PKeyFilename: `peer0.key.pem`},              // 5
		{CertFilename: `user2_issued_by_unknown_rootcert.pem`, PKeyFilename: `user2.key.pem`},      // 6
		{CertFilename: `user3_issued_by_tecnomatica.pem`, PKeyFilename: ``},                        // 7
		{CertFilename: `rootcert_tecnomatica_ca.pem`, PKeyFilename: ``},                            // 8
		{CertFilename: `rootcert_dc_central.pem`, PKeyFilename: ``},                                // 9
	}.UseReadFile(ReadLocal())
)

func ReadLocal() func(filename string) ([]byte, error) {
	_, curFile, _, ok := runtime.Caller(1)
	dir := path.Dir(curFile)
	if !ok {
		return nil
	}
	return func(filename string) ([]byte, error) {
		return ioutil.ReadFile(dir + "/" + filename)
	}
}

func (c *Cert) CertBytes() ([]byte, error) {
	//return c.readFile(`./` + c.CertFilename)
	return c.readFile(c.CertFilename)
}

func (c *Cert) PrivateBytes() ([]byte, error) {
	return c.readFile(c.PKeyFilename)
}
