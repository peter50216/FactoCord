package admin

import (
	"fmt"
	"io"
	"strings"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// LuaCommand executes a LUA command on server
func LuaCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(R == nil)
	if *R == false {
		s.ChannelMessageSend(support.Config.FactorioConsoleChannelID, "Server is not running!")
		return
	}
	TmpList := strings.Split(m.Content, " ")
	io.WriteString(*P, "/c "+strings.Join(TmpList[1:], " ")+"\n")
	return
}
