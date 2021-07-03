package cmd

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const AllNonNoticeChecks = "<all-non-notice>"
const AllChecks = "<all>"

type ParsedFlags struct {
	PprofHost string

	CPUProfile string
	MemProfile string

	MaxFileSize int

	FullAnalysisFiles  string
	IndexOnlyFiles     string
	CheckAutoGenerated bool

	RulesList string

	Output         string
	OutputJSON     bool
	OutputBaseline bool

	Debug              bool
	DebugParseDuration time.Duration

	MaxConcurrency int

	StubsDir     string
	CacheDir     string
	DisableCache bool

	IgnoreTriggerError bool

	ApplyQuickFixes bool

	KPHP bool

	Baseline             string
	ConservativeBaseline bool

	MisspellList string

	UnusedVarPattern string

	AllowAll     bool
	AllowChecks  string
	AllowDisable string

	ReportsExclude       string
	ReportsExcludeChecks string
	ReportsCritical      string

	PhpExtensionsArg string

	Gitignore bool

	GitPushArg                 string
	GitAuthorsWhitelist        string
	GitWorkTree                string
	GitSkipFetch               bool
	GitDisableCompensateMaster bool
	GitFullDiff                bool
	GitIncludeUntracked        bool
	GitRepo                    string

	// These two flags are mutated in prepareGitArgs.
	// This is bad, but it's easier for now than to fix this
	// without introducing other issues.
	Mutable struct {
		GitCommitFrom string
		GitCommitTo   string
	}
}

func DefaultCacheDir() string {
	defaultCacheDir, err := os.UserCacheDir()
	if err != nil {
		defaultCacheDir = ""
	} else {
		defaultCacheDir = filepath.Join(defaultCacheDir, "noverify-cache")
	}
	return defaultCacheDir
}

func RegisterCheckFlags(ctx *AppContext) *flag.FlagSet {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)

	fs.StringVar(&ctx.ParsedFlags.PprofHost, "pprof", "", "HTTP pprof endpoint (e.g. localhost:8080)")

	fs.StringVar(&ctx.ParsedFlags.Baseline, "baseline", "",
		"Path to a suppress profile created by -output-baseline")
	fs.BoolVar(&ctx.ParsedFlags.ConservativeBaseline, "conservative-baseline", false,
		"If enabled, baseline mode will have less false positive, but more false negatives")

	fs.StringVar(&ctx.ParsedFlags.ReportsCritical, "critical", AllNonNoticeChecks,
		"Comma-separated list of check names that are considered critical (all non-notice checks by default)")

	fs.StringVar(&ctx.ParsedFlags.RulesList, "rules", "",
		"Comma-separated list of rules files")

	fs.BoolVar(&ctx.ParsedFlags.ApplyQuickFixes, "fix", false,
		"Apply a quickfix where possible (updates source files)")

	fs.BoolVar(&ctx.ParsedFlags.Gitignore, "gitignore", false,
		"If enabled, noverify tries to use .gitignore files to exclude matched ignored files from the analysis")
	fs.BoolVar(&ctx.ParsedFlags.KPHP, "kphp", false,
		"If enabled, treat the code as KPHP")

	fs.StringVar(&ctx.ParsedFlags.GitRepo, "git", "", "Path to git repository to analyze")
	fs.StringVar(&ctx.ParsedFlags.Mutable.GitCommitFrom, "git-commit-from", "", "Analyze changes between commits <git-commit-from> and <git-commit-to>")
	fs.StringVar(&ctx.ParsedFlags.Mutable.GitCommitTo, "git-commit-to", "", "Analyze changes between commits <git-commit-from> and <git-commit-to>")
	fs.StringVar(&ctx.ParsedFlags.GitPushArg, "git-push-arg", "", "In {pre,post}-receive hooks a whole line from stdin can be passed")
	fs.StringVar(&ctx.ParsedFlags.GitAuthorsWhitelist, "git-author-whitelist", "", "Whitelist (comma-separated) for commit authors, if needed")
	fs.StringVar(&ctx.ParsedFlags.GitWorkTree, "git-work-tree", "", "Work tree. If specified, local changes will also be examined")
	fs.BoolVar(&ctx.ParsedFlags.GitSkipFetch, "git-skip-fetch", false, "Do not fetch ORIGIN_MASTER (use this option if you already fetch to ORIGIN_MASTER before that)")
	fs.BoolVar(&ctx.ParsedFlags.GitDisableCompensateMaster, "git-disable-compensate-master", false, "Do not try to compensate for changes in ORIGIN_MASTER after branch point")
	fs.BoolVar(&ctx.ParsedFlags.GitFullDiff, "git-full-diff", false, "Compute full diff: analyze all files, not just changed ones")
	fs.BoolVar(&ctx.ParsedFlags.GitIncludeUntracked, "git-include-untracked", true, "Include untracked (new, uncommitted files) into analysis")

	fs.StringVar(&ctx.ParsedFlags.ReportsExclude, "exclude", "", "Exclude regexp for filenames in reports list")
	fs.StringVar(&ctx.ParsedFlags.ReportsExcludeChecks, "exclude-checks", "", "Comma-separated list of check names to be excluded")
	fs.StringVar(&ctx.ParsedFlags.AllowDisable, "allow-disable", "", "Regexp for filenames where '@linter disable' is allowed")
	fs.StringVar(&ctx.ParsedFlags.AllowChecks, "allow-checks", AllChecks,
		"Comma-separated list of check names to be enabled")
	fs.BoolVar(&ctx.ParsedFlags.AllowAll, "allow-all-checks", false,
		"Enables all checks. Has the same effect as passing '<all>' to the -allow-checks parameter")
	fs.StringVar(&ctx.ParsedFlags.MisspellList, "misspell-list", "Eng",
		"Comma-separated list of misspelling dicts; predefined sets are Eng, Eng/US and Eng/UK")

	fs.StringVar(&ctx.ParsedFlags.PhpExtensionsArg, "php-extensions", "php,inc,php5,phtml", "List of PHP extensions to be recognized")

	fs.StringVar(&ctx.ParsedFlags.FullAnalysisFiles, "full-analysis-files", "", "Comma-separated list of files to do full analysis")
	fs.StringVar(&ctx.ParsedFlags.IndexOnlyFiles, "index-only-files", "", "Comma-separated list of files to do indexing")

	fs.StringVar(&ctx.ParsedFlags.Output, "output", "", "Output reports to a specified file instead of stderr")
	fs.BoolVar(&ctx.ParsedFlags.OutputJSON, "output-json", false, "Format output as JSON")
	fs.BoolVar(&ctx.ParsedFlags.OutputBaseline, "output-baseline", false, "Output a suppression profile instead of reports")

	fs.BoolVar(&ctx.ParsedFlags.CheckAutoGenerated, `check-auto-generated`, false, "Whether to lint auto-generated PHP file")
	fs.BoolVar(&ctx.ParsedFlags.Debug, "debug", false, "Enable debug output")
	fs.DurationVar(&ctx.ParsedFlags.DebugParseDuration, "debug-parse-duration", 0, "Print files that took longer than the specified time to analyse")
	fs.IntVar(&ctx.ParsedFlags.MaxFileSize, "max-sum-filesize", 20*1024*1024, "Max total file size to be parsed concurrently in bytes (limits max memory consumption)")
	fs.IntVar(&ctx.ParsedFlags.MaxConcurrency, "cores", runtime.NumCPU(), "Max cores")

	fs.StringVar(&ctx.ParsedFlags.StubsDir, "stubs-dir", "", "Directory with phpstorm-stubs")
	fs.StringVar(&ctx.ParsedFlags.CacheDir, "cache-dir", DefaultCacheDir(), "Directory for linter cache (greatly improves indexing speed)")
	fs.BoolVar(&ctx.ParsedFlags.DisableCache, "disable-cache", false, "If set, cache is not used and cache-dir is ignored")
	fs.BoolVar(&ctx.ParsedFlags.IgnoreTriggerError, "ignore-trigger-error", false, "If set, trigger_error control flow will be ignored")

	fs.StringVar(&ctx.ParsedFlags.UnusedVarPattern, "unused-var-regex", `^_$`,
		"Variables that match such regexp are marked as discarded; not reported as unused, but should not be used as values")

	fs.StringVar(&ctx.ParsedFlags.CPUProfile, "cpuprofile", "", "Write cpu profile to `file`")
	fs.StringVar(&ctx.ParsedFlags.MemProfile, "memprofile", "", "Write memory profile to `file`")

	var encodingUnused string
	fs.StringVar(&encodingUnused, "encoding", "", "Deprecated and unused")

	return fs
}
