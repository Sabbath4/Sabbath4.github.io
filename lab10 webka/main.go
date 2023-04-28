package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type secondHand struct {
	Id       int
	Name     string
	Price    string
	Original string
}

func main() {
	s := database()
	defer s.db.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/add", s.add)
	http.HandleFunc("/all", s.all)
	http.HandleFunc("/remove", s.remove)
	http.HandleFunc("/update", update)
	http.HandleFunc("/update2", s.update2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type server struct {
	db *sql.DB
}

func database() server {
	database, _ := sql.Open("sqlite3", "secondHand.db")
	server := server{db: database}
	return server
}

func (s *server) add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		price := r.FormValue("price")
		original := r.FormValue("original")
		_, err := s.db.Exec("insert into secondHand(name, price, original) values($1, $2, $3)", name, price, original)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles("static/add.html")
	t.Execute(w, nil)
}

func (s *server) all(w http.ResponseWriter, r *http.Request) {
	var secondHands []secondHand
	info, _ := s.db.Query("select * from secondHand;")
	for info.Next() {
		var secondHand secondHand
		err := info.Scan(&secondHand.Id, &secondHand.Name, &secondHand.Price, &secondHand.Original)
		if err != nil {
			log.Fatal(err)
		}
		secondHands = append(secondHands, secondHand)
	}
	t, _ := template.ParseFiles("static/all.html")
	t.Execute(w, secondHands)
}

func (s *server) remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		_, _ = s.db.Exec("delete from secondHand where id=$1", id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles("static/remove.html")
	t.Execute(w, nil)
}

func update(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/update.html")
	t.Execute(w, nil)
}

func (s *server) update2(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	price := r.FormValue("price")
	original := r.FormValue("original")
	_, _ = s.db.Exec("update secondHand set name=$1, price=$2, original=$3 where id=$4", name, price, original, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
