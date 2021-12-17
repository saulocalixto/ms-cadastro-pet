package dto

type PetDTO struct {
	Nome             string       `json:"nome"`
	DataDeNascimento string       `json:"dataDeNascimento"`
	Peso             float32      `json:"peso"`
	Proprietario     Proprietario `json:"proprietario"`
	Veterinario      Veterinario  `json:"veterinario"`
	Vacinas          []Vacina     `json:"vacinas"`
	Raca             string       `json:"raca"`
	Especie          string       `json:"especie"`
}

type Proprietario struct {
	Nome             string `json:"nome"`
	DataDeNascimento string `json:"dataDeNascimento"`
	Endereco         string `json:"endereco"`
	Telefone         string `json:"telefone"`
}

type Veterinario struct {
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Telefone string `json:"telefone"`
	CRMV     string `json:"crmv"`
}

type Vacina struct {
	Marca           string `json:"marca"`
	DataDeAplicacao string `json:"dataAplicacao"`
}
