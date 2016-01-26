package main

import (
	"net/http"
	"encoding/json"
	"strings"
)

func main(){
	http.HandleFunc("/hello", hello)
	//http handler for spotify route
	http.HandleFunc("/spotify", function(w http.ResponseWriter, r *http.Request){
			artist := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(artist)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}


		})

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.newEncoder(w).Encode(data)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello nick!"))
}

func query(artist string) (spotifyData, error) {
    response, err := http.Get("https://api.spotify.com/v1/search?q=" + artist + "&type=artist")
    if err != nil {
        return spotifyData{}, err
    }

    defer response.Body.Close()

    var data spotifyData

    if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
        return spotifyData{}, err
    }

    return data[0], nil
}

type spotifyData struct{
	Name string 'json:"name"'
	Popularity int 'json:"popularity"'
}
