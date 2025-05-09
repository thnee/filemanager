package cmd

import (
	"fmt"
	"main/internal/api"
	"main/internal/g"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  "Serve HTTP API and web app.",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().String("bind", "127.0.0.1:4000", "Host and port that the server listens on.")
	viper.BindPFlag("server.bind", serverCmd.Flags().Lookup("bind"))

	setServerCorsConfigDefaults()
	setServerSessionConfigDefaults()
}

func runServer(cmd *cobra.Command, args []string) {
	g.Session = newSessionManager()

	r := chi.NewRouter()

	customCors := newCors()

	r.Use(middleware.Logger)
	r.Use(customCors.Handler)
	r.Use(g.Session.LoadAndSave)

	api.RegisterRoutes(r)

	r.NotFound(api.NotFoundHandler)

	fmt.Printf("Starting server, listening on: %s\n", viper.GetString("server.bind"))
	http.ListenAndServe(viper.GetString("server.bind"), r)
}

func setServerCorsConfigDefaults() {
	viper.SetDefault("server.cors.allowed-methods", []string{"GET", "POST", "PUT", "OPTIONS"})
	viper.SetDefault("server.cors.allowed-headers", []string{"Content-Type"})
	viper.SetDefault("server.cors.allow-credentials", true)
	viper.SetDefault("server.cors.debug", false)
}

func setServerSessionConfigDefaults() {
	viper.SetDefault("server.session.cookie.name", "fm_session")
	viper.SetDefault("server.session.cookie.http-only", true)
	viper.SetDefault("server.session.cookie.http-path", "/")
	viper.SetDefault("server.session.cookie.persist", true)
	viper.SetDefault("server.session.cookie.same-site", "Default")
	viper.SetDefault("server.session.cookie.secure", true)
}

func newCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("server.cors.allowed-origins"),
		AllowedMethods:   viper.GetStringSlice("server.cors.allowed-methods"),
		AllowedHeaders:   viper.GetStringSlice("server.cors.allowed-headers"),
		AllowCredentials: viper.GetBool("server.cors.allow-credentials"),

		Debug: viper.GetBool("server.cors.debug"),
	})
}

func newSessionManager() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 3 * time.Hour
	s.IdleTimeout = 20 * time.Minute
	s.Cookie.Name = viper.GetString("server.session.cookie.name")
	s.Cookie.Domain = viper.GetString("server.session.cookie.domain")
	s.Cookie.HttpOnly = viper.GetBool("server.session.cookie.http-only")
	s.Cookie.Path = viper.GetString("server.session.cookie.path")
	s.Cookie.Persist = viper.GetBool("server.session.cookie.persist")
	s.Cookie.SameSite = ParseSameSite(viper.GetString("server.session.cookie.same-site"))
	s.Cookie.Secure = viper.GetBool("server.session.cookie.secure")
	return s
}

func ParseSameSite(s string) http.SameSite {
	switch s {
	case "Default":
		return http.SameSiteDefaultMode
	case "Strict":
		return http.SameSiteStrictMode
	case "Lax":
		return http.SameSiteLaxMode
	case "None":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}
