package models

import "example.com/crud/db"

type Cliente struct {
	CliId       int64
	CliNome     string
	CliEmail    string
	CliTelefone string
}

func (c Cliente) New() error {
	query := "insert into tblCLI_Cliente(cliNome, cliEmail, cliTelefone) values (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.CliNome, c.CliEmail, c.CliTelefone)

	if err != nil {
		return err
	}

	return nil
}

func List() ([]Cliente, error) {
	var clientes []Cliente
	query := "select * from tblCLI_Cliente"
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
