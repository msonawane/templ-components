package main

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/msonawane/fe/components"
)

func main() {

	// Start the server.
	fmt.Println("listening on http://localhost:8000")
	e := echo.New()

	metaTagsOpts := components.MetaTagsOpts{
		AppName: "Components",
		Title:   "Title",
	}

	layoutOpts := components.LayoutOpts{
		MetaTagsOpts: metaTagsOpts,
		BodyClass:    "bg-azure-lt",
	}

	successAlertOpt := components.SuccessAlertOpts
	successAlertOpt.Title = "Success Title"
	successAlertOpt.Text = "Success Text"

	dangerAlertOpts := components.DangerAlertOpts
	dangerAlertOpts.Title = "Danger Title"
	dangerAlertOpts.Text = "Danger Alert"

	infoAlertOpts := components.InfoAlertOpts
	infoAlertOpts.Title = "Info Title"

	warningAlertOpts := components.WarningAlertOpts
	warningAlertOpts.Title = "Warning Title"

	links := []components.LinkOpts{{
		HREF: "/",
		Text: "Components",
	},
		{
			HREF:   "#",
			Text:   "Alerts",
			Active: true,
		},
	}
	bc := components.BreadCrumbs(links...)
	// cardOption := components.CardOptions{
	// 	Title:       "Title",
	// 	BreadCrumbs: bc,
	// }

	pageOpts := components.PageOpts{
		Title:       "Alerts",
		BreadCrumbs: bc,
		Children: []templ.Component{
			components.Alert(successAlertOpt),
			components.Alert(dangerAlertOpts),
			components.Alert(infoAlertOpts),
			components.Alert(warningAlertOpts),
		},
	}

	e.GET("/", func(c echo.Context) error {
		return components.P(layoutOpts, pageOpts).Render(c.Request().Context(), c.Response().Writer)
	})

	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":8000"))
}
