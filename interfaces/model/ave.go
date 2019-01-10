package model

//Galinha representa uma ave do tipo Galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja representa o som feito por uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

//Quack representa o som feito for uma pata
func (a Ave) Quack() string {
	return "Quack, quack..."
}
