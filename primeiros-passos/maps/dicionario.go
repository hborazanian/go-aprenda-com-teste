package main

const (
	ErrNaoEncontrado      = ErroDicionario("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente   = ErroDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErroDicionario("não foi possível atualizar a palavra pois ela não existe")
)

type ErroDicionario string

func (e ErroDicionario) Error() string {
	return string(e)
}

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, erro := d.Busca(palavra)

	switch erro {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return erro
	}

	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, erro := d.Busca(palavra)

	switch erro {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return erro
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}
