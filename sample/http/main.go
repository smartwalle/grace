package main

import (
	"fmt"
	"github.com/smartwalle/nx"
	"net/http"
	"os"
)

func main() {
	nx.Serve(
		&http.Server{Addr: ":9900", Handler: newHandler()},
	)
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "hello %d \n", os.Getpid())

		//duration, err := time.ParseDuration(r.FormValue("duration"))
		//if err != nil {
		//	http.Error(w, err.Error(), 400)
		//	return
		//}
		//time.Sleep(duration)
		//fmt.Fprintf(
		//	w,
		//	"%s started at %s slept for %d nanoseconds from pid %d.\n",
		//	name,
		//	now,
		//	duration.Nanoseconds(),
		//	os.Getpid(),
		//)
	})
	return mux
}