package workspace

import (
	"bytes"
	"strings"

	"github.com/VKCOM/noverify/src/git"
)

type File struct {
	name     string
	contents []byte

	lineRanges []git.LineRange
	lines      [][]byte
	linesPos   []int

	autoGenerated bool
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func NewFileWithContents(name string, contents []byte) *File {
	lines, linesPos := parseContents(contents)

	return &File{
		name:          name,
		contents:      contents,
		lines:         lines,
		linesPos:      linesPos,
		autoGenerated: FileIsAutoGenerated(contents),
	}
}

func parseContents(contents []byte) (lines [][]byte, linesPos []int) {
	lines = bytes.Split(contents, []byte("\n"))
	linesPos = make([]int, len(lines))
	pos := 0
	for idx, ln := range lines {
		linesPos[idx] = pos
		pos += len(ln) + 1
	}
	return lines, linesPos
}

func (f *File) Name() string {
	return f.name
}

func (f *File) CountLines() int {
	return len(f.lines)
}

func (f *File) Line(index int) []byte {
	if index < 0 || index >= len(f.lines) {
		return nil
	}

	return f.lines[index]
}

func (f *File) CountLinePosition() int {
	return len(f.linesPos)
}

func (f *File) LinePosition(index int) int {
	if index < 0 || index >= len(f.linesPos) {
		return -1
	}
	return f.linesPos[index]
}

func (f *File) LineRanges() []git.LineRange {
	return f.lineRanges
}

func (f *File) SetLineRanges(r []git.LineRange) {
	f.lineRanges = r
}

func (f *File) ResetContents(newContents []byte) {
	f.contents = newContents
	f.lines, f.linesPos = parseContents(f.contents)
}

func (f *File) Contents() []byte {
	return f.contents
}

func (f *File) Part(start, end int) []byte {
	if start > end {
		return nil
	}
	if start < 0 || end < 0 {
		return nil
	}
	if start >= len(f.contents) || end >= len(f.contents) {
		return nil
	}

	return f.contents[start:end]
}

func (f *File) AutoGenerated() bool {
	return f.autoGenerated
}

func FileIsAutoGenerated(contents []byte) bool {
	src := contents
	maxLinesPeek := 15

	doNotEdit := false
	autoGenerated := false

	for line := 0; line < maxLinesPeek; line++ {
		if doNotEdit && autoGenerated {
			return true
		}
		newline := bytes.IndexByte(src, '\n')
		if newline == -1 {
			break
		}

		l := strings.ToLower(string(src[:newline]))
		src = src[newline+len("\n"):]

		looksLikeComment := strings.HasPrefix(l, "//") ||
			strings.HasPrefix(l, "/*") ||
			strings.HasPrefix(l, " *")
		if !looksLikeComment {
			continue
		}

		doNotEdit = doNotEdit ||
			strings.Contains(l, "do not edit") ||
			strings.Contains(l, "don't edit") ||
			strings.Contains(l, "do not modify") ||
			strings.Contains(l, "don't modify")
		autoGenerated = autoGenerated ||
			strings.Contains(l, "auto-generated") ||
			strings.Contains(l, "autogenerated") ||
			strings.Contains(l, "generated by")
	}

	return doNotEdit && autoGenerated
}
