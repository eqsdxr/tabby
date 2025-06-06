package tabby

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Tabby is returned when New() is called.
type Tabby struct {
	writer *tabwriter.Writer
}

// New returns a new *tabwriter.Writer with default config
func New() *Tabby {
	return &Tabby{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

// NewCustom returns a new *Tabby with custom *tabwriter.Writer set
func NewCustom(writer *tabwriter.Writer) *Tabby {
	return &Tabby{writer: writer}
}

// AddLine will write a new table line
func (t *Tabby) AddLine(args ...interface{}) {
	formatString := t.buildFormatString(args)
	fmt.Fprintf(t.writer, formatString, args...)
}

// AddHeader will write a new table line followed by a separator
func (t *Tabby) AddHeader(args ...interface{}) {
	t.AddLine(args...)
	t.addSeparator(args)
}

// Print will write the table to the terminal
func (t *Tabby) Print() {
	t.writer.Flush()
}

// addSeparator will write a new dash seperator line based on the args length
func (t *Tabby) addSeparator(args []interface{}) {
	var b bytes.Buffer
	for idx, arg := range args {
		length := len(fmt.Sprintf("%v", arg))
		b.WriteString(strings.Repeat("-", length))
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	b.WriteTo(t.writer)
}

// buildFormatString will build up the formatting string used by the *tabwriter.Writer
func (t *Tabby) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	return b.String()
}
