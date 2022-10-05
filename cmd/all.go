package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

var path string

func init() {
	sendCmd.AddCommand(allCmd)
	allCmd.Flags().StringVarP(&path, "p", "p", "p", "p")
	allCmd.MarkFlagRequired("p")
}

func sendAllFiles(s *discordgo.Session, event *discordgo.Ready) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, f := range files {
		folder := path + "/" + f.Name()
		file, err := os.Open(folder)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		s.ChannelFileSend(BotChan, filepath.Base(folder), file)
	}
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Envois tous les fichiers du dossier",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start()
		time.Sleep(2 * time.Second)
		fmt.Println("Tous les fichiers ont été envoyés")
	},
}
