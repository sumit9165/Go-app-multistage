package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ImageData struct {
	Size      string `json:"size"`
	BuildTime string `json:"buildTime"`
	StartTime string `json:"startTime"`
}

type Response struct {
	LargeImage     ImageData `json:"largeImage"`
	OptimizedImage ImageData `json:"optimizedImage"`
}

func imageHandler(w http.ResponseWriter, r *http.Request) {

	data := Response{
		LargeImage: ImageData{
			Size:      "1.2 GB",
			BuildTime: "45 seconds",
			StartTime: "10 seconds",
		},
		OptimizedImage: ImageData{
			Size:      "120 MB",
			BuildTime: "8 seconds",
			StartTime: "2 seconds",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/api/images", imageHandler)

	log.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}
