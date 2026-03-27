package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
)

type AuthParams struct {
	Title    string
	Body     []nodx.Node
	Language string
}

func Auth(params AuthParams) nodx.Node {
	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	body := nodx.Group(
		nodx.ClassMap{
			"w-screen h-screen px-4 py-[40px]":    true,
			"grid grid-cols-1 place-items-center": true,
			"bg-base-300 overflow-y-auto":         true,
		},
		nodx.Div(
			nodx.Class("w-full max-w-[600px] space-y-4"),
			nodx.Main(
				nodx.Class("rounded-box shadow-md bg-base-100 p-4"),
				nodx.Group(params.Body...),
			),
			nodx.Div(
				nodx.Class("flex justify-start space-x-2 items-center"),
				component.ChangeThemeButton(component.ChangeThemeButtonParams{
					Position:    component.DropdownPositionTop,
					AlignsToEnd: false,
					Size:        component.SizeMd,
					Language:    params.Language,
				}),
				component.ChangeLanguageButton(component.ChangeLanguageButtonParams{
					Position:    component.DropdownPositionTop,
					AlignsToEnd: false,
					Size:        component.SizeMd,
					Language:    params.Language,
				}),
			),
		),
	)

	return commonHtmlDoc(title, body)
}
