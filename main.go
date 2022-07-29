package main

import (
 "fmt"
 "log"
 "native-go-api/db"
 "native-go-api/handler"
 "native-go-api/models"
 "net/http"
 "os"
)

func main() {

log.Print("The is Server Running on localhost port 3000")

// initialize the database
db.Moviedb["001"] = models.Movie{ID: "001", Title: "A Space Odyssey", Description: "Science fiction"}
db.Moviedb["002"] = models.Movie{ID: "002", Title: "Citizen Kane", Description: "Drama"}
db.Moviedb["003"] = models.Movie{ID: "003", Title: "Raiders of the Lost Ark", Description: "Action and adventure"}
db.Moviedb["004"] = models.Movie{ID: "004", Title: "66. The General", Description: "Comedy"}


// route goes here

// test route
http.HandleFunc("/", handler.TestHandler)
// get movies
http.HandleFunc("movies", handler.GetMovies)
// get a single movie
http.HandleFunc("/movie", handler.GetMovie)
// Add movie
http.HandleFunc("/movie/add", handler.AddMovie)
// delete movie
http.HandleFunc("/movie/delete", handler.DeleteMovie)


// listen port
err := http.ListenAndServe(":3000", nil)
// print any server-based error messages
if err != nil {
 fmt.Println(err)
 os.Exit(1)
}
}