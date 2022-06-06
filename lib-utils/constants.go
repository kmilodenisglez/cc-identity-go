package libutils

// Error responses
// errorXXX occurs when XXX
const (
	ErrorParseJWS          = `error parsing into JWS`
	ErrorParseX509         = `error parsing into X509`
	ErrorBase64            = `error decoding into base64`
	ErrorVerifying         = `error verifying signature`
	ErrorGetMSPID          = `failed getting the client's MSPID: %v`
	ErrorGetIdentity       = `failed to get identity %s`
	ErrorIdentityExists    = `identity %s already exists`
	ErrorDefaultNotExist   = `%s does not exist`
	ErrorRequiredParameter = "a required parameter (%s) was not provided"
)
