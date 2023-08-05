package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	// Initialize the HTML template engine
	htmlTemplate := template.Must(template.ParseGlob("*.tmpl"))

	// Add the HTML template engine to the Gin engine
	r.SetHTMLTemplate(htmlTemplate)

	// Define your API endpoints
	r.GET("/example", executeScriptHandler)
	r.GET("/websocket", func(c *gin.Context) {
		// Redirect to the output page when accessing /websocket
		c.Redirect(http.StatusFound, "/output")
	})
	r.GET("/output", websocketPageHandler) // HTTP handler for the output page
	r.GET("/ws", wsHandler)                // WebSocket handler

	// Run the server on port 8088
	r.Run(":8088")
}

func executeScriptHandler(c *gin.Context) {
	// Your Bash script command to execute
	cmd := exec.Command("bash", "./start.sh") // Replace with the actual path to your Bash script

	// Execute the Bash script
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error occurred while executing the script.",
		})
		return
	}

	// Respond with the output of the Bash script
	c.HTML(http.StatusOK, "result.tmpl", gin.H{
		"output": string(output),
	})
}

func websocketPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "websocket.tmpl", nil)
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	cmd := exec.Command("bash", "./start.sh") // Replace with the actual path to your Bash script

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error creating command pipe: %v", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting the script: %v", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := cmdReader.Read(buf)
		if err != nil {
			break
		}

		// Send the output of the Bash script to the client through the WebSocket connection
		err = conn.WriteMessage(websocket.TextMessage, buf[:n])
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			break
		}
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for the script to finish: %v", err)
		return
	}
}
