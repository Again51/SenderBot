package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

var Token string
var BotChan string
var fileName string
var BotId string
var config *configStruct
var goBot *discordgo.Session

type configStruct struct {
	Token   string `json:"Token"`
	BotChan string `json:"BotChan"`
}

func init() {
	file, err := ioutil.ReadFile("/home/user/Bureau/sender/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&fileName, "n", "n", "", "")
	sendCmd.MarkFlagRequired("n")
	sendCmd.Flags().StringVarP(&BotChan, "c", "c", "", "")

	Token = config.Token
	BotChan = config.BotChan
}

func start() {
	goBot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID
	goBot.AddHandler(sendAllFiles)
	goBot.AddHandler(sendF)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Envois du fichier...")
}

func sendF(s *discordgo.Session, event *discordgo.Ready) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s.ChannelFileSend(BotChan, filepath.Base(fileName), file)
}

var sendCmd = &cobra.Command{
	Use:   "file",
	Short: "Envois un fichier vers le serveur discord",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start()
		time.Sleep(2 * time.Second)
		fmt.Println("Le fichier", filepath.Base(fileName), "a été envoyé")
	},
}
