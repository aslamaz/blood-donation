package constant

const (
	JwtSigningKey = "my_secret_key"
)

var (
	BloodGroups = []string{"A+", "A-", "B+", "B-", "O+", "O-", "AB+", "AB-"}
)

func IsValidBloodGroup(bloodGroup string) bool {
	for _, bg := range BloodGroups {
		if bg == bloodGroup {
			return true
		}
	}
	return false
}
