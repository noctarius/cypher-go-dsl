package cypher_go_dsl

type Limit struct {
	limitAmount NumberLiteral
	key         string
	notNil      bool
	err         error
}

func LimitCreate(number int) Limit {
	if number == 0 {
		return Limit{}
	}
	literal := NumberLiteralCreate(number)
	l := Limit{limitAmount: literal}
	l.key = getAddress(&l)
	return l
}

func (l Limit) getError() error {
	return l.err
}

func (l Limit) isNotNil() bool {
	return l.notNil
}

func (l Limit) getKey() string {
	return l.key
}

func (l Limit) accept(visitor *CypherRenderer) {
	(*visitor).enter(l)
	l.limitAmount.accept(visitor)
	(*visitor).leave(l)
}

func (l Limit) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" LIMIT ")
}

func (l Limit) leave(renderer *CypherRenderer) {
}
