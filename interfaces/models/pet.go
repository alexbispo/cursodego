package models

//Gato representa um gato
type Gato interface {
	Mia() string
}

//Cachorro representa um cachorro
type Cachorro interface {
	Late() string
}

//Pet representa um animal de estimação
type Pet struct {
	Nome string
}

//Mia faz miau
func (p Pet) Mia() string {
	return p.Nome + " Miaaau, miaauuu!...."
}

//Late faz Au au Au!
func (p Pet) Late() string {
	return p.Nome + " Au, au, au Au..."
}
