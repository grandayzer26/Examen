package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var tasks []Task

type Task struct {
	Description string
	Done        bool
	ID          int //Question 2.2 ajout de ID
}

func serialize(v any) []byte {
	ser, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return ser
}

//Question 2.1
func list(rw http.ResponseWriter, _ *http.Request) {

	//Question 2.2
	RJson := string(serialize(tasks))

	//Question 2.3
	s := ""
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done == false {
			s = s + string(serialize(tasks[i]))
		}
	}
	//Question 2.4
	rw.WriteHeader(http.StatusOK)

	//Question 2.5
	rw.Write([]byte("voici la liste des taches : \n"))
	rw.Write([]byte(RJson))
	rw.Write([]byte("\n"))
	rw.Write([]byte("voici la liste des taches non faites : \n"))
	rw.Write([]byte(s))

}

func done(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//Question 4.2
	case http.MethodGet:
		s := ""
		for i := 0; i < len(tasks); i++ {
			if tasks[i].Done {
				s = s + string(serialize(tasks[i]))
			}
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(s))

	//Question 4.4
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body : %v", err)
			http.Error(rw, "can't read body", http.StatusBadRequest)
			return
		}
		//Question 4.5
		rBody := string(body)
		intVar, _ := strconv.Atoi(rBody)
		if intVar < len(tasks) {
			tasks[intVar].Done = true
			rw.WriteHeader(http.StatusOK)

		} else {
			log.Fatal("Vous essayer de taper dans une case non existante ")
		}

	//Question 4.6
	default:
		rw.WriteHeader(http.StatusBadRequest)
	}

}

func add(rw http.ResponseWriter, r *http.Request) {
	//Question 3.1
	if r.Method == http.MethodPost {
		//Question 3.3
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body : %v", err)
			http.Error(rw, "can't read body", http.StatusBadRequest)
			return
		}

		//Question 3.4
		rBody := string(body)
		tache := Task{Description: rBody, Done: false, ID: len(tasks)}
		tasks = append(tasks, tache)
		//Question 3.5
		rw.WriteHeader(http.StatusCreated)
	} else { //Question 3.2
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Vous avez demander une mauvaise requete"))
	}

}

func main() {
	fmt.Println("suuu")

	//Question 1.5
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	//Question 1.6
	log.Fatal(http.ListenAndServe(":8083", nil))

	//Bonus
	/*
		en essayant de passer les handler dans le listenAndServe,
		 une erreur signal que le handler n'implemente pas la methode ServeHttp

	*/
}
