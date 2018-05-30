package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"flag"

	tail "github.com/hpcloud/tail"
	tb "gopkg.in/tucnak/telebot.v2"
)

// MyPoller - class needed for telebot
type MyPoller struct {
}

// Poll - The main telegram bot poller
func (p *MyPoller) Poll(b *tb.Bot, dest chan tb.Update, stop chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Unable to read further from stdin: ", err)
			break
		}
		user := tb.User{ID: -305152601}
		b.Send(&user, "`"+line+"`", &tb.SendOptions{ParseMode: "Markdown"})
	}
}

// fileTail open as file (filename) filters it trough `filter`
// and sends matching output to `output`
// `filter` == "" means all lines match
func fileTail(filename string, output chan string, filter string) {
	t, err := tail.TailFile(filename, tail.Config{Follow: true})
	if err != nil {
		log.Fatalln("Failed to tail file: ", filename)
		return
	}

	defer t.Cleanup()

	for line := range t.Lines {
		if line.Err != nil {
			log.Fatalln("Got error while reading file: ", filename)
		}
		if strings.Contains(line.Text, filter) {
			output <- line.Text
		}
	}

}
func main() {

	filename := flag.String("filename", "", "Filename to read from")
	botToken := flag.String("token", "", "Telegram bot token (TODO how to get)")
	receiverID := flag.Int("receiver", 0, "Telegram receiver to send messages to")
	filter := flag.String("filter", "", "Filter each output line, match with filter")

	flag.Parse()

	log.Println("Bot token:", *botToken)
	log.Println("Reading from file:", *filename)
	log.Println("Filter:", *filter)

	if *botToken == "" {
		log.Panic("Bot token not specified")
	}

	if *receiverID == 0 {
		log.Panic("Need receiver to send messages to")
	}

	b, err := tb.NewBot(tb.Settings{
		Token: *botToken,
		//Poller: &MyPoller{},
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	b.Start() // Currently the poller is not used TODO

	if err != nil {
		log.Panic(err)
	}

	lineOut := make(chan string)
	go fileTail(*filename, lineOut, *filter)

	user := tb.User{ID: *receiverID}
	for {
		select {
		case line := <-lineOut:
			//log.Println("Outter, got line: ", line)
			b.Send(&user, "`"+line+"`", &tb.SendOptions{ParseMode: "Markdown"})
		default:
			time.Sleep(1)
		}
	}

	//b.Start()

}
