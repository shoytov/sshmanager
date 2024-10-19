package main

import "github.com/rivo/tview"

func addConnectionForm() *tview.Form {
	server := Server{}

	addForm.AddInputField(connectionName, "", 20, nil, func(name string) {
		server.Name = name
	})

	addForm.AddInputField(connectionString, "", 200, nil, func(connectionString string) {
		server.ConnectionString = connectionString
	})

	addForm.AddInputField(connectionPort, "22", 200, nil, func(connectionPort string) {
		server.ConnectionPort = connectionPort
	})

	addForm.AddInputField(connectionDescription, "", 200, nil, func(description string) {
		server.Description = description
	})

	addForm.AddButton(connectionFormSaveButton, func() {
		addServer(server)
	})

	return addForm
}

func editConnectionForm(server Server, index int) *tview.Form {
	editForm.AddInputField(connectionName, server.Name, 20, nil, func(name string) {
		server.Name = name
	})

	editForm.AddInputField(connectionString, server.ConnectionString, 200, nil, func(connectionString string) {
		server.ConnectionString = connectionString
	})

	editForm.AddInputField(connectionPort, server.ConnectionPort, 200, nil, func(connectionPort string) {
		server.ConnectionPort = connectionPort
	})

	editForm.AddInputField(connectionDescription, server.Description, 200, nil, func(description string) {
		server.Description = description
	})

	editForm.AddButton(connectionFormSaveButton, func() {
		editServer(server, index)
	})

	return editForm
}
