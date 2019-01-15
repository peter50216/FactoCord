package admin

import (
	"fmt"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// Executes a LUA command on server
func LuaCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(R == nil)
	if *R == false {
		s.ChannelMessageSend(support.Config.FactorioChannelID, "Server is not running!")
		return
	}
	//io.WriteString(*P, "/save\n")
	//io.WriteString(*P, "/quit\n")
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Lua command would run now...")
	//time.Sleep(3 * time.Second)
	//*R = false
	//RestartCount = RestartCount + 1
	return
}
