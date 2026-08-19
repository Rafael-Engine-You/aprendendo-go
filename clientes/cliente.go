package clientes

import  t "time"

func NomeDoMetodo() {}

type Cliente struct {
	NomeCompleto string
	DataNascimento t.Time 
	Cpf string
	Contatos []string
}