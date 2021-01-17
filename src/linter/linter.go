package linter

import (
	"sync"
	"time"

	"github.com/VKCOM/noverify/src/lintdebug"
	"github.com/VKCOM/noverify/src/meta"
	"github.com/VKCOM/noverify/src/workspace"
)

type Linter struct {
	config *Config
}

func NewLinter(config *Config) *Linter {
	return &Linter{
		config: config,
	}
}

func (l *Linter) NewLintingWorker(id int) *Worker {
	w := newWorker(l.config, id)
	w.needReports = true
	return w
}

func (l *Linter) NewIndexingWorker(id int) *Worker {
	w := newWorker(l.config, id)
	w.needReports = false
	return w
}

// AnalyzeFiles runs linter on the files that are provided by the readFileNamesFunc function.
func (l *Linter) AnalyzeFiles(readFileNamesFunc workspace.ReadCallback) []*Report {
	return l.analyzeFiles(readFileNamesFunc)
}

func (l *Linter) analyzeFiles(readFileNamesFunc workspace.ReadCallback) []*Report {
	start := time.Now()
	defer func() {
		lintdebug.Send("Processing time: %s", time.Since(start))

		meta.Info.Lock()
		defer meta.Info.Unlock()

		lintdebug.Send("Funcs: %d, consts: %d, files: %d", meta.Info.NumFunctions(), meta.Info.NumConstants(), meta.Info.NumFilesWithFunctions())
	}()

	needReports := meta.IsIndexingComplete()

	lintdebug.Send("Parsing using %d cores", l.config.MaxConcurrency)

	filenamesCh := make(chan workspace.FileInfo, 512)

	go func() {
		readFileNamesFunc(filenamesCh)
		close(filenamesCh)
	}()

	var wg sync.WaitGroup
	reportsCh := make(chan []*Report, l.config.MaxConcurrency)

	wg.Add(l.config.MaxConcurrency)
	for i := 0; i < l.config.MaxConcurrency; i++ {
		go func(id int) {
			var w *Worker
			if needReports {
				w = l.NewLintingWorker(id)
			} else {
				w = l.NewIndexingWorker(id)
			}
			w.AllowDisable = l.config.AllowDisable
			var rep []*Report
			for f := range filenamesCh {
				rep = append(rep, w.doParseFile(f)...)
			}
			reportsCh <- rep
			wg.Done()
		}(i)
	}
	wg.Wait()

	var allReports []*Report
	for i := 0; i < l.config.MaxConcurrency; i++ {
		allReports = append(allReports, <-reportsCh...)
	}

	return allReports
}

func (l *Linter) InitStubs(readFileNamesFunc workspace.ReadCallback) {
	meta.SetLoadingStubs(true)
	l.AnalyzeFiles(readFileNamesFunc)
	meta.Info.InitStubs()
	if l.config.KPHP {
		meta.Info.InitKphpStubs()
	}
	meta.SetLoadingStubs(false)
}

// InitStubsFromDir parses directory with PHPStorm stubs which has all internal PHP classes and functions declared.
func (l *Linter) InitStubsFromDir(dir string) {
	l.InitStubs(workspace.ReadFilenames([]string{dir}, nil, l.config.PhpExtensions))
}
