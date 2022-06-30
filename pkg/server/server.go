package server

import (
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository/boltdb"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server          *http.Server
	pocketClient    *pocket.Client
	tokenRepository *boltdb.TokenRepository
	redirectURL     string
}

func NewAuthorizationServer(pc *pocket.Client, tr *boltdb.TokenRepository,
	ru string) *AuthorizationServer {
	return &AuthorizationServer{pocketClient: pc, tokenRepository: tr, redirectURL: ru}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":80",
		Handler: s,
	}
	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	chatId, err := strconv.ParseInt(r.URL.Query().Get("chat_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	requestToken, err := s.tokenRepository.Get(chatId, repository.RequestToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authResp, err := s.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.tokenRepository.Save(chatId, authResp.AccessToken, repository.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("char_id\n%d\nrequest_token\n%s\naccess_token\n%s", chatId, requestToken, authResp.AccessToken)

	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
