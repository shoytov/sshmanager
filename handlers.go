package main

import (
	"fmt"
	"github.com/creack/pty"
	"golang.org/x/term"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func executeShell(server *Server, action Action) {
	var c *exec.Cmd

	// Create arbitrary command.
	if action == Shell {
		command := fmt.Sprintf("reset && ssh %s -p %s", server.ConnectionString, server.ConnectionPort)
		c = exec.Command("bash", "-c", command)
	} else if action == Transfer {
		c = exec.Command("bash", "-c", "mc sh://"+server.ConnectionString+":"+server.ConnectionPort)
	} else {
		os.Exit(1)
	}

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				panic(err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)
}

func addServer(newServer Server) {
	// save add form button handler
	data.Items = append(data.Items, newServer)
	itemsCount := itemsList.GetItemCount()
	itemsList.AddItem(newServer.Name+" ("+newServer.Description+")", "", rune(49+itemsCount), nil)

	saveConfig()

	pages.SwitchToPage(savedConnectionsPage)
}

func editServer(server Server, index int) {
	// save edit form button handler
	data.Items[index].Name = server.Name
	data.Items[index].Description = server.Description
	data.Items[index].ConnectionString = server.ConnectionString
	data.Items[index].ConnectionPort = server.ConnectionPort

	saveConfig()

	itemsList.Clear()
	fillItemsListData()
	pages.SwitchToPage(savedConnectionsPage)
}

func editConnection() {
	editForm.Clear(true)
	index := itemsList.GetCurrentItem()
	editConnectionForm(data.Items[index], index)
}

func deleteConnection() {
	index := itemsList.GetCurrentItem()
	data.Items = append(data.Items[:index], data.Items[index+1:]...)
	itemsList.Clear()
	fillItemsListData()
	saveConfig()
	pages.SwitchToPage(savedConnectionsPage)
}
