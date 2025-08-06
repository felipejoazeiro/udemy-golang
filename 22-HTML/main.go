package main



var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob(".html"))

	http.HandleFunc("/home", func home(w http.ResponseWriter, r *http.Request){

		u := usuario{"João", "joao.pedro@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServer(":5000", nil))

}