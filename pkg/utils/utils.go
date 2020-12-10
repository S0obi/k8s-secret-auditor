package utils

import (
	"math"
	"os"
	"strings"
	"unicode"

	"github.com/S0obi/k8s-secret-auditor/pkg/config"
	"github.com/olekukonko/tablewriter"
)

// IsPassword : determine if the value is a password
func IsPassword(value string, config *config.Config) bool {
	for _, password := range config.PasswordPatterns {
		if strings.Contains(strings.ToLower(value), password) {
			return true
		}
	}
	return false
}

// ComputeEntropy : https://stackoverflow.com/questions/6151576/how-to-check-password-strength
func ComputeEntropy(value string) float64 {
	//cardinality := 0
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

// PrintResultTable : Render result table
func PrintResultTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Secret name", "Value", "Info"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
