package book

import (
	"strings"
	"testing"
)

// ***************
// First approach
// ***************

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_Execute(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T){
		s := &TestStruct{}
		res := s.Execute(s)
		expectedOrError(res, " world ", t)
	})
	t.Run("Using anonymous functions", func(t *testing.T){

		m := new(AnonymousTemplate)
		res := m.Execute(func() string {
			return "world"
		})
		expectedOrError(res, " world ", t)
	})
	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T){
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}

		template := Template{}
		res := template.Execute(messageRetriever)

		expectedOrError(res, " world ", t)
	})

}


func expectedOrError(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("Expected string '%s' was not found on returned string\n", expected)
	}
}