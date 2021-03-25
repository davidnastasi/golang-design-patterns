package book

import "strings"

type MessageRetriever interface {
  Message() string
}

type Templater interface {
  first() string
  third() string
  Execute(MessageRetriever) string
}

type Template struct {
}

func (t Template) first() string {
	return "hello"
}

func (t Template) third() string {
	return "template"
}

func (t Template) Execute(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

// ******************************
// Second approach. Using function
// ******************************

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) Execute(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

// ******************************
// Third approach. Adapter pattern
// ******************************

type templateAdapter struct {
	myFunc func() string
}

func (a *templateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &templateAdapter{myFunc:f}
}
