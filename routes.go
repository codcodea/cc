package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	pk "github.com/codcodea/cc/db"
	"github.com/codcodea/cc/types"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// HandleRoot GET /
func HandleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Color API is live and kicking!")
}

// HandleColor GET /colors/:hex
func HandleColor(c echo.Context) error {
	color := "#" + c.Param("hex")
	if len(color) != 7 {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	jsonData, err := getColor(color)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	return c.JSON(http.StatusOK, jsonData)
}

// getColor 
// - constructs a new response object
// - call and populate basic conversions
// - call and populate advanced conversions
// - call and populate color names
// - call and populate gradient
func getColor(color string) (types.Response, error) {

	var res = new(types.Response)

	ref, err := pk.Convert(color, res)

	if err != nil {
		return types.Response{}, err
	}

	res.Base.Color.Name = GetColorName(&ref)
	AddNames(&ref, res)
	AddRAL(&ref, res)
	AddPAN(&ref, res)
	AddNCS(&ref, res)
	AddMono(color, &ref, res)

	return *res, nil
}


// HandleLookup GET /lookup
// This is a new feature to be launched at mycolorpicker.com, currently in testing
// It route will not work without the client side feature branch

func HandleLookup(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		session := NewFormSession()
		fmt.Println("Session initialized")

		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				fmt.Println("Read socket err", err) // add logging later
				return
			}

			fmt.Println("Message received: ", msg)

			if len(msg) > 1 {

				// // limit the number of msg to 50
				// if len(session.LastResult.Names) > 50 {
				// 	session.LastResult.Names = session.LastResult.Names[:50]
				// }

				Send(msg, session, ws)
			}
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func Send(query string, s *FormSession, socket *websocket.Conn) error {

	ColorLookUp(query, s)

	if(len(s.LastResult) > 50) {
		s.LastResult = s.LastResult[50:]
	}

	// Marshal the LastResult.Names to JSON
	data, err := json.Marshal(s.LastResult)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return err
	}

	// Send the JSON data over the WebSocket
	err = websocket.Message.Send(socket, string(data))
	if err != nil {
		fmt.Println("Write socket error:", err) // Add logging later
		return err
	}
	return nil
}
