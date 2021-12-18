package async

import (
	"database/sql"
	"github.com/saulocalixto/ms-cadastro-pet/infraestructure/database"
	"github.com/saulocalixto/ms-cadastro-pet/infraestructure/database/repository"
	"github.com/saulocalixto/ms-cadastro-pet/internal/dto"
	"github.com/saulocalixto/ms-cadastro-pet/internal/model"
	"github.com/saulocalixto/ms-cadastro-pet/internal/services"
	"log"
	"math"
	"sync"
)

var wg sync.WaitGroup

func Salvar(pets []dto.PetDTO) []dto.PetValidatedDTO {
	if pets == nil {
		log.Panic("Não foi informado nenhum pet para cadastro")
	}
	numeroDeRotinas := 5
	tamanhoLista := int(math.Ceil(float64 (len(pets) / numeroDeRotinas)))
	savePetChannels := obterSaveChannels(numeroDeRotinas)
	errorPetChannels := obterErrosChannels(numeroDeRotinas)
	listaParticionada := obterListaParticionada(pets, numeroDeRotinas, tamanhoLista)
	errors := make([]dto.PetValidatedDTO, 0)
	salvar(listaParticionada, savePetChannels, errorPetChannels, &errors)
	return errors
}

func salvar(listaParticionada [][]dto.PetDTO, savePetChannels map[int]chan dto.PetDTO, errorPetChannels map[int]chan dto.PetValidatedDTO, errors *[]dto.PetValidatedDTO) {
	connection := database.ObterConexao()
	defer connection.Close()
	for i, lista := range listaParticionada {
		go validarPets(lista, savePetChannels[i], errorPetChannels[i])
		wg.Add(2)
		go salvarPetsValidos(savePetChannels[i], connection)
		go obterErros(errorPetChannels[i], errors)
	}
	wg.Wait()
}

func obterListaParticionada(pets []dto.PetDTO, numeroDeRotinas int, tamanhoLista int) [][]dto.PetDTO {
	listaParticionada := make([][]dto.PetDTO, numeroDeRotinas)
	quantidadPetsAdicionados := 0
	for i := 0; i < numeroDeRotinas; i++ {
		for j := quantidadPetsAdicionados; j <= (tamanhoLista+quantidadPetsAdicionados)-1; j++ {
			if j < len(pets) {
				listaParticionada[i] = append(listaParticionada[i], pets[j])
			}
		}
		quantidadPetsAdicionados += len(listaParticionada[i])
	}
	return listaParticionada
}

func obterSaveChannels(numeroDeRotinas int) map[int]chan dto.PetDTO {
	savePetChannels := make(map[int]chan dto.PetDTO)
	for i := 0; i <= numeroDeRotinas; i++ {
		savePetChannels[i] = make(chan dto.PetDTO)
	}
	return savePetChannels
}

func obterErrosChannels(numeroDeRotinas int) map[int]chan dto.PetValidatedDTO {
	errorChannels := make(map[int]chan dto.PetValidatedDTO)
	for i := 0; i <= numeroDeRotinas; i++ {
		errorChannels[i] = make(chan dto.PetValidatedDTO)
	}
	return errorChannels
}

func validarPets(pets []dto.PetDTO, savePetChannel chan dto.PetDTO, errorPetChannel chan dto.PetValidatedDTO) {
	for _, pet := range pets {
		validarPet(pet, savePetChannel, errorPetChannel)
	}
	close(savePetChannel)
	close(errorPetChannel)
}

func obterErros(errorPetChannel chan dto.PetValidatedDTO, errors* []dto.PetValidatedDTO) {
	for err := range errorPetChannel {
		*errors = append(*errors, err)
	}
	wg.Done()
}

func salvarPetsValidos(savePetChannel chan dto.PetDTO, connection *sql.DB) {
	for pet := range savePetChannel {
		petModel := obterModel(pet)
		repository.Insert(petModel, connection)
	}
	wg.Done()
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

func validarPet(pet dto.PetDTO, savePetChannel chan dto.PetDTO, errorPetChannel chan dto.PetValidatedDTO) {
	result := services.Valide(pet)
	if result.IsValid() {
		savePetChannel <- result.Pet
	} else {
		errorPetChannel <- result
	}
}
