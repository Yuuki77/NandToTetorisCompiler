package src

import (
	"bufio"
	"os"
	"regexp"
	"strings"
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
	if tokenizer.tokenType != StringConst {
		panic("token type should be val")
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
	// コンストラクタの関数内で、構造体をnew
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

	// justString := strings.Join(tokenizer.keywordLists, " ")

	// concatSymbols := append(tokenizer.keywordLists, tokenizer.symbolLists...)

	// joinedString := strings.Join(tokenizer.symbolLists, ",")
	// fmt.Println("joined string", joinedString)

	// // 以下、構造体の各フィールドを引数で受け取った値に設定
	// l.Name = name
	// l.LangType = langType
	// // 構造体のインスタンスを返す
	// return l

	// files, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("cpy:", string(files))

	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	re := regexp.MustCompile(`//|/.*|.*/`)

	for scanner.Scan() {
		text := scanner.Text()
		// normal comment out
		text = strings.TrimSpace(text)

		if re.MatchString(text) {
			continue
		}

		if len(text) == 0 {
			continue
		}

		line := tokenizer.Split(text)
		tokenizer.lines = append(tokenizer.lines, line...)

	}
	return tokenizer
}

// HasMoreTokens checks if has more token
func (tokenizer JackTokenizer) hasMoreTokens() bool {
	//code
	return false
}

// Advance gets next token and it can be called
func (tokenizer JackTokenizer) Advance() bool {
	if tokenizer.hasMoreTokens() {
		return true
	} else {
		return false
	}
}

func (tokenizer JackTokenizer) Split(line string) []string {
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == string(e) {
			return true
		}
	}
	return false
}

type TokenType int

const (
	KEYWORD TokenType = iota
	SYMBOL
	IDENTIFIER
	StringConst
)

type KeyWord int

const (
	CLASS KeyWord = iota
	METHOD
	FUNCTION
	CONSTRUCTOR
	INT
	BOOLEAN
	CHAR
	VOID
	STATIC
	FILED
	LET
	DL
	IF
	ELSE
	WHILE
	RETURN
	TRUE
	FALSE
	NULL
	THIS
)
