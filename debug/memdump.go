// Package debug contains assorted small helper functions for debugging.
package debug

import (
	"bytes"
	"fmt"
)

//
// Contains configuration for MemDump output.
//
// BytesPerLine - show up to this many bytes in each line. If zero, no limit.
//
// ShowCounts - if true, prefix each new line with a position counter.
//
// CountsInHex - if true, the position counter is in hex (otherwise decimal)
//
type MemDumpConfig struct {
	BytesPerLine int
	ShowCounts   bool
	CountsInHex  bool
}

//
// Print a single-line hex dump of the bytes in ptr, prefixed by an optional heading.
//
func PrintMemDump(heading string, ptr []byte) {

	fmt.Println(MemDump(heading, ptr))
}

//
// Return a single-line hex dump of the bytes in ptr, prefixed by an optional heading.
//
func MemDump(heading string, ptr []byte) string {

	lines := memdump(heading, ptr, false, false, 0, false)
	return lines[0]
}

//
// Print a hex dump of the bytes in ptr, prefixed by an optional heading.
// The formatting of the output is controlled by the MemDumpConfig settings.
//
func (c *MemDumpConfig) PrintMemDump(heading string, ptr []byte) {

	lines := memdump(heading, ptr, c.ShowCounts, c.CountsInHex, c.BytesPerLine, true)
	fmt.Println(lines[0])
}

//
// Return a hex dump of the bytes in ptr, prefixed by an optional heading.
// The formatting of the output is controlled by the MemDumpConfig settings.
//
func (c *MemDumpConfig) MemDump(heading string, ptr []byte) []string {
	return memdump(heading, ptr, c.ShowCounts, c.CountsInHex, c.BytesPerLine, false)
}

//
// memdump does the actual work for all the related convenience functions.
//
// heading is an optional heading printed as the first line (in multi-line mode)
// or as a line prefix (in single line mode).
//
// ptr is the memory buffer to process.
//
// If counts is true, a position counter is added as the first column of each new line.
//
// If countsInHex is true, the position counter is in hex, otherwise decimal.
//
// bytesPerLine limits how many bytes shown per line. If zero, there is no limit.
//
// forceOneLiner fits all the output in one line, possibly containing newlines.
//
func memdump(heading string, ptr []byte,
	counts bool, countsInHex bool, bytesPerLine int, forceOneLiner bool) []string {

	var buf bytes.Buffer
	var lines []string

	if heading != "" {
		buf.WriteString(heading)
		buf.WriteString(": ")

		if bytesPerLine > 0 {
			buf.WriteString("\n")
			lines = append(lines, buf.String())
			buf.Reset()
		}
	}

	if bytesPerLine > 0 && counts {
		if countsInHex {
			buf.WriteString("0x0000  ")
		} else {
			buf.WriteString("0000  ")
		}
	}

	for i := 0; i < len(ptr); i++ {
		buf.WriteString(fmt.Sprintf("%02x ", ptr[i]))
		if bytesPerLine > 0 && (i+1)%bytesPerLine == 0 {
			buf.WriteString("\n")
			lines = append(lines, buf.String())
			buf.Reset()
			if counts {
				if countsInHex {
					buf.WriteString(fmt.Sprintf("0x%04x  ", i+1))
				} else {
					buf.WriteString(fmt.Sprintf("%04d  ", i+1))
				}
			}
		}
	}

	if buf.Len() > 0 {
		buf.WriteString("\n")
		lines = append(lines, buf.String())
	}

	if !forceOneLiner {
		return lines
	}

	if len(lines) == 1 {
		return lines
	}

	var out []string
	var line bytes.Buffer
	for _, l := range lines {
		line.WriteString(l)
	}
	return append(out, line.String())
}
