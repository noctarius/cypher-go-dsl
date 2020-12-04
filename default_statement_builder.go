package cypher_go_dsl

type DefaultStatementBuilder struct {
	invalidReason             string
	currentSinglePartElements []Visitable
	multipartElements         []MultiPartElement
	currentOngoingMatch       MatchBuilder
	currentOngoingUpdate      DefaultStatementWithUpdateBuilder
	currentOngoingCall        DefaultStatementWithUpdateBuilder
}

func (d DefaultStatementBuilder) where(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementBuilder) addWith(with With) {
	if with.isNotNil() {
		d.multipartElements = append(d.multipartElements, MultiPartElementCreate(d.BuildListOfVisitable(true), with))
	}
}

func (d DefaultStatementBuilder) wherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDefault(distinct bool, expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) deleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) delete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) merge(pattern ...PatternElement) {
	panic("implement me")
}

func (d DefaultStatementBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwinds(expression ...Expression) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwindByString(variable string) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwind(expression Expression) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) call(statement Statement) OngoingReadingWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) callExposes(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	panic("implement me")
}

func (d DefaultStatementBuilder) asCondition() Expression {
	panic("implement me")
}

func (d DefaultStatementBuilder) and(condition Condition) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) andPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) or(condition Condition) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) orPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) OptionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) create(element ...PatternElement) {
	panic("implement me")
}

func NewDefaultBuilder() DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentSinglePartElements: make([]Visitable, 0),
	}
}

func (d DefaultStatementBuilder) Match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	if pattern == nil || len(pattern) == 0 {
		return DefaultStatementBuilder{invalidReason: "patterns to match is required"}
	}
	if d.currentOngoingMatch.notNil {
		d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingMatch.buildMatch())
	}
	d.currentOngoingMatch = MatchBuilder{
		optional: false,
		notNil:   true,
	}
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, pattern...)
	return d
}

func (d DefaultStatementBuilder) closeCurrentOngoingMatch() {

}

func (d DefaultStatementBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(false, expression...)
}

func (d DefaultStatementBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(true, expression...)
}

func (d DefaultStatementBuilder) returningDefault(distinct bool, expression ...Expression) OngoingReadingAndReturn {
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct:       distinct,
		defaultBuilder: d,
	}
	withReturnBuilder.AddExpression(expression...)
	return withReturnBuilder
}

func (d DefaultStatementBuilder) Build() Statement {
	return d.BuildImpl(false, Return{})
}

func (d DefaultStatementBuilder) BuildImpl(clearCurrentBuildSteps bool, returning Return) Statement {
	singlePartQuery, _ := SinglePartQueryCreate(d.BuildListOfVisitable(clearCurrentBuildSteps), returning)
	if len(d.multipartElements) == 0 {
		return singlePartQuery
	}
	return MultiPartQueryCreate(d.multipartElements, singlePartQuery)
}

func (d *DefaultStatementBuilder) BuildListOfVisitable(clearAfter bool) []Visitable {
	visitables := make([]Visitable, 0)
	copy(visitables, d.currentSinglePartElements)
	if d.currentOngoingMatch.notNil {
		visitables = append(visitables, d.currentOngoingMatch.buildMatch())
	}
	if d.currentOngoingUpdate.isNotNil() {
		visitables = append(visitables, d.currentOngoingUpdate.builder.build())
	}
	if d.cu {

	}
	d.currentSinglePartElements = d.currentSinglePartElements[:0]
	return visitables
}
