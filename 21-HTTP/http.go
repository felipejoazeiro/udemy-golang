package main

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Carregar páginas!"))
	}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServer(":5000", nil))

}