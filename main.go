package main

import (
	"anya/enums"
	"anya/handlers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var ListenPort = 80

func init() {

}

func main() {
	router := chi.NewRouter()
	
	router.Group(func(r chi.Router) {
		r.Get("/ping", handlers.Pong)
	})

	router.Group(func(r chi.Router) {
		r.Get("/services_health", handlers.ServiceHealth)
		r.Post("register_service", handlers.RegisterService)
		r.Post("update_health", handlers.UpdateHealth)
	})

	router.Group(func(r chi.Router) {
		r.Get("shared_svc_config", handlers.SharedSVConfig)
	})

	router.Group(func(r chi.Router) {
		r.Post("update_config", handlers.SetConfig)
		r.Post("send_to_topic", handlers.SendToTopic)
	})

	router.NotFound(handlers.NotFound)
	router.MethodNotAllowed(handlers.MethodNotAllowed)

	fmt.Println(enums.Terminalbold + enums.TerminalcCyan + "Anya Listening on port: " + enums.Terminalitalic + strconv.Itoa(ListenPort) + enums.Terminalreset)
	fmt.Println(enums.Terminalbold + enums.TerminalcRed + "Check Directory:" + enums.TerminalcWhite + enums.Terminalitalic + " logs/" + enums.Terminalreset + " for detailed logging and Error reports" + enums.Terminalreset)
	http.ListenAndServe(":"+strconv.Itoa(ListenPort), router)
}
