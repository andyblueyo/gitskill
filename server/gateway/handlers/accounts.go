package handlers

import (
	"net/http"
)

func (ctx *HandlerContext) AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}
	accountName := r.URL.Query().Get("account")
	ctx.AccountsQueue <- accountName
	msg := struct{ Message string }{"account queued successfully!"}
	respond(w, msg)
}
