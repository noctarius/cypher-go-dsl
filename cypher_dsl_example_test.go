package cypher_go_dsl

import (
	"testing"
)

func TestFindAllMovies(t *testing.T) {
	movie := NewNode("Movie").Named("m")
	statement := Matchs(movie).
		returningByNamed(movie).
		build()
	query := NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}

func TestDefaultStatementBuilder_OptionalMatch(t *testing.T) {
	farm := NewNode("Farm").Named("b")
	statement := Matchs(farm).
		where(ConditionsNot(farm.RelationshipFrom(AnyNode(), "HAS"))).
		withByString("b").
		optionalMatch(farm.RelationshipTo(AnyNode1("p"), "HAS")).
		returningByString("b", "p").
		build()
	query := NewRenderer().Render(statement)
	expect := "MATCH (m:`Movie`) RETURN m"
	if query != expect {
		t.Errorf("%s is incorrect \n %s", query, expect)
	}
}
