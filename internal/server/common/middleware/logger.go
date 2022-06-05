package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go-api-server-example/internal/logging"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			log = logging.GetLoggerFromContext(r.Context())

			ww = middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 = time.Now()

			headersIn = make(map[string]string)
		)

		for _, k := range []string{"Accept", "Accept-Encoding", "Content-Type", "Authorization"} {
			v := r.Header.Get(k)
			if v != "" {
				headersIn[k] = v
			}
		}
		//delete(headersIn, "Authorization")

		//buf, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	log.Errorf("Error reading request body: %v", err.Error())
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//
		//reader := ioutil.NopCloser(bytes.NewBuffer(buf))
		//r.Body = reader

		defer func() {
			log.Debugw(fmt.Sprintf("%s %s", r.Method, r.URL.Path),
				"remoteIp", r.RemoteAddr,
				//"host", r.Host,
				"path", r.URL.Path,
				"query", r.URL.RawQuery,
				"proto", r.Proto,
				"method", r.Method,
				//"userAgent", r.UserAgent(),
				"status", ww.Status(),
				"latencyMs", time.Since(t1),
				"headersIn", headersIn,
				//"headersOut", ww.Header(),
				//"bytesIn", r.Header.Get("Content-Length"),
				//"bytesOut", ww.BytesWritten(),
				//"body", buf,
			)
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
