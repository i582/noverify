package langsrv

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	dbg "runtime/debug"
	"strconv"
	"strings"
	"sync"

	"go.lsp.dev/uri"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/lintdebug"
	"github.com/VKCOM/noverify/src/linter"
	"github.com/VKCOM/noverify/src/vscode"
	"github.com/VKCOM/noverify/src/workspace"
)

const maxLength = 16 << 20

var respMutex sync.Mutex

type baseRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      *int   `json:"id"`
	Method  string `json:"method"`
	Params  json.RawMessage
}

type response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      *int        `json:"id"`
	Result  interface{} `json:"result"`
}

type methodCall struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func (response) IMessage()   {}
func (methodCall) IMessage() {}

func writeLog(msg string) {
	_ = writeMessage(&methodCall{
		JSONRPC: "2.0",
		Method:  "window/logMessage",
		Params: map[string]interface{}{
			"type":    3,
			"message": msg,
		},
	})
}

type LangServer struct {
	lint *linter.Linter

	openMapMutex sync.Mutex
	openMap      map[string]openedFile

	changingMutex sync.Mutex

	indexer  *linter.Worker
	analyzer *linter.Worker
}

func NewLangServer(lint *linter.Linter) *LangServer {
	return &LangServer{
		lint:     lint,
		openMap:  map[string]openedFile{},
		indexer:  lint.NewIndexingWorker(0),
		analyzer: lint.NewLintingWorker(0),
	}
}

// Start starts Microsoft LSP server with stdin/stdout I/O.
func (ls *LangServer) Start() error {
	lintdebug.Register(writeLog)

	rd := bufio.NewReader(os.Stdin)
	// connWr := os.Stdout

	for {
		ln, err := rd.ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read: %s", err.Error())
		}

		if !strings.HasPrefix(ln, "Content-Length: ") {
			return fmt.Errorf("wrong line: expected 'Content-Length:', got '%s'", ln)
		}

		length, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(ln, "Content-Length: ")))
		if err != nil {
			return fmt.Errorf("could not parse content length: %s", err.Error())
		}

		// should be empty line
		if _, err := rd.ReadString('\n'); err != nil {
			return fmt.Errorf("could not read delimiter: %v", err)
		}

		if length > maxLength {
			return fmt.Errorf("length too high: %d, max: %d", length, maxLength)
		}

		buf := make([]byte, length)
		if _, err = io.ReadFull(rd, buf); err != nil {
			return fmt.Errorf("could not read message: %v", err)
		}

		// lintdebug.Send(string(buf))
		if err := ls.handleMessage(buf); err != nil {
			return fmt.Errorf("could not write message: %v", err)
		}
	}
}

func (ls *LangServer) handleMessage(buf []byte) error {
	defer func() {
		if r := recover(); r != nil {
			lintdebug.Send("Panic occurred: %s, trace: %s", r, dbg.Stack())
		}
	}()

	var req baseRequest
	err := json.Unmarshal(buf, &req)
	if err != nil {
		return err
	}

	// lintdebug.Send("%+v", string(buf))

	switch req.Method {
	case "initialize":
		return ls.handleInitialize(&req)
	// case "textDocument/didOpen":
	// 	return handleTextDocumentDidOpen(&req)
	case "textDocument/didChange":
		return ls.handleTextDocumentDidChange(&req)
	// case "textDocument/definition":
	// 	return handleTextDocumentDefinition(&req)
	// case "textDocument/references":
	// 	return handleTextDocumentReferences(&req)
	// case "textDocument/didClose":
	// 	return handleTextDocumentDidClose(&req)
	// case "textDocument/completion":
	// 	return handleTextDocumentCompletion(&req)
	// case "textDocument/hover":
	// 	return handleTextDocumentHover(&req)
	// case "textDocument/documentSymbol":
	// 	return handleTextDocumentSymbol(&req)
	case "workspace/didChangeWatchedFiles":
		return ls.handleTextDocumentDidChange(&req)
	default:
		lintdebug.Send("Got %s, data: %s", req.Method, req.Params)
	}

	if req.ID == nil {
		return nil
	}

	return writeMessage(&response{
		JSONRPC: req.JSONRPC,
		ID:      req.ID,
		Result:  map[string]interface{}{},
	})
}

func writeMessage(message interface{ IMessage() }) error {
	respMutex.Lock()
	defer respMutex.Unlock()

	buf := bytes.NewBuffer(nil)

	_, err := buf.Write([]byte("Content-Type: application/vscode-jsonrpc; charset=utf8\r\n"))
	if err != nil {
		return err
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(buf, "Content-Length: %d\r\n\r\n", len(data))
	if err != nil {
		return err
	}

	_, err = buf.Write(data)
	if err != nil {
		return err
	}

	// lintdebug.Send(buf.String())

	_, err = os.Stdout.Write(buf.Bytes())

	return err
}

func (ls *LangServer) handleInitialize(req *baseRequest) error {
	var params vscode.InitializeParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		return err
	}

	// lintdebug.Send("Params: %+v", params)

	if params.RootURI != "" {
		lintdebug.Send("Root dir: %s", params.RootURI.Filename())
	}

	// go func() {
	// 	analysisFiles := []string{params.RootURI.Filename()}
	//
	// 	filter := workspace.NewFilenameFilter(ls.lint.Config().ExcludeRegex)
	// 	ls.lint.AnalyzeFiles(workspace.ReadFilenames(analysisFiles, filter, ls.lint.Config().PhpExtensions))
	// 	ls.lint.MetaInfo().SetIndexingComplete(true)
	//
	// 	// fully analyze all opened files
	// 	// other files are not analyzed fully at all
	// 	ls.openMapMutex.Lock()
	// 	for filename, op := range ls.openMap {
	// 		lintdebug.Send("open: %s", filename)
	// 		go ls.openFile(filename, string(op.file.Contents()))
	// 	}
	// 	ls.openMapMutex.Unlock()
	// }()

	lintdebug.Send(req.JSONRPC)

	writeMessage(&response{
		JSONRPC: req.JSONRPC,
		ID:      req.ID,
		Result: map[string]interface{}{
			"capabilities": map[string]interface{}{
				// "codeActionProvider":               nil,
				// "codeLensProvider":                 nil,
				"textDocumentSync": 1, // FULL
				"textDocument": map[string]interface{}{
					"publishDiagnostics": true,
				},
				// "documentSymbolProvider":           true,
				// "workspaceSymbolProvider":          true,
				// "definitionProvider":               true,
				// "dependenciesProvider":             nil,
				// "documentFormattingProvider":       nil,
				// "documentHighlightProvider":        nil,
				// "documentOnTypeFormattingProvider": nil,
				// "documentRangeFormattingProvider":  nil,
				// "referencesProvider":               true,
				// "hoverProvider":                    true,
				// "completionProvider": map[string]interface{}{
				// 	"resolveProvider":   true,
				// 	"triggerCharacters": []string{"$", ">", "\\"},
				// },
				// "renameProvider": nil,
				// "signatureHelpProvider": map[string]interface{}{
				// 	"triggerCharacters": []string{"(", ","},
				// },
				// "xworkspaceReferencesProvider": true,
				// "xdefinitionProvider":          true,
				// "xdependenciesProvider":        true,
			},
		},
	})

	ls.flushReports(`c:/Users/p.makhnev/Music/2.php`, nil)

	return nil
}

type openedFile struct {
	rootNode ir.Node
	// scopes   map[ir.Node]*meta.Scope
	file *workspace.File
}

func (ls *LangServer) openFile(filename, contents string) {
	ls.changingMutex.Lock()
	defer ls.changingMutex.Unlock()

	if ls.lint.MetaInfo().IsIndexingComplete() {
		ls.changeFileNonLocked(filename, contents)
		return
	}

	rawContents := []byte(contents)

	// just parse file, do not fully analyze it as indexing is not yet done
	res, err := ls.indexer.ParseContents(workspace.FileInfo{
		Name:     filename,
		Contents: rawContents,
	})
	if err != nil {
		log.Printf("Could not parse %s: %s", filename, err.Error())
		lintdebug.Send("Could not parse %s: %s", filename, err.Error())
		return
	}

	ls.openMapMutex.Lock()
	ls.openMap[filename] = openedFile{
		rootNode: res.RootNode,
		file:     workspace.NewFile(filename, rawContents),
	}
	ls.openMapMutex.Unlock()
}

// Handle changed contents of a file in the editor
func (ls *LangServer) changeFile(filename, contents string) {
	ls.changingMutex.Lock()
	defer ls.changingMutex.Unlock()

	ls.changeFileNonLocked(filename, contents)
}

func (ls *LangServer) changeFileNonLocked(filename, contents string) {
	// if !ls.lint.MetaInfo().IsIndexingComplete() {
	// 	return
	// }
	//
	// // parse file, update index for it, and then generate diagnostics based on new index
	// ls.lint.MetaInfo().SetIndexingComplete(false)
	//
	// res, err := ls.indexer.ParseContents(workspace.FileInfo{
	// 	Name:     filename,
	// 	Contents: []byte(contents),
	// })
	// if err != nil {
	// 	log.Printf("Could not parse %s: %s", filename, err.Error())
	// 	lintdebug.Send("Could not parse %s: %s", filename, err.Error())
	// 	return
	// }
	//
	// // w.UpdateMetaInfo()
	//
	// ls.lint.MetaInfo().SetIndexingComplete(true)
	//
	// newWalker := linter.NewWalkerForLangServer(ls.lint.MetaInfo(), ls.lint.Config(), linter.NewWorkerContext(), res.Walker)
	//
	// newWalker.InitCustom()
	// res.RootNode.Walk(newWalker)
	//
	// linter.AnalyzeFileRootLevel(res.RootNode, newWalker)
	//
	// ls.openMapMutex.Lock()
	// f := openedFile{
	// 	rootNode: res.RootNode,
	// 	// scopes:   w.Scopes,
	// 	file: res.Walker.File(),
	// }
	// ls.openMap[filename] = f
	// ls.openMapMutex.Unlock()

	ls.flushReports(filename, nil)
}

func (ls *LangServer) flushReports(filename string, d *linter.RootWalker) {
	// diag := d.Diagnostics
	// if len(diag) == 0 && diag == nil {
	// 	diag = make([]vscode.Diagnostic, 0)
	// }

	var diags []vscode.Diagnostic

	for i := 0; i < 10; i++ {
		diags = append(diags, vscode.Diagnostic{
			Range: vscode.Range{
				Start: vscode.Position{
					Line:      i,
					Character: 0,
				},
				End: vscode.Position{
					Line:      i,
					Character: 3,
				},
			},
			Severity: vscode.Error,
			Code:     "unused",
			Source:   "noverify",
			Message:  "some unused error",
			Tags:     []int{1},
		})
	}

	js, _ := json.Marshal(&methodCall{
		JSONRPC: "2.0",
		Method:  "textDocument/publishDiagnostics",
		Params: &vscode.PublishDiagnosticsParams{
			URI:         uri.File(filename),
			Diagnostics: diags,
		},
	})
	lintdebug.Send("reports %s", string(js))
	err := writeMessage(&methodCall{
		JSONRPC: "2.0",
		Method:  "textDocument/publishDiagnostics",
		Params: &vscode.PublishDiagnosticsParams{
			URI:         uri.File(filename),
			Diagnostics: diags,
		},
	})

	if err != nil {
		lintdebug.Send("error send diagnostic: %v", err)
	}
}

func (ls *LangServer) handleTextDocumentDidChange(req *baseRequest) error {
	var params vscode.TextDocumentDidChangeParams
	if err := json.Unmarshal(req.Params, &params); err != nil {
		return err
	}

	// if len(params.ContentChanges) != 1 {
	// 	lintdebug.Send("Unexpected number of content changes: %d", len(params.ContentChanges))
	// 	return nil
	// }

	u := params.TextDocument.URI
	lintdebug.Send("uri: %s", string(u))
	if u == "" {
		lintdebug.Send("uri empty")
		return nil
	}
	//
	// if isFileScheme(u) {
	ls.changeFile(u.Filename(), params.ContentChanges[0].Text)
	// }

	lintdebug.Send("file changed %s", params.ContentChanges)

	return nil
}
