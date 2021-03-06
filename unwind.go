package cypher_go_dsl

type Unwind struct {
	expressionToUnwind Expression
	variable           string
	key                string
	notNil             bool
	err                error
}

func UnwindCreate(expressionToUnwind Expression, variable string) Unwind {
	var expression Expression
	if aliased, isAliased := expressionToUnwind.(Aliased); isAliased {
		expression = aliased.AsName()
	} else {
		expression = expressionToUnwind
	}
	unwind := Unwind{
		expressionToUnwind: expression,
		variable:           variable,
		notNil:             true,
	}
	unwind.key = getAddress(&unwind)
	return unwind
}

func (u Unwind) getError() error {
	return u.err
}

func (u Unwind) accept(visitor *CypherRenderer) {
	visitor.enter(u)
	u.expressionToUnwind.accept(visitor)
	visitor.leave(u)
}

func (u Unwind) enter(renderer *CypherRenderer) {
	renderer.append("UNWIND ")
}

func (u Unwind) leave(renderer *CypherRenderer) {
	renderer.append(" AS ")
	renderer.append(u.variable)
	renderer.append(" ")
}

func (u Unwind) getKey() string {
	panic("implement me")
}

func (u Unwind) isNotNil() bool {
	panic("implement me")
}
