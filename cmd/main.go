package main

import (
	"flag"

	"github.com/Oppodelldog/cruda-test/cases/crud"
	"github.com/Oppodelldog/cruda-test/cases/jsx"
	"github.com/Oppodelldog/cruda-test/cases/tools"
	"github.com/Oppodelldog/cruda-test/runner"
)

const defaultURL = "http://172.17.0.2:3000"

func main() {
	var url string

	flag.StringVar(&url, "url", defaultURL, "-url=http://172.17.0.2:3000")
	flag.Parse()

	runner.Suites(url,
		runner.TestSuites{
			"CRUD": runner.TestSuite{
				"Case01PageWorks":                crud.Case01PageWorks,
				"Case02CreateItem":               crud.Case02CreateItem,
				"Case02CreateItemFormValidation": crud.Case02CreateItemFormValidation,
				"Case02CreateItemServerError":    crud.Case02CreateItemServerError,
				"Case03SelectItem":               crud.Case03SelectItem,
				"Case03SelectItemServerError":    crud.Case03SelectItemServerError,
				"Case04UpdateItem":               crud.Case04UpdateItem,
				"Case04UpdateItemServerError":    crud.Case04UpdateItemServerError,
				"Case05DeleteItem":               crud.Case05DeleteItem,
				"Case05DeleteItemServerError":    crud.Case05DeleteItemServerError,
			},
			"JSX Page": runner.TestSuite{
				"Case01CallJSXPage": jsx.Case01CallJSXPage,
			},
			"Tools": runner.TestSuite{
				"Case01ToolsExtension_CreatedEvent":  tools.Case01ToolsExtensionReceiveCreateEvent,
				"Case02ToolsExtension_SelectedEvent": tools.Case02ToolsExtensionReceiveSelectEvent,
				"Case03ToolsExtension_UpdatedEvent":  tools.Case03ToolsExtensionReceiveUpdateEvent,
				"Case04ToolsExtension_DeleteEvent":   tools.Case04ToolsExtensionReceiveDeleteEvent,
			},
		},
	)
}
