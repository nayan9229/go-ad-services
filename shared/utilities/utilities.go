package utilities

import (
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogSetup performs logging setup common to all services.
func LogSetup(appname string, dev bool) {
	baselog := zerolog.New(os.Stdout)
	if dev {
		baselog = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	applog := baselog.With().Timestamp().Str("service", appname).Logger()
	log.Logger = applog

	stdlog.SetFlags(0)
	stdlog.SetOutput(applog.With().Logger())
}

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				log.Info().
					Str("method", r.Method).
					Str("path", r.URL.Path).
					Int("status", ww.Status()).
					Int("bytes", ww.BytesWritten()).
					Dur("duration", time.Since(start)).
					Str("remote_addr", r.RemoteAddr).
					Str("request_id", middleware.GetReqID(r.Context())).
					Msg("Request completed")
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
