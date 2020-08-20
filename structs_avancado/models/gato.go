package models

// Gato reprsenta um gatinho
type Gato struct {
	Nome            string `json:"nome"`
	Idade           int    `json:"idade"`
	numeroDeVacinas int
}

// AdicionaVacina incrementa o número de vacinas
func (g *Gato) AdicionaVacina() {
	g.numeroDeVacinas++
}

// NumeroDeVacinas retorna o numero de vacinas
func (g *Gato) NumeroDeVacinas() int {
	return g.numeroDeVacinas
}

//Cachorro representa um cachorro
type Cachorro struct {
	Nome  string
	idade int
}

//NovoCachorro nova instancia de cachorro.
func NovoCachorro(nome string) Cachorro {
	return Cachorro{Nome: nome}
}

//SetIdade define a idade do cachorro
func (c *Cachorro) SetIdade(idade int) {
	c.idade = idade
}

//GetIdade obtem a idade do cachorro
func (c *Cachorro) GetIdade() int {
	return c.idade
}
