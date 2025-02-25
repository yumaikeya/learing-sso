package main

import (
	"angya-backend/internal/photoApplication"
	"angya-backend/internal/spotApplication"
	"angya-backend/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SpotHandler struct{}
type PhotoHandler struct{}

func (h *SpotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.EnhanceResponseWriter(&w)
	spotUsecase := spotApplication.NewUsecase()

	switch r.Method {
	case "POST":
		body := make([]byte, r.ContentLength)
		r.Body.Read(body) // byte 配列にリクエストボディを読み込む

		spot, err := spotUsecase.Register(r.Context(), body)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, string(func() (b []byte) { j, _ := json.Marshal(spot); return j }()))
	case "GET":
		spots, err := spotUsecase.List(r.Context())
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, string(func() (b []byte) { j, _ := json.Marshal(spots); return j }()))
	}
}

func (h *PhotoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.EnhanceResponseWriter(&w)
	photoUsecase := photoApplication.NewUsecase()

	switch r.Method {
	case "POST":
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body) //ReadAllでResponse Bodyを読み切る

		photo, err := photoUsecase.Register(r.Context(), b)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Fprint(w, string(func() (b []byte) { j, _ := json.Marshal(photo); return j }()))
	case "GET":
		// photo, err := photoUsecase.List(r.Context())
		// if err != nil {
		// 	fmt.Fprint(w, err.Error())
		// }
		// fmt.Fprint(w, string(func() (b []byte) { j, _ := json.Marshal(photo); return j }()))
	}
}

func main() {
	s, p := SpotHandler{}, PhotoHandler{}

	server := http.Server{
		Addr:    ":8088",
		Handler: nil,
	}

	http.Handle("/api/spots", &s)
	http.Handle("/api/photos", &p)

	server.ListenAndServe()
}
