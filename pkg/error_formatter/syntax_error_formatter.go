package error_formatter

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
	case strings.Contains(msg, "missing '}'") || (strings.Contains(msg, "extraneous input '<EOF>' expecting") && strings.Contains(msg, "}")):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: Unclosed function block. Missing closing brace '}' for function body",
			actualLine, actualCol, codeLine, caretLine)

	case strings.Contains(msg, "extraneous input") && !isValidIdentifier(text):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: Invalid identifier '%s'",
			actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "missing ';'"):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: missing ';' at '%s'", actualLine, actualCol, codeLine, caretLine, text)

	case strings.Contains(msg, "mismatched input") && isTypeKeyword(text):
		if strings.Contains(msg, "expecting Identifier") {
			fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: missing identifier after type '%s'",
				actualLine, actualCol, codeLine, caretLine, text)
		} else {
			prevLine := ""
			if actualLine > 1 && actualLine-1 <= len(h.Lines) {
				prevLine = h.Lines[actualLine-2]
			}
			if !strings.Contains(prevLine, ";") && !strings.Contains(prevLine, "}") {
				fullMessage = fmt.Sprintf("line %d:%d\n%s\nSyntax Error: missing ';' at end of declaration\n%s\n%s\nNote: this affects the next declaration on line %d",
					actualLine-1, len(prevLine),
					prevLine,
					codeLine,
					caretLine,
					actualLine)
			} else {
				fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: unexpected type keyword '%s'. Check for missing identifier or invalid syntax",
					actualLine, actualCol, codeLine, caretLine, text)
			}
		}

	case strings.Contains(msg, "mismatched input") && !isTypeKeyword(text) && len(text) > 0:
		if strings.Contains(msg, "expecting ')'") {
			fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: missing closing parenthesis ')'",
				actualLine, actualCol, codeLine, caretLine)
		} else if strings.Contains(msg, "expecting '{'") {
			fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: missing opening brace '{' for function body",
				actualLine, actualCol, codeLine, caretLine)
		} else if strings.Contains(msg, "expecting {<EOF>, 'int', 'float', 'bool', 'char', 'void'}") {
			fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: unexpected type '%s'. ",
				actualLine, actualCol, codeLine, caretLine, text)
		} else {
			fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: unexpected token '%s'.",
				actualLine, actualCol, codeLine, caretLine, text)
		}

	case strings.Contains(msg, "missing ']'") || (strings.Contains(msg, "expecting") && strings.Contains(msg, "]")):
		fullMessage = fmt.Sprintf("line %d:%d\n%s\n%s\nSyntax Error: incomplete array declaration. Missing closing bracket ']' at '%s'", actualLine, actualCol, codeLine, caretLine, text)

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

// Add this helper function
func isTypeKeyword(text string) bool {
	types := []string{"int", "float", "bool", "char", "void"}
	for _, t := range types {
		if text == t {
			return true
		}
	}
	return false
}

// Add this helper function
func isValidIdentifier(text string) bool {
	if len(text) == 0 {
		return false
	}

	// First character must be [a-zA-Z_]
	first := text[0]
	if !((first >= 'a' && first <= 'z') ||
		(first >= 'A' && first <= 'Z') ||
		first == '_') {
		return false
	}

	// Check for any characters that are not [a-zA-Z0-9_]
	for _, c := range text[1:] {
		// Check if character is NOT in allowed set
		isLetter := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		isDigit := c >= '0' && c <= '9'
		isUnderscore := c == '_'

		if !isLetter && !isDigit && !isUnderscore {
			// Character is a special symbol or other invalid character
			return false
		}
	}
	return true
}

// Required interface methods for antlr.ErrorListener
func (h *SyntaxErrorHandler) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (h *SyntaxErrorHandler) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (h *SyntaxErrorHandler) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}
