package action

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"rust-cmd-webrcon/structures"
	"time"
)

func SendCommand() error {
	var u = url.URL{Scheme: "ws", Host: Config.Host, Path: Config.Password}
	fmt.Printf("Connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("Dial error:" + err.Error())
	}
	defer c.Close()

	message, err := json.Marshal(Message)
	if err != nil {
		return errors.New("bad message:" + err.Error())
	}

	fmt.Println("Send Message:", string(message))

	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return errors.New("write error:" + err.Error())
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			return errors.New("read error:" + err.Error())
		}
		fmt.Println("Response:", string(message))

		var response structures.Response

		if err = json.Unmarshal(message, &response); err != nil {
			return errors.New("read error:" + err.Error())
		}

		if response.Identifier == Message.Identifier {
			time.AfterFunc(time.Second*3, func() {
				os.Exit(0)
			})
		}
	}

	return nil
}
