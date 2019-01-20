package support

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
)

// Chat pipes in-game chat to Discord.
func Chat(s *discordgo.Session, m *discordgo.MessageCreate) {
	for {
		t, err := tail.TailFile("factorio.log", tail.Config{Follow: true})
		if err != nil {
			ErrorLog(fmt.Errorf("%s: An error occurred when attempting to tail factorio.log\nDetails: %s", time.Now(), err))
		}
		for line := range t.Lines {
			if strings.Contains(line.Text, "[CHAT]") {
				if !strings.Contains(line.Text, "<server>: [Discord]") {
					TmpList := strings.Split(line.Text, " ")
					TmpList[3] = strings.Replace(TmpList[3], ":", "", -1)
					if strings.Contains(strings.Join(TmpList, " "), "@") {
						index := LocateMentionPosition(TmpList)

						for _, position := range index {
							User := SearchForUser(TmpList[position])

							if User == nil {
								continue
							}
							TmpList[position] = User.Mention()
						}

					}
					s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("<%s>: %s", TmpList[3], strings.Join(TmpList[4:], " ")))
				}

			} else if strings.Contains(line.Text, "[JOIN]") ||
				strings.Contains(line.Text, "[LEAVE]") {
				TmpList := strings.Split(line.Text, " ")
				s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", strings.Join(TmpList[3:], " ")))

			} else if strings.Contains(line.Text, "FactoCord.ext: ") {
				TmpList := strings.Split(line.Text, " ")
				s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", strings.Join(TmpList[1:], " ")))

			} else {
				if strings.Contains(Config.LogConsole, "TRUE") {
					s.ChannelMessageSend(Config.FactorioConsoleChannelID, fmt.Sprintf("%s", line.Text))
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
