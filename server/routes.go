package server

import (
	"net/http"
	"time"

	"github.com/Eldius/things-manager-go/config"
	"github.com/Eldius/things-manager-go/model"
	"github.com/Eldius/webapp-healthcheck-go/health"
	"github.com/eldius/jwt-auth-go/auth"
	authRepo "github.com/eldius/jwt-auth-go/repository"
)

/*
Routes creates the app router
*/
func Routes() http.Handler {
	mux := http.NewServeMux()

	repo := model.NewRepository()
	db := repo.GetDB()
	authRepo := authRepo.NewRepositoryCustom(db)
	svc := auth.NewAuthServiceCustom(authRepo)
	h := auth.NewAuthHandlerCustom(svc)
	mux.HandleFunc("/login", h.HandleLogin())
	mux.Handle("/user", h.AuthInterceptor(h.HandleNewUser()))

	// TODO add auth
	mux.HandleFunc("/things", HandleThingRoot(repo))
	mux.HandleFunc("/things/", HandleThingWithId(repo))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	// Health check
	mux.HandleFunc("/health", health.BuildChecker([]health.ServiceChecker{
		health.NewDBChecker("main-db", repo.GetDB().DB(), time.Duration(2*time.Second)),
	}, map[string]string{
		"app":         "things-manager",
		"version":     config.GetVersion(),
		"build-date":  config.GetBuildDate(),
		"branch-name": config.GetBranchName(),
	}))

	return mux
}
