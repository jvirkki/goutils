package debug

import "testing"

var data1 = []byte{0x45, 0x10, 0x00, 0x3c, 0xaa, 0x2e, 0x40, 0x00, 0x40, 0x06, 0x92,
	0x7b, 0x7f, 0x00, 0x00, 0x01, 0x7f, 0x00, 0x00, 0x01, 0xb2, 0x64, 0x00, 0x63, 0x58, 0xd1,
	0x24, 0xd9, 0x00, 0x00, 0x00, 0x00, 0xa0, 0x02, 0xaa, 0xaa, 0xfe, 0x30, 0x00, 0x00, 0x02,
	0x04, 0xff, 0xd7, 0x04, 0x02, 0x08, 0x0a, 0x00, 0xc5, 0x8e, 0xf7, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x03, 0x03, 0x07}

func TestMemdump(t *testing.T) {

	lines := memdump("packet info", data1, true, true, 16, false)
	t.Log(lines)
	if len(lines) != 5 {
		t.Error("Expected 5 lines")
	}
}

func TestMemdumpOneLine(t *testing.T) {

	lines := memdump("packet info", data1, true, true, 90, false)
	t.Log(lines)
	if len(lines) != 2 {
		t.Error("Expected 2 lines")
	}
	if len(lines[1]) != 189 {
		t.Errorf("Unexpected line length %d", len(lines[1]))
	}
}

func TestMemdumpOneLineNoHeader(t *testing.T) {

	lines := memdump("", data1, false, true, 90, false)
	t.Log(lines)
	if len(lines) != 1 {
		t.Error("Expected 1 lines")
	}
	if len(lines[0]) != 181 {
		t.Errorf("Unexpected line length %d", len(lines[0]))
	}
}

func TestMemdumpOneLineWithHeader(t *testing.T) {

	lines := memdump("bytes", data1, false, true, 0, false)
	t.Log(lines)
	if len(lines) != 1 {
		t.Error("Expected 1 lines")
	}
	if len(lines[0]) != 188 {
		t.Errorf("Unexpected line length %d", len(lines[0]))
	}
}

func TestMemdumpOneLineAnyway(t *testing.T) {

	lines := memdump("", data1, false, true, 90, true)
	t.Log(lines)
	if len(lines) != 1 {
		t.Error("Expected 1 lines")
	}
	if len(lines[0]) != 181 {
		t.Errorf("Unexpected line length %d", len(lines[0]))
	}
}

func TestMemdumpForceOneLine(t *testing.T) {

	lines := memdump("decimal counter", data1, true, false, 32, true)
	t.Log(lines)
	if len(lines) != 1 {
		t.Error("Expected 1 lines")
	}
	if len(lines[0]) != 212 {
		t.Errorf("Unexpected line length %d", len(lines[0]))
	}
}

func TestPrintMemDump(t *testing.T) {
	PrintMemDump("print bytes", data1)
}

func TestConfigDump(t *testing.T) {

	var c MemDumpConfig

	c.BytesPerLine = 16
	c.ShowCounts = true
	c.CountsInHex = false

	lines := c.MemDump("by config", data1)
	t.Log(lines)
	if len(lines) != 5 {
		t.Error("Expected 5 lines")
	}
}

func TestConfigPrint(t *testing.T) {

	var c MemDumpConfig

	c.BytesPerLine = 8
	c.ShowCounts = true
	c.CountsInHex = true

	c.PrintMemDump("print by config", data1)
}
