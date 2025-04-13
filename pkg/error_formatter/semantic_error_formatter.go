package error_formatter

import (
	"fmt"
	"strings"
)

// Formats semantic errors in Rust style
type SemanticErrorFormatter struct {
	SourceLines []string
}

// Creates a new formatter with source lines
func NewSemanticErrorFormatter(sourceLines []string) *SemanticErrorFormatter {
	return &SemanticErrorFormatter{
		SourceLines: sourceLines,
	}
}

// Formats a single semantic error
func (f *SemanticErrorFormatter) FormatError(rawError string) string {
	// Extract line number and error message
	parts := strings.SplitN(rawError, ": ", 2)
	if len(parts) != 2 {
		return rawError
	}

	var line int
	_, err := fmt.Sscanf(parts[0], "Line %d", &line)
	if err != nil {
		return rawError
	}

	errorMsg := parts[1]
	codeLine := ""
	if line > 0 && line <= len(f.SourceLines) {
		codeLine = f.SourceLines[line-1]
	} else {
		return rawError
	}

	column := findErrorPosition(codeLine, errorMsg)
	caretLine := strings.Repeat(" ", column) + "^"
	errorType := "Semantic Error"
	detailedMsg := errorMsg

	// Format the error message
	fullMessage := fmt.Sprintf("line %d:%d\n%s\n%s\n%s: %s",
		line, column, codeLine, caretLine, errorType, detailedMsg)

	return fullMessage
}

// Formats all semantic errors
func (f *SemanticErrorFormatter) FormatErrors(rawErrors []string) []string {
	formatted := make([]string, len(rawErrors))
	for i, rawError := range rawErrors {
		formatted[i] = f.FormatError(rawError)
	}
	return formatted
}

func findErrorPosition(codeLine, errorMsg string) int {
	// Specific error handling
	if strings.Contains(errorMsg, "argument count mismatch") {
		funcName := extractFunctionCall(errorMsg)
		if funcName != "" {
			callPos := strings.Index(codeLine, funcName+"(")
			if callPos >= 0 {
				return callPos + len(funcName)
			}
		}
		parenPos := strings.Index(codeLine, "(")
		if parenPos >= 0 {
			return parenPos
		}
	}

	if strings.Contains(errorMsg, "return type mismatch") {
		retPos := strings.Index(codeLine, "return ")
		if retPos >= 0 {
			return retPos
		}
	}

	if strings.Contains(errorMsg, "array access on non-array type") {
		ids := extractIdentifiers(errorMsg)
		for _, id := range ids {
			bracketPos := strings.Index(codeLine, id+"[")
			if bracketPos >= 0 {
				return bracketPos + len(id)
			}
			idPos := strings.Index(codeLine, id)
			if idPos >= 0 {
				return idPos
			}
		}
	}

	// General identifier handling
	identifiers := extractIdentifiers(errorMsg)
	var targetIdentifier string
	if strings.Contains(errorMsg, "undefined symbol: ") {
		parts := strings.SplitN(errorMsg, "undefined symbol: ", 2)
		if len(parts) > 1 {
			targetIdentifier = strings.Fields(parts[1])[0]
			targetIdentifier = strings.Trim(targetIdentifier, ".,():;'\"")
		}
	}

	if targetIdentifier != "" {
		pos := strings.Index(codeLine, targetIdentifier)
		if pos >= 0 {
			return pos
		}
	}

	for _, id := range identifiers {
		pos := strings.Index(codeLine, id)
		if pos >= 0 {
			return pos
		}
	}

	// Operator handling
	operators := []string{"==", "!=", "<=", ">=", "<", ">", "+", "-", "*", "/", "&&", "||", "!"}
	for _, op := range operators {
		if strings.Contains(errorMsg, fmt.Sprintf("operator '%s'", op)) || strings.Contains(errorMsg, op) {
			pos := strings.Index(codeLine, op)
			if pos >= 0 {
				return pos
			}
		}
	}

	// Default to first non-whitespace character
	for i, c := range codeLine {
		if c != ' ' && c != '\t' {
			return i
		}
	}

	return 0
}

// Extracts function name from error message
func extractFunctionCall(errorMsg string) string {
	pattern := "function '"
	if idx := strings.Index(errorMsg, pattern); idx != -1 {
		start := idx + len(pattern)
		end := strings.Index(errorMsg[start:], "'")
		if end != -1 {
			return errorMsg[start : start+end]
		}
	}
	return ""
}

// Extracts identifiers from error message
func extractIdentifiers(errorMsg string) []string {
	identifiers := []string{}
	patterns := []string{
		"undefined symbol: ",
		"variable ",
		"function ",
		"type mismatch: expected ",
		", got ",
		"array ",
		"parameter ",
		"type '",
	}
	for _, pattern := range patterns {
		if strings.Contains(errorMsg, pattern) {
			parts := strings.SplitN(errorMsg, pattern, 2)
			if len(parts) > 1 {
				word := strings.Split(parts[1], " ")[0]
				word = strings.Split(word, "'")[0]
				word = strings.Trim(word, ".,():;")
				if word != "" && !isKeyword(word) {
					identifiers = append(identifiers, word)
				}
			}
		}
	}
	return identifiers
}

// Checks if a word is a keyword
func isKeyword(word string) bool {
	keywords := map[string]bool{"if": true, "else": true, "while": true, "return": true, "int": true, "bool": true}
	return keywords[word]
}
