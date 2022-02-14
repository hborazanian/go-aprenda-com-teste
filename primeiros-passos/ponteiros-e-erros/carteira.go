package carteira

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
