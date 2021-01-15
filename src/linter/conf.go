package linter

import (
	"regexp"
	"runtime"
	"time"

	"github.com/client9/misspell"

	"github.com/VKCOM/noverify/src/baseline"
	"github.com/VKCOM/noverify/src/inputs"
	"github.com/VKCOM/noverify/src/rules"
)

var (
	// BaselineProfile is a suppression database for warnings.
	// Nil profile is an empty suppression profile.
	BaselineProfile      *baseline.Profile
	ConservativeBaseline bool

	ApplyQuickFixes bool

	// KPHP tells whether we're working in KPHP-compatible mode.
	KPHP bool

	CacheDir string

	// TypoFixer is a rule set for English typos correction.
	// If nil, no misspell checking is performed.
	// See github.com/client9/misspell for details.
	TypoFixer *misspell.Replacer

	// SrcInput implements source code reading from files and buffers.
	//
	// TODO(quasilyte): avoid having it as a global variable?
	SrcInput = inputs.NewDefaultSourceInput()

	// Rules is a set of dynamically loaded linter diagnostics.
	Rules = &rules.Set{}

	// settings

	StubsDir       string
	Debug          bool
	MaxConcurrency = runtime.NumCPU()
	MaxFileSize    int

	// DebugParseDuration specifies the minimum parse duration for it to be printed to debug output.
	DebugParseDuration time.Duration

	CheckAutoGenerated bool

	IsDiscardVar = isUnderscore

	ExcludeRegex *regexp.Regexp

	// actually time.Duration
	initParseTime int64
	initWalkTime  int64
)
