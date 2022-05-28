package domain

type RolePermissions struct {
	permissions map[string][]string
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}

func (r RolePermissions) IsAuthorizedFor(role string, resource string) bool {
	// Get all resource permissions for the role in the claim.
	// Resource string should be listed for authorization
	permissions := r.permissions[role]
	for _, p := range permissions {
		if p == resource {
			return true
		}
	}
	return false
}
