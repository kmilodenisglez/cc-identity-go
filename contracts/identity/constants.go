package identity

// docType
const (
	ParticipantDocType = "did.participant"
	RoleDocType        = "did.role"
	AccessDocType      = "did.access"
	IssuerDocType      = "did.issuer"
)

const (
	// index
	Deleted = "deleted"

	// objectType
	ObjectTypeParticipantByDidUuid = ParticipantDocType + ":did~uuid"
	ObjectTypeParticipantDeleted   = ParticipantDocType + ":" + Deleted + "~did~uuid" // use to index deleted participant
	ObjectTypeIssuerByDefault      = IssuerDocType + ":default~uuid"
)
