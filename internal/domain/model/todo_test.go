package model

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTodo_NewTodo(t *testing.T) {
	// given
	name := "Play piano"

	// when
	todo, err := NewTodo(name)

	// then
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.Equal(t, name, todo.Task())
	require.False(t, todo.Complete())
}

func TestTodo_NewTodo_EmptyName(t *testing.T) {
	// given
	name := ""

	// when
	todo, err := NewTodo(name)

	// then
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrInvalidTaskName))
	require.Nil(t, todo)
}

func TestTodo_Of(t *testing.T) {
	// given
	id := 1
	name := "Play piano"
	complete := true

	// when
	todo := Of(id, name, complete)

	// then
	require.NotNil(t, todo)
	require.Equal(t, id, todo.Id())
	require.Equal(t, name, todo.Task())
	require.True(t, todo.Complete())
}

func TestTodo_Check(t *testing.T) {
	// given
	todo := Todo{}

	// when
	todo.Check()

	// then
	require.True(t, todo.Complete())
}

func TestTodo_Uncheck(t *testing.T) {
	// given
	todo := Todo{complete: true}

	// when
	todo.Uncheck()

	// then
	require.False(t, todo.Complete())
}
