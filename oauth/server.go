package oauth

import (
	"encoding/json"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"
	"log"
	"net/http"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Oauth struct {
	*server.Server
	clientStore *store.ClientStore
	config      *Config
}

type Config struct {
	Domain string
}

type Credential struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func New(config *Config) *Oauth {

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return &Oauth{
		srv,
		clientStore,
		config,
	}
}

func (s *Oauth) GetCredentials() (*Credential, error) {
	// {"CLIENT_ID":"189fe9e3","CLIENT_SECRET":"28f6f908"}
	credential := &Credential{
		uuid.New().String()[:28],
		uuid.New().String()[:28],
	}
	domain := s.config.Domain
	if domain == "" {
		domain = "http://localhost:9094"
	}
	if err := s.clientStore.Set(credential.ClientID, &models.Client{
		ID:     credential.ClientID,
		Secret: credential.ClientSecret,
		Domain: domain,
	}); err != nil {
		return nil, err
	}
	return credential, nil
}

func (s *Oauth) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	s.HandleTokenRequest(w, r)
}

func (s *Oauth) MiddlewareFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := s.ValidationBearerToken(r)
		if err != nil {
			a, _ := json.Marshal(map[string]string{
				"message": err.Error(),
			})
			http.Error(w, string(a), http.StatusBadRequest)
			return
		}
		f.ServeHTTP(w, r)
	}
}

func (s *Oauth) Middleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := s.ValidationBearerToken(r)
		if err != nil {
			a, _ := json.Marshal(map[string]string{
				"message": err.Error(),
			})
			http.Error(w, string(a), http.StatusBadRequest)
			return
		}
		f.ServeHTTP(w, r)
	})
	//return http.HandlerFunc(f.ServeHTTP)
}
