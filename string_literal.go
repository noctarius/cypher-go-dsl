package cypher_go_dsl

type StringLiteral struct {
	content string
	key     string
	notNil  bool
	err     error
}

func StringLiteralCreate(content string) StringLiteral {
	stringLiteral := StringLiteral{
		content: content,
		notNil:  true,
	}
	stringLiteral.key = getAddress(&stringLiteral)
	return stringLiteral
}

func (s StringLiteral) getError() error {
	return s.err
}

func (s StringLiteral) isNotNil() bool {
	return s.notNil
}

func escapeStringLiteral(value string) string {
	return "'" + value + "'"
}

func (s StringLiteral) getKey() string {
	return s.key
}

func (s StringLiteral) GetExpressionType() ExpressionType {
	return "StringLiteral"
}

func (s StringLiteral) GetContent() interface{} {
	return s.content
}

func (s StringLiteral) AsString() string {
	return s.content
}

func (s StringLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s StringLiteral) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeStringLiteral(s.AsString()))
}

func (s StringLiteral) leave(renderer *CypherRenderer) {
}
