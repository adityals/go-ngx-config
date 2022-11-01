package lexer

import (
	"bufio"
	"bytes"
	"io"

	"github.com/adityals/go-ngx-config/internal/token"
)

type runeCheck func(rune) bool

type Lexer struct {
	File   string
	Line   int
	Column int
	Reader *bufio.Reader
	Latest token.Token
}

func NewStringLexer(confString string) *Lexer {
	return NewLexer(bytes.NewBuffer([]byte(confString)))
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		Line:   1,
		Reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Scan() token.Token {
	l.Latest = l.getNextToken()
	return l.Latest
}

func (l *Lexer) getNextToken() token.Token {
reTokenizing:
	ch := l.peek()
	switch {
	case isSpace(ch):
		l.skipWhitespace()
		goto reTokenizing
	case isEOF(ch):
		return l.NewToken(token.EOF).Lit(string(l.read()))
	case ch == ';':
		return l.NewToken(token.Semicolon).Lit(string(l.read()))
	case ch == '{':
		return l.NewToken(token.BlockStart).Lit(string(l.read()))
	case ch == '}':
		return l.NewToken(token.BlockEnd).Lit(string(l.read()))
	case ch == '#':
		return l.scanComment()
	case ch == '$':
		return l.scanVariable()
	case isQuote(ch):
		return l.scanQuotedString(ch)
	default:
		return l.scanKeyword()
	}
}

func (l *Lexer) NewToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:   tokenType,
		Line:   l.Line,
		Column: l.Column,
	}
}

func (l *Lexer) peek() rune {
	r, _, _ := l.Reader.ReadRune()
	_ = l.Reader.UnreadRune()
	return r
}

func (l *Lexer) read() rune {
	ch, _, err := l.Reader.ReadRune()

	if err != nil {
		return rune(token.EOF)
	}

	if ch == '\n' {
		l.Column = 1
		l.Line++
	} else {
		l.Column++
	}

	return ch
}

func (l *Lexer) readWhile(while runeCheck) string {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		if ch := l.peek(); while(ch) {
			buf.WriteRune(l.read())
		} else {
			break
		}
	}

	// unread the latest char we consume.
	return buf.String()
}

func (l *Lexer) readUntil(until runeCheck) string {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		if ch := l.peek(); isEOF(ch) {
			break
		} else if until(ch) {
			break
		} else {
			buf.WriteRune(l.read())
		}
	}

	return buf.String()
}

func (l *Lexer) skipWhitespace() {
	l.readWhile(isSpace)
}

func (l *Lexer) scanComment() token.Token {
	return l.NewToken(token.Comment).Lit(l.readUntil(isEndOfLine))
}

func (l *Lexer) scanKeyword() token.Token {
	return l.NewToken(token.Keyword).Lit(l.readUntil(isKeywordTerminator))
}

func (l *Lexer) scanVariable() token.Token {
	return l.NewToken(token.Variable).Lit(l.readUntil(isKeywordTerminator))
}

/**
\” – To escape “ within double quoted string.
\\ – To escape the backslash.
\n – To add line breaks between string.
\t – To add tab space.
\r – For carriage return.
*/
func (l *Lexer) scanQuotedString(delimiter rune) token.Token {
	var buf bytes.Buffer
	tok := l.NewToken(token.QuotedString)

	buf.WriteRune(l.read()) // consume delimiter

	for {
		ch := l.read()

		if ch == rune(token.EOF) {
			panic("unexpected end of file while scanning a string, maybe an unclosed quote?")
		}

		if ch == '\\' {
			if needsEscape(l.peek(), delimiter) {
				switch l.read() {
				case 'n':
					buf.WriteRune('\n')
				case 'r':
					buf.WriteRune('\r')
				case 't':
					buf.WriteRune('\t')
				case '\\':
					buf.WriteRune('\\')
				case delimiter:
					buf.WriteRune(delimiter)
				}
				continue
			}
		}

		buf.WriteRune(ch)
		if ch == delimiter {
			break
		}
	}

	return tok.Lit(buf.String())
}
