package homeHandler

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello REST!"))

}

