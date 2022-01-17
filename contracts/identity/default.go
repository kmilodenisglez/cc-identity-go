package identity

// GetIgnoredFunctions returns functions are still callable by the code just not directly by outside users
func (ci *ContractIdentity) GetIgnoredFunctions() []string {
	return []string{"CreateAccess"}
}
