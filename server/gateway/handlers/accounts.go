package handlers

import (
	"net/http"
	"fmt"
)

func (ctx *HandlerContext) AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}
	accountName := r.URL.Query().Get("account")
	go ctx.AddAccountToChan(accountName)
	fmt.Printf("added to queue to scrape: %s\n", accountName)
	msg := struct{ Message string }{fmt.Sprintf("gh account: %s queued successfully!", accountName)}
	respond(w, msg)
}

func (ctx *HandlerContext) AddAccountToChan(accountName string) {
	ctx.AccountsQueue <- accountName
}
