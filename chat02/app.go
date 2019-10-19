package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Meddlers *Middleware
	config   *Env
}

type shortenReq struct {
	URL                 string `json:"url"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes"`
}

type shortLinkResp struct {
	ShortLink string `json:"short_link"`
}

func (a *App) Init(env *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile) //时间日期 | 行号文件名
	a.config = env
	a.Router = mux.NewRouter()
	a.Meddlers = &Middleware{}
	a.initRoutes()
}

func (a *App) initRoutes() {
	m := alice.New(a.Meddlers.LoggingHandler, a.Meddlers.RecoverHandler)
	//a.Router.HandleFunc("/api/shorten", a.createShortenLink).Methods("POST")
	//a.Router.HandleFunc("/api/info", a.getShortLinkInfo).Methods("GET")
	//a.Router.HandleFunc("/{short_link:[a-zA-Z0-9]{1,11}}", a.redirect).Methods("GET")
	a.Router.Handle("/api/shorten",
		m.ThenFunc(a.createShortenLink)).Methods("POST")
	a.Router.Handle("/api/info",
		m.ThenFunc(a.getShortLinkInfo)).Methods("POST")
	a.Router.Handle("/{short_link:[a-zA-Z0-9]{1,11}}",
		m.ThenFunc(a.redirect)).Methods("POST")
}

func (a *App) createShortenLink(w http.ResponseWriter, r *http.Request) {
	var (
		req shortenReq
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("parse parameters failed %v", req),
		})
		return
	}

	if req.ExpirationInMinutes <= 0 {
		respondWithError(w, StatusError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("valid data err %v ", req),
		})
		return
	}

	defer r.Body.Close()

	if str, err := a.config.S.Shorten(req.URL, req.ExpirationInMinutes); err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortLinkResp{ShortLink: str})
	}
}

func (a *App) getShortLinkInfo(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()
	link := vals.Get("short_link")

	if detail, err := a.config.S.ShortLinkInfo(link); err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, detail)
	}
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Printf("%s \n", vars["short_link"])
	if link, err := a.config.S.UnShorten(vars["short_link"]); err != nil {
		respondWithError(w, err)
	} else {
		http.Redirect(w, r, link, http.StatusTemporaryRedirect)
	}
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e.Error())
		respondWithJSON(w, e.Status(), e.Error())
	default:
		respondWithJSON(w, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resp)
}
