### [PoC] Telegram Bot

Proof-of-Concept of a Telegram Bot develop with Golang.
It uses [telebot.v2](https://github.com/tucnak/telebot/) library.

[Create Telegram Bot](https://core.telegram.org/bots#6-botfather)

The idea was to try to well organise the code and discover golang.

Some features:
- Limit the bot to answer only selected channels;
- Add commands (with help command and limit to admin usage only);
- Filter messages;
- User interaction for a command;

#### Configuration

All fields are strings.
```yaml
---
telegram_token: ''
authorized_channels:
  - ''
  - ''
```

