package main

/*type Server struct {
	//DB     *gorm.DB
	Router *mux.Router
}*/

/*func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
	fmt.Println("RESTfulServ. on:8093, Controller:", r.URL.Path[1:])
}*/

func main() {
	Run()
	/*http.HandleFunc("/", handler)
	fmt.Println("Starting Restful services...")
	fmt.Println("Using port:8093")
	err := http.ListenAndServe(":8093", nil)
	log.Print(err)
	errorHandler(err)*/
}

/*func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}*/
