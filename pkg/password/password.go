package password

import (
	"math"
	"strings"
	"unicode"

	"github.com/S0obi/k8s-secret-auditor/pkg/config"
)

// Password : Password information
type Password struct {
	Name      string
	Value     string
	Compliant bool
	Entropy   float64
}

// NewPassword : Constructor of Password struct
func NewPassword(key string, value string) *Password {
	return &Password{Name: key, Value: value, Entropy: computeEntropy(value)}
}

// IsCompliant : Check if the password is compliant regarding a config
func (pwdInfo *Password) IsCompliant(config *config.Config) bool {
	if len(pwdInfo.Value) < config.Policy.Length || pwdInfo.Entropy < config.Policy.Entropy {
		return false
	}
	return true
}

// IsPassword : determine if the value is a password
func IsPassword(value string, config *config.Config) bool {
	for _, password := range config.PasswordPatterns {
		if strings.Contains(strings.ToLower(value), password) {
			return true
		}
	}
	return false
}

// computeEntropy : https://stackoverflow.com/questions/6151576/how-to-check-password-strength
func computeEntropy(value string) float64 {
	characteristics := map[string]float64{
		"lower":   0.,
		"upper":   0.,
		"digit":   0.,
		"symbols": 0.,
	}

	for _, character := range value {
		if unicode.IsDigit(character) {
			characteristics["digit"] = 10.
		} else if unicode.IsLower(character) {
			characteristics["lower"] = 26.
		} else if unicode.IsUpper(character) {
			characteristics["upper"] = 26.
		} else {
			characteristics["symbols"] = 36.
		}
	}

	result := characteristics["digit"] + characteristics["lower"] + characteristics["upper"] + characteristics["symbols"]
	return math.Log2(result) * float64(len(value))
}
