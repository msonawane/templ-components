package main

import (
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"github.com/msonawane/fe/components"
)

var (
	metaTagsOpts = components.MetaTagsOpts{
		AppName: "Components",
		Title:   "Title",
	}

	layoutOpts = components.LayoutOpts{
		MetaTagsOpts: metaTagsOpts,
		BodyClass:    "bg-blue-lt",
	}
)

func main() {
	e := echo.New()

	g := e.Group("")
	g.GET("/", alerts).Name = "alerts"
	g.GET("/buttons", buttons).Name = "buttons"
	g.GET("/cards", cards).Name = "cards"
	g.GET("/tables", tables).Name = "tables"

	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":8000"))
}

func alerts(ctx echo.Context) error {
	links := []components.LinkOpts{
		{
			HREF: "/",
			Text: "Components",
		},
		{
			HREF:   ctx.Echo().Reverse("alerts"),
			Text:   "Alerts",
			Active: true,
		},
	}
	bc := components.BreadCrumbs(links...)
	if htmx.IsHTMX(ctx.Request()) {
		return htmx.NewResponse().
			RenderTempl(ctx.Request().Context(), ctx.Response(), components.AlertCard(bc))
	}

	return components.AlertPage(layoutOpts, bc).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func buttons(ctx echo.Context) error {
	nm := NavMenu(ctx)["buttons"]
	nm.Active = true
	links := []components.LinkOpts{
		NavMenu(ctx)["root"],
		nm,
	}
	bc := components.BreadCrumbs(links...)
	if htmx.IsHTMX(ctx.Request()) {
		return htmx.NewResponse().
			RenderTempl(ctx.Request().Context(), ctx.Response(), components.ButtonCard(bc))
	}

	return components.ButtonPage(layoutOpts, bc).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func cards(ctx echo.Context) error {
	nm := NavMenu(ctx)["cards"]
	nm.Active = true
	links := []components.LinkOpts{
		NavMenu(ctx)["root"],
		nm,
	}
	bc := components.BreadCrumbs(links...)

	if htmx.IsHTMX(ctx.Request()) {
		return htmx.NewResponse().
			RenderTempl(ctx.Request().Context(), ctx.Response(), components.CardShowCase(bc))
	}
	return components.CardPage(layoutOpts, bc).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func tables(ctx echo.Context) error {
	nm := NavMenu(ctx)["tables"]
	nm.Active = true
	links := []components.LinkOpts{
		NavMenu(ctx)["root"],
		nm,
	}
	bc := components.BreadCrumbs(links...)
	if htmx.IsHTMX(ctx.Request()) {
		return htmx.NewResponse().
			RenderTempl(ctx.Request().Context(), ctx.Response(), components.TableCard(bc))
	}

	return components.TablePage(layoutOpts, bc).Render(ctx.Request().Context(), ctx.Response().Writer)
}

func NavMenu(ctx echo.Context) map[string]components.LinkOpts {
	menuLinks := map[string]components.LinkOpts{}
	menuLinks["root"] = components.LinkOpts{
		HREF: "/",
		Text: "Components",
	}
	menuLinks["alerts"] = components.LinkOpts{
		HREF: ctx.Echo().Reverse("alerts"),
		Text: "Alerts",
	}
	menuLinks["buttons"] = components.LinkOpts{
		HREF: ctx.Echo().Reverse("buttons"),
		Text: "Buttons",
	}
	menuLinks["cards"] = components.LinkOpts{
		HREF: ctx.Echo().Reverse("cards"),
		Text: "Card with Action buttons and Pagination",
	}
	menuLinks["tables"] = components.LinkOpts{
		HREF: ctx.Echo().Reverse("tables"),
		Text: "Tables",
	}

	return menuLinks
}
