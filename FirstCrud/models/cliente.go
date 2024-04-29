package models

import (
	"example.com/crud/db"
	"example.com/crud/utils"
)

type Cliente struct {
	CliId       int64
	CliNome     string
	CliEmail    string
	CliPassword string
	CliTelefone string
}

func (c Cliente) New() error {
	query := "insert into tblCLI_Cliente(cliNome, cliEmail, cliPassword, cliTelefone) values (?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashed, err := utils.HashPassword(c.CliPassword)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.CliNome, c.CliEmail, hashed, c.CliTelefone)

	if err != nil {
		return err
	}

	return nil
}

func List() ([]Cliente, error) {
	var clientes []Cliente
	query := "select cliID, cliNome, cliEmail, cliTelefone from tblCLI_Cliente"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cliente Cliente
		if err := rows.Scan(&cliente.CliId, &cliente.CliNome, &cliente.CliEmail, &cliente.CliTelefone); err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}

func Read(CliID int64) (Cliente, error) {
	query := "select cliID, cliNome, cliEmail, cliTelefone from tblCLI_Cliente where cliID = ?"

	row := db.DB.QueryRow(query, CliID)

	var cliente Cliente

	if err := row.Scan(&cliente.CliId, &cliente.CliNome, &cliente.CliEmail, &cliente.CliTelefone); err != nil {
		return Cliente{}, err
	}

	return cliente, nil
}

func (c Cliente) Update() error {
	query := `update tblCLI_Cliente set cliNome = ?, cliEmail = ?, cliTelefone = ? where cliID = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.CliNome, c.CliEmail, c.CliTelefone, c.CliId)

	if err != nil {
		return err
	}

	return nil
}

func (c Cliente) Delete() error {
	query := "delete from tblCLI_Cliente where cliID = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&c.CliId)

	if err != nil {
		return err
	}

	return nil
}
