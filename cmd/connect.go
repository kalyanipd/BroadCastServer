package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the broadcast server as a client",
	Run: func(cmd *cobra.Command, args []string) {
		url := "ws://localhost:8080/ws"
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Fatal("Dial error:", err)
		}
		defer conn.Close()

		// Receive messages
		go func() {
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println("Read error:", err)
					return
				}
				fmt.Printf("Received: %s\n", msg)
			}
		}()

		// Send messages
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("You: ")
			if scanner.Scan() {
				text := scanner.Text()
				if text == "exit" {
					break
				}
				err := conn.WriteMessage(websocket.TextMessage, []byte(text))
				if err != nil {
					log.Println("Write error:", err)
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
