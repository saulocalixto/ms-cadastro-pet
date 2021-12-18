package dto

type PetDTO struct {
	Nome             string       `json:"nome"`
	DataDeNascimento string       `json:"dataDeNascimento"`
	Peso         float32         `json:"peso"`
	Proprietario ProprietarioDTO `json:"proprietario"`
	Veterinario  VeterinarioDTO `json:"veterinario"`
	Vacinas      []VacinaDTO    `json:"vacinas"`
	Raca             string     `json:"raca"`
	Especie          string       `json:"especie"`
}

type ProprietarioDTO struct {
	Nome             string `json:"nome"`
	DataDeNascimento string `json:"dataDeNascimento"`
	Endereco         string `json:"endereco"`
	Telefone         string `json:"telefone"`
}

type VeterinarioDTO struct {
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Telefone string `json:"telefone"`
	CRMV     string `json:"crmv"`
}

type VacinaDTO struct {
	Marca           string `json:"marca"`
	DataDeAplicacao string `json:"dataAplicacao"`
}
