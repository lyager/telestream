package main

import (
	"bufio"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

type MyPoller struct {
}

func (p *MyPoller) Poll(b *tb.Bot, dest chan tb.Update, stop chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//log.("Failed to read from Stdin", err)
			break
		}
		user := tb.User{ID: -305152601}
		b.Send(&user, line)
	}
}

func main() {
	bot_token := os.Getenv("BOT_TOKEN")

	if bot_token == "" {
		log.Panic("Bot token not specified")
	}
	b, err := tb.NewBot(tb.Settings{
		Token:  bot_token,
		Poller: &MyPoller{},
		//Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Panic(err)
	}

	//b.Send(tb.Recipient("TelestreamGroup"), "YO!")
	b.Handle("/hi", func(m *tb.Message) {
		// TODO Why is 'b' actually visible in here?
		log.Println("Channel: ", m.Chat)
		b.Send(m.Sender, "Hi "+m.Sender.Username)
	})

	b.Start()

}
