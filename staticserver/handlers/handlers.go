package handlers

import (
	"net/http"
	"os"
	"strings"
)

type MainHandler struct {
	landingPageHandler http.HandlerFunc
	reactPageHandler   http.HandlerFunc
}

func New(landingPageHandler http.HandlerFunc, reactPageHandler http.HandlerFunc) *MainHandler {
	return &MainHandler{landingPageHandler: landingPageHandler, reactPageHandler: reactPageHandler}
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	found := strings.HasPrefix(r.URL.Path, "/"+os.Getenv("LP_PATH"))
	if found {
		h.landingPageHandler(w, r)
		return
	} else {
		h.reactPageHandler(w, r)
		return
	}
}
