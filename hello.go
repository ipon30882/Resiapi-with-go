package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID  string  `json:"imdbID"`
	Title   string  `json:"title"`
	Year    int     `json:"year"`
	Rating  float32 `json:"rating"`
	IsSuper bool    `json:"isSuperHero"`
}

var movies []Movie

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //Errorตอบกลับ cilen StatusBadRequest
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //Errorตอบกลับ cilen StatusBadRequest
			fmt.Fprintf(w, "error : %s", err)
			return
		}

		movies = append(movies, t)
		w.WriteHeader(http.StatusOK) // StatusOK ใช้งานได้
		fmt.Fprintf(w, "hello %s Create movies", "POST")
		return
	}

	b, err := json.Marshal(movies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //Errorตอบกลับ cilen StatusInternalServerError
		fmt.Fprintf(w, "error : %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8") //เปลี่ยนtextเป็นjson
	w.Write(b)
	//fmt.Fprintf(w, " %s ", string(b))
}

func main() {
	http.HandleFunc("/movie", moviesHandler)

	err := http.ListenAndServe("localhost:2556", nil)
	log.Fatal(err)

}
