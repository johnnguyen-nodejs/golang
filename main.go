package main
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	hh := handlers.NewHello()
	http.ListenAndServe(":9090", nil)
}
