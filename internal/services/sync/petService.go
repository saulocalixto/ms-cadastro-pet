package sync

import (
	"github.com/saulocalixto/ms-cadastro-pet/internal/dto"
	"github.com/saulocalixto/ms-cadastro-pet/internal/model"
	"github.com/saulocalixto/ms-cadastro-pet/internal/services"
	"log"
)

import (
	"github.com/saulocalixto/ms-cadastro-pet/infraestructure/database"
	"github.com/saulocalixto/ms-cadastro-pet/infraestructure/database/repository"
)

func Salvar(pets []dto.PetDTO) []dto.PetValidatedDTO {
	if pets == nil {
		log.Panic("Não foi informado nenhum pet para cadastro")
	}
	errors := make([]dto.PetValidatedDTO, 0)
	connection := database.ObterConexao()
	defer connection.Close()
	for _ ,pet := range pets {
		result := services.Valide(pet)
		if result.IsValid() {
			petModel := obterModel(pet)
			repository.Insert(petModel, connection)
		} else {
			errors = append(errors, result)
		}
	}
	return errors
}

func obterModel(pet dto.PetDTO) model.Pet {
	petModel := model.Pet{
		Nome:             pet.Nome,
		Raca:             pet.Raca,
		Especie:          pet.Especie,
		Peso:             pet.Peso,
		DataDeNascimento: pet.DataDeNascimento,
		Imunizado:        len(pet.Vacinas) >= 4}

	if (dto.ProprietarioDTO{} != pet.Proprietario) {
		petModel.ProprietarioNome = pet.Proprietario.Nome
		petModel.ProprietarioTelefone = pet.Proprietario.Telefone
		petModel.ProprietarioEndereco = pet.Proprietario.Endereco
		petModel.ProprietarioDataDeNascimento = pet.Proprietario.DataDeNascimento
	}

	if (dto.VeterinarioDTO{} != pet.Veterinario) {
		petModel.VeterinarioNome = pet.Veterinario.Nome
		petModel.VeterinarioTelefone = pet.Veterinario.Telefone
		petModel.VeterinarioEndereco = pet.Veterinario.Endereco
		petModel.VeterinarioCrm = pet.Veterinario.CRMV
	}
	return petModel
}