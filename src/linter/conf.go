package linter

import (
	"regexp"
	"runtime"
	"time"

	"github.com/VKCOM/noverify/src/baseline"
	"github.com/VKCOM/noverify/src/inputs"
	"github.com/VKCOM/noverify/src/rules"
	"github.com/client9/misspell"
)

var (
	// LangServer represents whether or not we run in a language server mode.
	LangServer bool

	// BaselineProfile is a suppression database for warnings.
	// Nil profile is an empty suppression profile.
	BaselineProfile      *baseline.Profile
	ConservativeBaseline bool

	CacheDir string

	// TypoFixer is a rule set for English typos correction.
	// If nil, no misspell checking is performed.
	// See github.com/client9/misspell for details.
	TypoFixer *misspell.Replacer

	// AnalysisFiles is a list of files that are being analyzed (in non-git mode)
	AnalysisFiles []string

	// SrcInput implements source code reading from files and buffers.
	//
	// TODO(quasilyte): avoid having it as a global variable?
	SrcInput = inputs.NewDefaultSourceInput()

	// Rules is a set of dynamically loaded linter diagnostics.
	Rules = &rules.Set{}

	// settings
	StubsDir        string
	Debug           bool
	MaxConcurrency  = runtime.NumCPU()
	MaxFileSize     int
	DefaultEncoding string
	PHPExtensions   []string

	// DebugParseDuration specifies the minimum parse duration for it to be printed to debug output.
	DebugParseDuration time.Duration

	CheckAutoGenerated bool

	IsDiscardVar = isUnderscore

	ExcludeRegex *regexp.Regexp

	// actually time.Duration
	initParseTime int64
	initWalkTime  int64
)
