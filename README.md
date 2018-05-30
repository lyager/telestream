## Introduction

I created this small program, both to play a little more with Go, but also
because I have numerous unattended machines at home. Wanting to keep their Linux
distributions up to date, I run
[cron-apt](https://debian-administration.org/article/162/A_short_introduction_to_cron-apt)
(yeah, they are Debian based), in order
to monitor if new packages become available.

But the default way of using `cron-apt` is to have it send an email. Setting
up an email server has become increasingly difficult, as an effect of Google
tightening the screw on spam - and thank God for that.

So I made a small program that can watch a file for a certain keyword and send
that line to a Telegram receiver (channel or person)

## Setup

You'll need the Telegram receiver ID and a Bot token before your start. It's not
that bad. Bot tokens you can get from the almighty
[BotFather](https://core.telegram.org/bots#6-botfather). The receiver ID is a
little bit more tricky, but after creating the Bot (with a name) you can invite
the bot into the intended group, and [follow this
guide](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id).
You'll only need to do this once.

## Running

Example

    go run main.go --token <BotToken> --receiver <ReceiverID> --filename /var/log/system.log --filter "respawn"

this creates a watcher of file `/var/log/system.log` looking for keyword
`respawn` and sends those lines to <ReceiverID>

## TODO

* Create JSON configuration - make behaviour more like LogWatcher
* Make the program shut down proper, by letting all go-routines listen to
  signals from channel