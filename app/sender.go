package app

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"os"
	"strconv"
	"time"
)

func HelloHandler(ctx telebot.Context) error {
	user := ctx.Sender()
	info := user.FirstName
	if user.LastName != "" {
		info = fmt.Sprintf("%s %s", info, user.LastName)
	}
	return ctx.Send(fmt.Sprintf("Welcome %s", info))
}

func SendNewLine(b *telebot.Bot, l *logrus.Logger) {
	var fileSize int64
	fileName := os.Getenv("FILE")

	for {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			l.Fatal(err)
			return
		}

		if fileInfo.Size() > fileSize {
			file, err := os.Open(fileName)
			if err != nil {
				l.Fatal(err)
				return
			}

			_, err = file.Seek(fileSize, 0)
			if err != nil {
				l.Fatal(err)
				return
			}

			scanner := bufio.NewScanner(file)
			chatID, _ := strconv.Atoi(os.Getenv("USERID"))

			for scanner.Scan() {
				line := scanner.Text()
				if line != "" {
					newMessage := "Received 😍📨\n\n\n" + line
					_, err = b.Send(telebot.ChatID(chatID), newMessage)
					if err != nil {
						l.Fatal(err)
						return
					}
				}
			}
			fileSize = fileInfo.Size()
		}
		time.Sleep(time.Minute)
	}
}
