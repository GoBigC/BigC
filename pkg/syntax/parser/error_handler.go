package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// SyntaxErrorHandler handles syntax errors during parsing
type SyntaxErrorHandler struct {
	Errors []string
	Lines  []string // source lines from the file
}

func NewSyntaxErrorHandler(lines []string) *SyntaxErrorHandler {
	return &SyntaxErrorHandler{
		Errors: make([]string, 0),
		Lines:  lines,
	}
}

func (h *SyntaxErrorHandler) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	actualLine := line
	actualCol := column
	text := "<unknown>"

	if token, ok := offendingSymbol.(*antlr.CommonToken); ok && token != nil {
		text = token.GetText()
		if text == "<EOF>" {
			if parser, ok := recognizer.(antlr.Parser); ok {
				stream := parser.GetTokenStream()
				index := stream.Index()
				if index > 0 {
					prev := stream.Get(index - 1)
					text = prev.GetText()
					actualLine = prev.GetLine()
					actualCol = prev.GetColumn() + len(prev.GetText())
				}
			}
		}
	}

	codeLine := ""
	caretLine := ""
	if actualLine > 0 && actualLine <= len(h.Lines) {
		codeLine = h.Lines[actualLine-1]
		caretLine = strings.Repeat(" ", actualCol) + "^"
	}

	// Enhanced error message handling
	var fullMessage string
	switch {
	case strings.Contains(msg, "missing ';'"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: missing ';' at '%s'", actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "mismatched input"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: mismatched input '%s'. Check for unexpected tokens or typos.", actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "no viable alternative"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: unexpected structure near '%s'. Possibly an invalid expression or statement.", actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "extraneous input"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: extraneous input '%s'. Remove unnecessary characters or tokens.", actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "expecting"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: expected token missing near '%s'. %s", actualLine, actualCol, codeLine, caretLine, text, msg)

	default:
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\n%s", actualLine, actualCol, codeLine, caretLine, msg)
	}

	h.Errors = append(h.Errors, fullMessage)
}

// Required interface methods for antlr.ErrorListener
func (h *SyntaxErrorHandler) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (h *SyntaxErrorHandler) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (h *SyntaxErrorHandler) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}
