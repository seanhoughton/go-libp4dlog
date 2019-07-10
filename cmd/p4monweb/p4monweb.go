package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// func main() {
// 	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Welcome to my website!")
// 	})

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	http.ListenAndServe(":8001", nil)
// }

type msg struct {
	Message string `json:"msg"` //= {"msg":"Hello Golang Mpls Conf!"}
}

func main() {

	port := os.Getenv("PORT") // port string
	if port == "" {
		port = "8001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(&msg{Message: "Hello Golang MPLS Conf"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	msg := `	323 I p4dtguser  00:00:05 IDLE<br>
	1321 I p4dtguser  00:00:04 IDLE <br>
	1337 I p4dtguser  00:00:03 IDLE <br>
	1353 I p4dtguser  00:00:03 IDLE <br>
	1369 I p4dtguser  00:00:02 IDLE <br>
	1847 I swarm      00:04:36 IDLE <br>
	2123 B remote     193:06:33 ldapsync<br> 
	3283 I p4dtguser  00:00:03 IDLE <br>
	3318 R rcowham    00:00:00 monitor <br>
	3332 I p4dtguser  00:00:02 IDLE <br>
	9835 I svc_p4d_ha 00:38:30 IDLE <br>
	9837 I svc_p4d_fr 00:00:00 IDLE <br>
	10348 I svc_p4d_fr 00:15:19 IDLE <br>
	10891 I svc_p4d_fr 00:03:40 IDLE <br>
	12245 I svc_p4d_fr 00:05:34 IDLE <br>
	16719 I svc_p4d_ha 00:43:35 IDLE <br>
	16729 I svc_p4d_ha 00:52:00 IDLE <br>
	17590 I svc_p4d_ha 00:15:19 IDLE <br>
	17894 I svc_p4d_fr 00:02:39 IDLE <br>
	17895 I svc_p4d_ha 01:24:29 IDLE <br>
	19628 I svc_p4d_ha 00:54:10 IDLE <br>
	20271 R svc_p4d_ha 00:00:04 rmt-Journal<br> 
	22056 I svc_p4d_ha 00:19:10 IDLE <br>
	22157 I svc_p4d_fr 00:10:57 IDLE <br>
	24428 I svc_p4d_ha 00:26:54 IDLE <br>
	24429 I svc_p4d_ha 00:22:44 IDLE <br>
	24430 I svc_p4d_ha 00:02:39 IDLE <br>
	29455 I p4dtguser  00:00:01 IDLE <br>
	32663 I p4dtguser  00:00:03 IDLE <br>
   `

	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "%s", msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("Running server on port: %s\nType Ctr-c to shutdown server.\n", port)
	http.ListenAndServe(":"+port, nil)
}
