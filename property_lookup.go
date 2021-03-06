package cypher_go_dsl

type PropertyLookup struct {
	ExpressionContainer
	propertyKeyName string
	key             string
	notNil          bool
	err             error
}

func (p PropertyLookup) getError() error {
	return p.err
}

func (p PropertyLookup) isNotNil() bool {
	return p.notNil
}

func (p PropertyLookup) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	visitor.leave(p)
}

func (p PropertyLookup) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(".")
	renderer.builder.WriteString(p.GetPropertyKeyName())
}

func (p PropertyLookup) leave(renderer *CypherRenderer) {
}

func (p PropertyLookup) getKey() string {
	return p.key
}

func (p PropertyLookup) GetExpressionType() ExpressionType {
	panic("implement me")
}

func (p PropertyLookup) GetPropertyKeyName() string {
	return p.propertyKeyName
}

func (p *PropertyLookup) Wrap() {
	p.expression = p
}
