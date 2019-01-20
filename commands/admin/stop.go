package admin

import (
	"io"
	"time"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// P references the var Pipe in main
var P *io.WriteCloser

// StopServer saves and stops the server.
func StopServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	io.WriteString(*P, "/save\n")
	io.WriteString(*P, "/quit\n")
	time.Sleep(10 * time.Second)
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Server saved and shutting down; Cya!")
	s.Close()
	support.Exit(0)
}
