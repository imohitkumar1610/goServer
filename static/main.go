package main

impport(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w , http.ResponseWriter, r *http.Request){
	if err := r.Parseform(); err != nil {
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, " POST request successfull!")
	name:= r.FormValue("name")
	address := r.FornValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func hellloHandler(w , http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main(){
	fileserver := http.fileserver(http.dir("./static"))
	http.handle("/",fileserver)
	http.handlefunc("/form",formhandler)
	http.handlefunc("/hello",hellohandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != null {
		log.Fatal(err)
	}
}