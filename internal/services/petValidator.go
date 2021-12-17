package services

import (
	"fmt"
	"ms-cadastro-pet/internal/dto"
	"time"
)

func Valide(pet dto.PetDTO) dto.PetValidatedDTO {
	validations := make([]dto.ValidationDTO, 0)
	validations = valideNomePet(pet, validations)
	validations = valideDataDeNascimento(pet.DataDeNascimento, "dataDeNascimento", validations)
	validations = valideCampoNaoPodeSerNulo(pet.Raca, "raca", validations)
	validations = valideCampoNaoPodeSerNulo(pet.Especie, "especie", validations)
	validations = valideProprietario(pet.Proprietario, validations)
	validations = valideVeterinario(pet.Veterinario, validations)
	return dto.PetValidatedDTO{Pet: pet, Validations: validations}
}

func valideNomePet(pet dto.PetDTO, validations []dto.ValidationDTO) []dto.ValidationDTO {
	validations = valideCampoNaoPodeSerNulo(pet.Nome, "nome", validations)
	return validations
}

func valideProprietario(proprietario dto.Proprietario, validations []dto.ValidationDTO) []dto.ValidationDTO {
	if (dto.Proprietario{} == proprietario) {
		validations = append(validations, dto.ValidationDTO{
			Campo: "proprietario",
			Mensagem: "O properitário deve ser preenchido"})
	}
	validations = valideCampoNaoPodeSerNulo(proprietario.Nome, "proprietario.nome", validations)
	validations = valideDataDeNascimento(
		proprietario.DataDeNascimento,
		"proprietario.dataDeNascimento",
		validations)
	validations = valideCampoNaoPodeSerNulo(proprietario.Endereco, "proprietario.endereco", validations)
	return validations
}

func valideVeterinario(veterinario dto.Veterinario, validations []dto.ValidationDTO) []dto.ValidationDTO {
	if (dto.Veterinario{} != veterinario) {
		validations = valideCampoNaoPodeSerNulo(veterinario.Nome, "veterinario.nome", validations)
		validations = valideCampoNaoPodeSerNulo(veterinario.Endereco, "veterinario.endereco", validations)
		validations = valideCampoNaoPodeSerNulo(veterinario.CRMV, "veterinario.crmv", validations)
	}
	return validations
}

func valideDataDeNascimento(dataTexto string, nomeCampo string, validations []dto.ValidationDTO) []dto.ValidationDTO {
	if dataTexto == "" {
		validations = append(validations, dto.ValidationDTO{
			Campo:    nomeCampo,
			Mensagem: "A data deve ser preenchida"})
		return validations
	}
	data, err := time.Parse("2006-01-02", dataTexto)
	if err != nil {
		validations = append(validations, dto.ValidationDTO{
			Campo:    nomeCampo,
			Mensagem: "Erro ao converter a data, observe se ela foi passada no formato: yyyy-MM-dd"})
	}
	if data.After(time.Now()) {
		validations = append(validations, dto.ValidationDTO{
			Campo:    nomeCampo,
			Mensagem: "A data deve ser anterior a data atual"})
	}
	return validations
}

func valideCampoNaoPodeSerNulo(valorCampo string, nomeCampo string, validations []dto.ValidationDTO) []dto.ValidationDTO {
	if valorCampo == "" {
		validations = append(validations, dto.ValidationDTO{
			Campo: nomeCampo,
			Mensagem: fmt.Sprintf("O %s deve ser preenchido", nomeCampo)})
	}
	return validations
}
