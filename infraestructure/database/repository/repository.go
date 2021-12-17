package repository

import (
	"ms-cadastro-pet/infraestructure/database"
	"ms-cadastro-pet/internal/model"
)

func Insert(pet model.Pet) {
	sqlStatement := getInsertScript()
	database.GetConnection().QueryRow(sqlStatement,
		pet.Nome,
		pet.DataDeNascimento,
		pet.Peso,
		pet.Imunizado,
		pet.Raca,
		pet.Especie,
		pet.ProprietarioNome,
		pet.ProprietarioDataDeNascimento,
		pet.ProprietarioEndereco,
		pet.ProprietarioTelefone,
		pet.VeterinarioNome,
		pet.VeterinarioEndereco,
		pet.VeterinarioTelefone,
		pet.VeterinarioCrm)
	database.CloseConnection()
}

func getInsertScript() string {
	return `INSERT INTO pet
    	(nome,
    	 dataDeNascimento,
    	 peso,
    	 imunizado,
    	 raca,
    	 especie,
    	 proprietario_nome,
    	 proprietario_dataDeNascimento,
    	 proprietario_endereco,
    	 proprietario_telefone,
    	 veterinario_nome,
    	 veterinario_endereco,
    	 veterinario_telefone,
    	 veterinario_crmv) 
    	 VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
}