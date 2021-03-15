package types

const (
	// ModuleName is the name of the module
	ModuleName = "woodpecker"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	AttributePrefix = "attribute-value-"
	AttributeCountPrefix = "attribute-count-"
)
		
const (
	MedicalHistoryPrefix = "medicalHistory-value-"
	MedicalHistoryCountPrefix = "medicalHistory-count-"
)
		
const (
	BodyIndexPrefix = "bodyIndex-value-"
	BodyIndexCountPrefix = "bodyIndex-count-"
)
		