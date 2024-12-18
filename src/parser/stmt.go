package parser

import (
	"github.com/Dorakokce/DoraScriptGo/src/ast"
	"github.com/Dorakokce/DoraScriptGo/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	var assignedValue ast.Expr

	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration expected to find variable name").Value

	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		explicitType = parse_type(p, default_bp)
	}

	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Missing either right hand side in var declaration or expilicit type")
	}

	p.expect(lexer.SEMI_COLON)

	if isConstant && assignedValue == nil {
		panic("Cannot define constant without providing value")
	}

	return ast.VarDeclStmt{
		ExplicitType:  explicitType,
		VarName:       varName,
		IsConstant:    isConstant,
		AssignedValue: assignedValue,
	}
}

func parse_struct_decl(p *parser) ast.Stmt {
	p.expect(lexer.STRUCT)
	structName := p.expect(lexer.IDENTIFIER).Value

	return ast.StructDeclStmt{
		Properties: parse_obj_prop(p),
		StructName: structName,
	}
}

func parse_obj_prop(p *parser) map[string]ast.ObjectProperty {
	p.expect(lexer.OPEN_CURLY)

	props := map[string]ast.ObjectProperty{}

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		static := false
		if p.currentTokenKind() == lexer.STATIC {
			static = true
		}
		name := p.expect(lexer.IDENTIFIER).Value
		p.expect(lexer.COLON)
		propType := parse_type(p, default_bp)
		p.expect(lexer.SEMI_COLON)

		props[name] = ast.ObjectProperty{
			IsStatic: static,
			Type:     propType,
		}
	}

	p.expect(lexer.CLOSE_CURLY)

	return props
}
