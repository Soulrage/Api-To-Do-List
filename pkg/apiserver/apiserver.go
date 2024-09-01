package apiserver
import(
	"net/http"
	"time"
	"context"
)

type ApiServer struct{
	httpServer *http.Server
}

func (s *ApiServer) Run(port string, hanlder http.Handler) error{
	s.httpServer = &http.Server{
		Addr:			":" + port,
		Handler:       hanlder,
		MaxHeaderBytes:  1<<20,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}


func (s *ApiServer) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}