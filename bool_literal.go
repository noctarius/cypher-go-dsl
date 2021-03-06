package cypher_go_dsl

type BooleanLiteral struct {
	content bool
	key     string
	notNil  bool
	err     error
}

func BooleanLiteralCreate(content bool) BooleanLiteral {
	b := BooleanLiteral{
		content: content,
		notNil:  true,
	}
	b.key = getAddress(&b)
	return b
}

func (b BooleanLiteral) getError() error {
	return b.err
}

func (b BooleanLiteral) isNotNil() bool {
	return b.notNil
}

func (b BooleanLiteral) getKey() string {
	return b.key
}

func (b BooleanLiteral) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (b BooleanLiteral) GetContent() interface{} {
	return b.content
}

func (b BooleanLiteral) AsString() string {
	if b.content {
		return "true"
	}
	return "false"
}

func (b BooleanLiteral) accept(visitor *CypherRenderer) {
	(*visitor).enter(b)
	(*visitor).leave(b)
}

func (b BooleanLiteral) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(b.AsString())
}

func (b BooleanLiteral) leave(renderer *CypherRenderer) {
}
