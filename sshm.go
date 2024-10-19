package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	readConfig()

	itemsFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 113: // q
			app.Stop()
		case 102: //f
			index := itemsList.GetCurrentItem()
			executeShell(&data.Items[index], Transfer)
		case 97: // a
			addForm.Clear(true)
			addConnectionForm()
			pages.SwitchToPage(addConnectionsPage)
		case 101: // e
			editConnection()
			pages.SwitchToPage(editConnectionsPage)
		case 100: // d
			pages.SwitchToPage(deleteConnectionPage)
		}

		return event
	})

	addFormFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 126 { // ~
			pages.SwitchToPage(savedConnectionsPage)
		}
		return event
	})

	editFormFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 126 { // ~
			pages.SwitchToPage(savedConnectionsPage)
		}
		return event
	})

	fillItemsListData()

	itemsFlex.SetDirection(tview.FlexRow).
		AddItem(itemsList, 0, 1, true).
		AddItem(mainMenu, 0, 1, false)

	addFormFlex.SetDirection(tview.FlexRow).
		AddItem(addForm, 0, 1, true).
		AddItem(formMenu, 0, 1, false)

	editFormFlex.SetDirection(tview.FlexRow).
		AddItem(editForm, 0, 1, true).
		AddItem(formMenu, 0, 1, false)

	deleteFlex.SetDirection(tview.FlexRow).
		AddItem(modal, 0, 1, true)

	pages.AddPage(savedConnectionsPage, itemsFlex, true, true)
	pages.AddPage(addConnectionsPage, addFormFlex, true, false)
	pages.AddPage(editConnectionsPage, editFormFlex, true, false)
	pages.AddPage(deleteConnectionPage, deleteFlex, true, false)

	itemsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		executeShell(&data.Items[index], Shell)
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
