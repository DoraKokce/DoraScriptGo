package parser

import (
	"fmt"

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
	var properties = map[string]ast.StructProperty{}
	var methods = map[string]ast.StructMethod{}
	var structName = p.expect(lexer.IDENTIFIER).Value

	p.expect(lexer.OPEN_CURLY)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		isStatic := false
		var propertyName string
		if p.currentTokenKind() == lexer.STATIC {
			isStatic = true
			p.expect(lexer.STATIC)
		}

		if p.currentTokenKind() == lexer.IDENTIFIER {
			propertyName = p.expect(lexer.IDENTIFIER).Value
			p.expectError(lexer.COLON, "Expected to find colon following after struct property")
			structType := parse_type(p, default_bp)
			p.expect(lexer.SEMI_COLON)

			_, exists := properties[propertyName]

			if exists {
				panic(fmt.Sprintf("Property %s has already bean declared", propertyName))
			}

			properties[propertyName] = ast.StructProperty{
				IsStatic: isStatic,
				Type:     structType,
			}

			continue
		}

		panic("Cannot curently handle methods inside struct decl")
	}

	p.expect(lexer.CLOSE_CURLY)

	return ast.StructDeclStmt{
		Properties: properties,
		Methods:    methods,
		StructName: structName,
	}
}
