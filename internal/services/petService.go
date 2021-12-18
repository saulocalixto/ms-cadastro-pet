package services

import (
	"log"
	"ms-cadastro-pet/infraestructure/database/repository"
	"ms-cadastro-pet/internal/dto"
	"ms-cadastro-pet/internal/model"
	"sync"
)

var wg sync.WaitGroup

func SalvarAsync(pets []dto.PetDTO) []dto.PetValidatedDTO {
	if pets == nil {
		log.Panic("Não foi informado nenhum pet para cadastro")
	}
	savePetChannel := make(chan dto.PetDTO)
	errorPetChannel := make(chan dto.PetValidatedDTO)
	go validePets(pets, savePetChannel, errorPetChannel)
	wg.Add(2)
	go salvarPetsValidos(savePetChannel)
	errors := make([]dto.PetValidatedDTO, 0)
	go getErrors(errorPetChannel, &errors)
	wg.Wait()
	return errors
}

func SalvarSync(pets []dto.PetDTO) []dto.PetValidatedDTO {
	if pets == nil {
		log.Panic("Não foi informado nenhum pet para cadastro")
	}
	errors := make([]dto.PetValidatedDTO, 0)
	for _ ,pet := range pets {
		result := Valide(pet)
		if result.IsValid() {
			petModel := getModel(pet)
			repository.Insert(petModel)
		} else {
			errors = append(errors, result)
		}
	}
	return errors
}

func validePets(pets []dto.PetDTO, savePetChannel chan dto.PetDTO, errorPetChannel chan dto.PetValidatedDTO) {
	for _, pet := range pets {
		validaPet(pet, savePetChannel, errorPetChannel)
	}
	close(savePetChannel)
	close(errorPetChannel)
}

func getErrors(errorPetChannel chan dto.PetValidatedDTO, errors* []dto.PetValidatedDTO) {
	for r := range errorPetChannel {
		*errors = append(*errors, r)
	}
	wg.Done()
}

func salvarPetsValidos(savePetChannel chan dto.PetDTO) {
	for pet := range savePetChannel {
		petModel := getModel(pet)
		repository.Insert(petModel)
	}
	wg.Done()
}

func getModel(pet dto.PetDTO) model.Pet {
	petModel := model.Pet{
		Nome:             pet.Nome,
		Raca:             pet.Raca,
		Especie:          pet.Especie,
		Peso:             pet.Peso,
		DataDeNascimento: pet.DataDeNascimento,
		Imunizado:        len(pet.Vacinas) >= 4}

	if (dto.Proprietario{} != pet.Proprietario) {
		petModel.ProprietarioNome = pet.Proprietario.Nome
		petModel.ProprietarioTelefone = pet.Proprietario.Telefone
		petModel.ProprietarioEndereco = pet.Proprietario.Endereco
		petModel.ProprietarioDataDeNascimento = pet.Proprietario.DataDeNascimento
	}

	if (dto.Veterinario{} != pet.Veterinario) {
		petModel.VeterinarioNome = pet.Veterinario.Nome
		petModel.VeterinarioTelefone = pet.Veterinario.Telefone
		petModel.VeterinarioEndereco = pet.Veterinario.Endereco
		petModel.VeterinarioCrm = pet.Veterinario.CRMV
	}
	return petModel
}

func validaPet(pet dto.PetDTO, savePetChannel chan dto.PetDTO, errorPetChannel chan dto.PetValidatedDTO) {
	result := Valide(pet)
	if result.IsValid() {
		savePetChannel <- result.Pet
	} else {
		errorPetChannel <- result
	}
}
