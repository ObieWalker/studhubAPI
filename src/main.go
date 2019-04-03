package main

import (
  "encoding/json"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)

type Student struct {
  ID        string   `json:"id,omitempty"`
  Firstname string   `json:"firstname,omitempty"`
  Lastname  string   `json:"lastname,omitempty"`
  Level     *Level   `json:"level,omitempty"`
}
type Level struct {
  Faculty  string `json:"faculty,omitempty"`
  Class string `json:"class,omitempty"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range students {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Student{})
}

func createStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var student Student
  _ = json.NewDecoder(r.Body).Decode(&student)
  student.ID = strconv.Itoa(rand.Intn(10000000))
  students = append(students, student)
  json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range students {
    if item.ID == params["id"] 
      students = append(students[:index], students[index+1:]...)
      break
  }
  json.NewEncoder(w).Encode(students)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range students {
    if item.ID == params["id"]
      students = append(students[:index], students[index+1:]...)
      var student Student
      _ = json.NewDecoder(r.Body).Decode(&student)
      student.ID = params["id"]
      students = append(students, student)
      json.NewEncoder(w).Encode(student)
      return
  }
  json.NewEncoder(w).Encode(students)
}


func main() {
  router := mux.NewRouter()

  students = append(students, Student{ID: "1", Firstname: "John", Lastname: "Doe", Level: &Level{Faculty: "Science", Class: "Primary 5"}})
  students = append(students, Student{ID: "2", Firstname: "Koko", Lastname: "Doe", Level: &Level{Faculty: "Arts", Class: "Primary 5"}})
  students = append(students, Student{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

  router.HandleFunc("/students", getStudents).Methods("GET")
  router.HandleFunc("/student/{id}", getStudent).Methods("GET")
  router.HandleFunc("/student/{id}", createStudent).Methods("POST")
  router.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
  router.HandleFunc("/student/{id}", updateStudent).Methods("UPDATE")
  log.Fatal(http.ListenAndServe(":8000", router))




}