package server

import (
	"fmt"
	"net/http"

	"github.com/Eldius/cors-interceptor-go/cors"
)

func Start(port int) error {
	host := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(host, cors.CORS(Routes()))
}
