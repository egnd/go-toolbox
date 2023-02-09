# tgchain

Chained wrapper for handling tg updates

### Example
```go
package main

import (
	"log"
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/egnd/go-toolbox/tg/tgchain"
)

func main() {
	api, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := api.GetUpdatesChan(u)

	tgchain.NewListener(nil, nil)
		.Add(tgchain.EventMessage, 
			// @TODO: your handlers, children of IEventHandler
		)
		.Listen(context.Background(), updates)
}

```
