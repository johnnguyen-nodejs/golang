package main
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/johnnguyen-nodejs/golang/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api")
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	http.ListenAndServe(":9090", nil)
}
