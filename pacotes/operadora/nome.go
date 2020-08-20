package operadora

import (
	"cursodego/pacotes/prefixo"
	"strconv"
)

// NomeOperadora representa o nome da operadora.
var NomeOperadora = "XPTO Telecom"

// PrefixoMaisOperadora representa o prefixo mais a operadora
var PrefixoMaisOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
