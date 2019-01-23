package actions

import (
	"microlog/web/component"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	view := component.CreateView("web/tpl/layout.html", "web/tpl/users.html")

	view.Render(w)
}
