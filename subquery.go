package cypher_go_dsl

import (
	"errors"
)

type Subquery struct {
	statement Statement
	key       string
	notNil    bool
	err       error
}

func SubqueryCreate(statement Statement) Subquery {
	subQuery := Subquery{
		statement: statement,
		notNil:    true,
	}
	subQuery.key = getAddress(&subQuery)
	return subQuery
}

func SubqueryCall(statement Statement) (Subquery, error) {
	validReturn := false
	if singlePartQuery, isSinglePartQuery := statement.(SinglePartQuery); isSinglePartQuery {
		validReturn = singlePartQuery.doesReturnElements()
	} else if multiPartQuery, isMultiPartQuery := statement.(MultiPartQuery); isMultiPartQuery {
		validReturn = multiPartQuery.doesReturnElements()
	} else if procedureCall, isProcedureCall := statement.(ProcedureCall); isProcedureCall {
		validReturn = procedureCall.doesReturnElements()
	} else if _, isUnionQuery := statement.(UnionQuery); isUnionQuery {
		validReturn = true
	}
	if !validReturn {
		return Subquery{}, errors.New("only a statement that returns elements, either via return or yield, can be used in a subquery")
	}
	return SubqueryCreate(statement), nil
}

func (s Subquery) getError() error {
	return s.err
}

func (s Subquery) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	s.statement.accept(visitor)
	visitor.leave(s)
}

func (s Subquery) enter(renderer *CypherRenderer) {
	renderer.append("CALL {")
}

func (s Subquery) leave(renderer *CypherRenderer) {
	renderer.append("} ")
}

func (s Subquery) getKey() string {
	return s.key
}

func (s Subquery) isNotNil() bool {
	return s.notNil
}
