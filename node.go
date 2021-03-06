package cypher_go_dsl

import (
	"fmt"
)

type Node struct {
	symbolicName SymbolicName
	labels       []NodeLabel
	properties   Properties
	key          string
	notNil       bool
	err          error
}

func NodeCreate() Node {
	node := Node{
		notNil: true,
	}
	return node
}

func (node *Node) injectKey() {
	node.key = fmt.Sprint(fmt.Sprintf("%p", node))
}

func NodeCreate1(primaryLabel string, properties Properties, additionalLabels ...string) Node {
	labels := make([]NodeLabel, 0)
	if primaryLabel != "" {
		labels = append(labels, NodeLabel{value: primaryLabel})
	}
	for _, label := range additionalLabels {
		labels = append(labels, NodeLabelCreate(label))
	}
	node := Node{
		notNil:       true,
		symbolicName: SymbolicName{},
		properties:   properties,
		labels:       labels,
	}
	node.injectKey()
	return node
}

func NodeCreate2(primaryLabel string) Node {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabelCreate(primaryLabel))
	node := Node{
		labels: labels,
		notNil: true,
	}
	node.injectKey()
	return node
}

func NodeCreate3(primaryLabel string, additionalLabel ...string) Node {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	for _, label := range additionalLabel {
		labels = append(labels, NodeLabelCreate(label))
	}
	node := Node{
		labels: labels,
	}
	node.injectKey()
	return node
}

func NodeCreate4(newProperties MapExpression, node Node) Node {
	newNode := Node{symbolicName: node.symbolicName, labels: node.labels, notNil: true, properties: PropertiesCreate(newProperties)}
	node.injectKey()
	return newNode
}

func (node Node) getSymbolicName() SymbolicName {
	return node.symbolicName
}

func (node Node) isNotNil() bool {
	return node.notNil
}

func (node Node) IsPatternElement() bool {
	return true
}

func (node Node) getKey() string {
	return node.key
}

func (node Node) hasSymbolic() bool {
	return node.symbolicName.isNotNil()
}

func (node Node) getError() error {
	return node.err
}

func (node Node) accept(visitor *CypherRenderer) {
	(*visitor).enter(node)
	VisitIfNotNull(node.symbolicName, visitor)
	for _, label := range node.labels {
		label.accept(visitor)
	}
	VisitIfNotNull(node.properties, visitor)
	(*visitor).leave(node)
}

func (node Node) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("(")
	if !node.hasSymbolic() {
		return
	}
	_, renderer.skipNodeContent = renderer.visitedNamed[node.key]
	renderer.visitedNamed[node.key] = 1
	if renderer.skipNodeContent {
		renderer.builder.WriteString(node.symbolicName.value)
	}
}

func (node Node) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(")")
	renderer.skipNodeContent = false
}

func (node Node) RelationshipTo(other Node, types ...string) Relationship {
	return RelationshipCreate(node, LTR(), other, types...)
}

func (node Node) RelationshipFrom(other Node, types ...string) Relationship {
	return RelationshipCreate(node, RTL(), other, types...)
}

func (node Node) RelationshipBetween(nodeDest Node, types ...string) Relationship {
	panic("implement me")
}

func (node Node) WithRawProperties(keysAndValues ...interface{}) (Node, error) {
	var properties = &MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		var err error
		*properties, err = NewMapExpression(keysAndValues...)
		if err != nil {
			return Node{}, err
		}
	}
	return node.WithProperties(*properties), nil
}

func (node Node) WithProperties(newProperties MapExpression) Node {
	return NodeCreate4(newProperties, node)
}

func (node Node) Property(name string) {
	fmt.Print(name)
}

func (node Node) Named(name string) Node {
	node.symbolicName = SymbolicNameCreate(name)
	return node
}

type NodeLabel struct {
	value  string
	key    string
	notNil bool
	err    error
}

func NodeLabelCreate(value string) NodeLabel {
	n := NodeLabel{
		value:  value,
		notNil: true,
	}
	n.key = getAddress(&n)
	return n
}

func (n NodeLabel) getError() error {
	return n.err
}

func (n NodeLabel) isNotNil() bool {
	return n.notNil
}

func (n NodeLabel) getKey() string {
	return n.key
}

func (n NodeLabel) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	visitor.leave(n)
}

func (n NodeLabel) enter(renderer *CypherRenderer) {
	if n.value == "" {
		return
	}
	renderer.builder.WriteString(NodeLabelStart)
	renderer.builder.WriteString(escapeName(n.value))
}

func (n NodeLabel) leave(renderer *CypherRenderer) {
}
