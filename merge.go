package cypher_go_dsl

type Merge struct {
	pattern               Pattern
	onCreateOrMatchEvents []Visitable
	key                   string
	notNil                bool
	err                   error
}

var BLANK = StringLiteralCreate(" ")

func MergeCreate(pattern Pattern) Merge {
	m := Merge{
		pattern:               pattern,
		onCreateOrMatchEvents: make([]Visitable, 0),
		notNil:                true,
	}
	m.key = getAddress(&m)
	return m
}

func MergeCreate1(pattern Pattern, mergeActions []MergeAction) Merge {
	onCreateOrMatchEvents := make([]Visitable, 0)
	onCreateOrMatchEvents = append(onCreateOrMatchEvents, BLANK)
	for _, mergeAction := range mergeActions {
		onCreateOrMatchEvents = append(onCreateOrMatchEvents, mergeAction)
	}
	m := Merge{
		pattern:               pattern,
		onCreateOrMatchEvents: onCreateOrMatchEvents,
		notNil:                true,
	}
	m.key = getAddress(&m)
	return m
}

func (m Merge) hasEvent() bool {
	return len(m.onCreateOrMatchEvents) > 0
}

func (m Merge) getError() error {
	return m.err
}

func (m Merge) accept(visitor *CypherRenderer) {
	visitor.enter(m)
	m.pattern.accept(visitor)
	for _, event := range m.onCreateOrMatchEvents {
		event.accept(visitor)
	}
	visitor.leave(m)
}

func (m Merge) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("MERGE ")
}

func (m Merge) leave(renderer *CypherRenderer) {
	if !m.hasEvent() {
		renderer.append(" ")
	}
}

func (m Merge) getKey() string {
	return m.key
}

func (m Merge) isNotNil() bool {
	return m.notNil
}

func (m Merge) isUpdatingClause() bool {
	panic("implement me")
}
