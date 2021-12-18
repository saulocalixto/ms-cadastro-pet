package dto

type ValidationDTO struct{
	Campo string `json:"campo"`
	Mensagem string `json:"mensagem"`
}

type PetValidatedDTO struct{
	Validations []ValidationDTO `json:"validations"`
	Pet PetDTO `json:"pet"`
}

type ReturnDTO struct{
	Tempo string `json:"tempo"`
	Total int `json:"registros_processados"`
	Validations []PetValidatedDTO `json:"validations"`
}

func (dto PetValidatedDTO) IsValid() bool {
	return len(dto.Validations) == 0
}