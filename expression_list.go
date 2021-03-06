package cypher_go_dsl

type ExpressionList struct {
	expressions []Expression
	key         string
	notNil      bool
	err         error
}

func ExpressionListCreate(expression []Expression) ExpressionList {
	e := ExpressionList{
		expressions: expression,
		notNil:      true,
	}
	e.key = getAddress(&e)
	return e
}

func (e ExpressionList) GetExpressionType() ExpressionType {
	panic("implement me")
}

func (e ExpressionList) getError() error {
	return e.err
}

func (e ExpressionList) isNotNil() bool {
	return e.notNil
}

func (e ExpressionList) getKey() string {
	return e.key
}

func (e ExpressionList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(expression)
}

func (e ExpressionList) accept(visitor *CypherRenderer) {
	(*visitor).enter(e)
	for _, expression := range e.expressions {
		e.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(e)
}

func (e ExpressionList) enter(renderer *CypherRenderer) {
}

func (e ExpressionList) leave(renderer *CypherRenderer) {
}

func NewExpressionList(expression ...Expression) ExpressionList {
	expressions := make([]Expression, len(expression))
	expressions = append(expressions, expression...)
	return ExpressionList{expressions: expressions}
}
