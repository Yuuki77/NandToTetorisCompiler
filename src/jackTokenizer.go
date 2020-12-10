package src

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

type JackTokenizer struct {
	tokenType    TokenType
	keyWord      KeyWord
	symbol       string
	identifier   string
	intVal       int64
	stringVal    string
	keywordLists []string
	symbolLists  []string
	lines        []string
	currentIndex int
}

func (tokenizer JackTokenizer) GetKeyWord() KeyWord {
	if tokenizer.tokenType != KEYWORD {
		panic("token type should be keyword")
	}

	return tokenizer.keyWord
}

func (tokenizer JackTokenizer) GetSymbol() string {
	if tokenizer.tokenType != SYMBOL {
		panic("token type should be symbol")
	}

	return tokenizer.symbol
}

func (tokenizer JackTokenizer) GetIdentifier() string {
	if tokenizer.tokenType != IDENTIFIER {
		panic("token type should be identifier")
	}

	return tokenizer.identifier
}

func (tokenizer JackTokenizer) GetIntVal() int64 {
	if tokenizer.tokenType != INTCONST {
		panic("token type should be int")
	}

	return tokenizer.intVal
}

func (tokenizer JackTokenizer) GetStringVal() string {
	if tokenizer.tokenType != StringConst {
		panic("token type should be val")
	}

	return tokenizer.stringVal
}

func New(filePath string) *JackTokenizer {
	tokenizer := new(JackTokenizer)
	tokenizer.keywordLists = []string{
		"class", "constructor", "function", "method", "field", "static",
		"var", "int", "char", "boolean", "void", "true", "false", "null",
		"this", "let", "do", "if", "else", "while", "return",
	}

	tokenizer.symbolLists = []string{
		"{", "}", "(", ")", "[", "]",
		".", ",", ";", "+", "-", "*", "/",
		"&", "|", "<", ">", "=", "~",
	}

	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	re := regexp.MustCompile(`^//|^/.*`)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if re.MatchString(text) {
			continue
		}

		if len(text) == 0 {
			continue
		}

		// trim the inline comment

		text = trimInLIneComment(text)
		pretty.Println(text)

		line := tokenizer.split(text)
		tokenizer.lines = append(tokenizer.lines, line...)
	}

	for tokenizer.Advance() {
	}

	return tokenizer
}

// HasMoreTokens checks if has more token
func (tokenizer JackTokenizer) hasMoreTokens() bool {
	return tokenizer.currentIndex < len(tokenizer.lines)
}
func trimInLIneComment(text string) string {
	if strings.Contains(text, "/") {
		index := strings.Index(text, "/")
		return text[:index]
	}
	return text
}

// Advance gets next token and it can be called
func (tokenizer *JackTokenizer) Advance() bool {
	if tokenizer.hasMoreTokens() {
		tokenizer.tokenType = tokenizer.getTokenType()
		// pretty.Println("current word is %v, tokenType %v", tokenizer.getCurrentWord(), tokenizer.tokenType)

		if tokenizer.isKeyword() {
			tokenizer.keyWord = tokenizer.convertStringToKeyWord()
			pretty.Println("keyword is", tokenizer.keyWord)
		} else if tokenizer.isSymbol() {
			tokenizer.symbol = tokenizer.getCurrentWord()
			pretty.Println("symbol is", tokenizer.symbol)

		} else if tokenizer.isStringConst() {
			tokenizer.stringVal = tokenizer.getCurrentWord()
			pretty.Println("string val is", tokenizer.stringVal)

		} else if tokenizer.isIntConst() {
			tokenizer.intVal = tokenizer.GetIntVal()
			pretty.Println("int const is", tokenizer.intVal)
		} else {
			tokenizer.identifier = tokenizer.getCurrentWord()
			pretty.Println("identifier is", tokenizer.identifier)
		}

		// pretty.Println("current index is  %v", tokenizer.currentIndex)

		tokenizer.currentIndex++

		return true
	}
	return false
}

func (tokenizer *JackTokenizer) getTokenType() TokenType {
	if tokenizer.isKeyword() {
		return KEYWORD
	} else if tokenizer.isSymbol() {
		return SYMBOL
	} else if tokenizer.isStringConst() {
		return StringConst
	} else if tokenizer.isIntConst() {
		return INTCONST
	}

	return IDENTIFIER
}

func (tokenizer JackTokenizer) convertStringToKeyWord() KeyWord {
	word := tokenizer.getCurrentWord()

	switch word {
	case "class":
		return CLASS
	case "method":
		return METHOD
	case "function":
		return FUNCTION
	case "constructor":
		return CONSTRUCTOR
	case "int":
		return INT
	case "boolean":
		return BOOLEAN
	case "char":
		return CHAR
	case "void":
		return VOID
	case "var":
		return VAR
	case "static":
		return STATIC
	case "field":
		return FILED
	case "let":
		return LET
	case "do":
		return DO
	case "if":
		return IF
	case "else":
		return ELSE
	case "while":
		return WHILE
	case "return":
		return RETURN
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "null":
		return NULL
	case "this":
		return THIS
	}
	panic("what" + word)
}
func (tokenizer JackTokenizer) split(line string) []string {
	fields := []string{}

	last := 0
	stringStart := false
	for i := 0; i < len(line); i++ {
		value := string(line[i])
		isString := value == "\""

		if isString {
			stringStart = !stringStart
			if stringStart {
				continue
			} else {
				fields = append(fields, line[last:i+1])
				last = i + 1
			}
		}

		if stringStart {
			continue
		}

		if value == " " {
			if last == i {
				last = i + 1
				continue
			}

			fields = append(fields, line[last:i])
			last = i + 1

		} else if contains(tokenizer.symbolLists, value) {
			if i != last {
				fields = append(fields, line[last:i])
			}
			fields = append(fields, line[i:i+1])

			last = i + 1
		}
	}

	return fields
}

func (tokenizer JackTokenizer) isKeyword() bool {
	return contains(tokenizer.keywordLists, tokenizer.getCurrentWord())
}

func (tokenizer JackTokenizer) isSymbol() bool {
	return contains(tokenizer.symbolLists, tokenizer.getCurrentWord())
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == string(e) {
			return true
		}
	}
	return false
}

func (tokenizer JackTokenizer) isStringConst() bool {
	return string(tokenizer.getCurrentWord()[0]) == "\""
}
func (tokenizer JackTokenizer) isIntConst() bool {
	_, err := strconv.Atoi(tokenizer.getCurrentWord())
	return err == nil
}

func (tokenizer JackTokenizer) getInt() int {
	val, _ := strconv.Atoi(tokenizer.getCurrentWord())
	return val
}

func (tokenizer JackTokenizer) getCurrentWord() string {
	return tokenizer.lines[tokenizer.currentIndex]
}

type TokenType string

const (
	KEYWORD     = TokenType("token")
	SYMBOL      = TokenType("symbol")
	IDENTIFIER  = TokenType("identifer")
	INTCONST    = TokenType("InConst")
	StringConst = TokenType("StringConst")
)

type KeyWord string

const (
	CLASS       = KeyWord("class")
	METHOD      = KeyWord("method")
	FUNCTION    = KeyWord("function")
	CONSTRUCTOR = KeyWord("contactor")
	INT         = KeyWord("int")
	BOOLEAN     = KeyWord("boolean")
	CHAR        = KeyWord("char")
	VOID        = KeyWord("void")
	VAR         = KeyWord("var")
	STATIC      = KeyWord("static")
	FILED       = KeyWord("filed")
	LET         = KeyWord("let")
	DO          = KeyWord("do")
	IF          = KeyWord("if")
	ELSE        = KeyWord("else")
	WHILE       = KeyWord("while")
	RETURN      = KeyWord("return")
	TRUE        = KeyWord("true")
	FALSE       = KeyWord("false")
	NULL        = KeyWord("null")
	THIS        = KeyWord("this")
)
