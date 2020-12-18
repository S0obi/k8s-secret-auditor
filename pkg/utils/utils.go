package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// PrintResultTable : Render result table
func PrintResultTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Secret name", "Value", "Info"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
