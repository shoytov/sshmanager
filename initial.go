package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Action int

// Actions for servers
const (
	Shell Action = iota
	Transfer
)

var app = tview.NewApplication()

var mainMenu = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(q) to quit\n(f) to connect with file-transfer mode\n(a) add connection\n(e) edit connection\n(d) delete connection")

var itemsList = tview.NewList().ShowSecondaryText(false)

var pages = tview.NewPages()

var itemsFlex = tview.NewFlex()

var modal = tview.NewModal().
	SetText("Are you sure you want to delete this connection?").
	AddButtons([]string{"Yes", "Cancel"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			deleteConnection()
		} else {
			pages.SwitchToPage(savedConnectionsPage)
		}
	})

var addFormFlex = tview.NewFlex()
var editFormFlex = tview.NewFlex()
var deleteFlex = tview.NewFlex()

var addForm = tview.NewForm()
var editForm = tview.NewForm()

var formMenu = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(~) to go back")

type Server struct {
	Name             string
	ConnectionString string
	ConnectionPort   string
	Description      string
}

type Data struct {
	Items []Server
}

var data Data

var configFilePath = filepath.Join(os.Getenv("HOME"), ".config", "sshmanager", "config.json")

func readConfig() {
	dir := filepath.Dir(configFilePath)

	// make dirs if not exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	configFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	// json parse
	bytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	if len(bytes) == 0 {
		content := "{}"
		if _, err := configFile.WriteString(content); err != nil {
			panic(err)
		}
		bytes = []byte(content)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}
}

func saveConfig() {
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	// save updated json into file
	err = os.WriteFile(configFilePath, updatedJSON, 0644)
	if err != nil {
		panic(err)
	}
}

func fillItemsListData() {
	for index, item := range data.Items {
		itemsList.AddItem(item.Name+" ("+item.Description+")", "", rune(49+index), nil)
	}
}
