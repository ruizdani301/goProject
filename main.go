// main.go
package main
import (
	"log"
	"apidgraph/dgraph"
	"apidgraph/api"
	"net/http"
	"github.com/go-chi/chi" // pendiente intslacion
	"context"
	"github.com/go-chi/cors"

)

func main() {
	ctx := context.Background()
	cliente := dgraph.NewClient()

	apidg := api.NewApidgraph(cliente, ctx)

	//	newClient()

	router := chi.NewRouter()
	// configuration of cors 
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", /*"X-CSRF-Token"*/},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	// path to Get all programs 
	router.Get("/programs", apidg.GetAllPrograms)
	// path to get one program by uid
	router.Get("/programs/{uid}", apidg.GetProgram)
	//path to create a program
	router.Post("/programs", apidg.CreateProgram)
	 //path to update a program 
	router.Put("/programs", apidg.UpdateProgram)
	// path to delete a program by uid
	router.Delete("/programs", apidg.DeleteProgram)

	log.Println("corriendo")
	log.Fatal(http.ListenAndServe(":5000", router))
}
