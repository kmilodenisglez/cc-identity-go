package identity

// GetIgnoredFunctions returns functions are still callable by the code just not directly by outside users
func (ic *ContractIdentity) GetIgnoredFunctions() []string {
	return []string{"CreateAccess"}
}
