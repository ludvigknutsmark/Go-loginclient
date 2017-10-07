package main

import (
	"flag"
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "if yes, the app is in debug mode")
	window  *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		//Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/gopher.icns",
			AppIconDefaultPath: "resources/gopher.png",
		},
		Debug:    *debug,
		Homepage: "index.html",
		//MessageHandler: handleMessages,
		OnWait: func(_ *astilectron.Astilectron, w *astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			// Store global variables
			window = w

			// Listen to messages sent by webserver
			window.On(astilectron.EventNameWindowEventMessage, func(e astilectron.Event) (deleteListener bool) {
				//Parses the sent JSON string from index.js to get sepearate JSON key values.
				var m string
				e.Message.Unmarshal(&m)
				var parsed map[string]interface{}
				err := json.Unmarshal([]byte(m), &parsed)
				if err != nil {
					panic(err);
				}
				//Operation == 0 means Login
				if parsed["operation"] == "0"{
					err := Login(parsed["username"].(string), parsed["password"].(string))
					//Call the send function to answer the message sent from index.js
					sendMessage(window, err);
				} else {
					errR := RegisterUser(parsed["reg_username"].(string), parsed["reg_password"].(string), parsed["reg_password_confirm"].(string));
					sendMessage(window, errR)
				}
				return
			})


			// Add listeners on tray
			t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) { astilog.Info("Tray has been clicked!"); return })
			return nil
		},
		//RestoreAssets: RestoreAssets,
		TrayOptions: &astilectron.TrayOptions{
			Image:   astilectron.PtrStr("resources/tray.png"),
			Tooltip: astilectron.PtrStr("Wow, what a beautiful tray!"),
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#fff"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(600),
			Width:           astilectron.PtrInt(600),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
/* SendMessage takes an astilectron.Window pointer as an argument to send the
	 the message back to the same window*/
func sendMessage(w *astilectron.Window, message string) {
	w.Send(message)
	return
}
