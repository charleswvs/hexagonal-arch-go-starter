package app_test // para testar metodos internos (privados), nao precisa utilizar o "_test", mas caso sejam metodos externos, se usa o "_test"

import (
	"testing"

	"github.com/charles00willian/hexagonal-arch-go-starter/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal than 0", err.Error())
}
