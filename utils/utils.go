package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// type TableHeader []string

func Table(header []string, data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetHeaderLine(true)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	//table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
}
