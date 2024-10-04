package parser

import (
	"testing"

	"github.com/quinsberry/monkey-interpreter/internal/ast"
	"github.com/quinsberry/monkey-interpreter/internal/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let asd = 83;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if stLen := len(program.Statements); stLen != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", stLen)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"asd"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if tl := s.TokenLiteral(); tl != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", tl)
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if val := letStmt.Name.Value; val != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, val)
		return false
	}

	if tl := letStmt.Name.TokenLiteral(); tl != name {
		t.Errorf("letStmt.Name.TokenLiteral not '%s'. got=%s", name, tl)
		return false
	}
	return true
}
