package model

import "database/sql"

//Local armazena os dados da localização pelo
type Local struct {
	Pais             string         `json:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode"`
}
