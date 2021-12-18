CREATE TABLE pet (
                     id SERIAL PRIMARY KEY,
                     nome VARCHAR(255),
                     dataDeNascimento Date,
                     peso float,
                     imunizado boolean,
                     raca VARCHAR(255),
                     especie VARCHAR(255),
                     proprietario_nome VARCHAR(255),
                     proprietario_dataDeNascimento Date,
                     proprietario_endereco VARCHAR(255),
                     proprietario_telefone VARCHAR(255),
                     veterinario_nome VARCHAR(255),
                     veterinario_endereco VARCHAR(255),
                     veterinario_telefone VARCHAR(255),
                     veterinario_crmv VARCHAR(255)
);