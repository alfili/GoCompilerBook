package parser

import (
	"gocompilerbook/ast"
	"gocompilerbook/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 0;
	let y = 10;

	let test = 121212;
	`

	lex := lexer.New(input)

	pars := New(lex)

	prog := pars.ParseProgram()

	if prog == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"test"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}
