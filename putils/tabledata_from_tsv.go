package putils

import (
	"github.com/gemini/pterm"
)

// TableDataFromTSV converts TSV data into pterm.TableData.
//
// Usage:
//	pterm.DefaultTable.WithData(putils.TableDataFromTSV(tsv)).Render()
func TableDataFromTSV(csv string) (td pterm.TableData) {
	return TableDataFromSeparatedValues(csv, "\t", "\n")
}
