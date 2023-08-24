package main 

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Student struct {
	ID      gocql.UUID `json:"id"`
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Class   string     `json:"class"`
	Subject string     `json:"subject"`
	Deleted bool       `json:"deleted"`
}

var session *gocql.Session

func init() {
 cluster := gocql.NewCluster("127.0.0.1")
 cluster.Keyspace = "studentdb"
 var err error
 fmt.Println("Initalization Started")
 session, err = cluster.CreateSession()
 if err != nil {
  log.Fatal(err)
 }
}

func main() {
 defer session.Close()
 fmt.Println("main Started")
 router := mux.NewRouter()
 router.HandleFunc("/student/v1/students", createStudent).Methods("POST")
 router.HandleFunc("/student/v1/students/archived", getArchivedStudents).Methods("GET")
 router.HandleFunc("/student/v1/students", getAllStudents).Methods("GET")
 router.HandleFunc("/student/v1/students/{studentId}", getStudent).Methods("GET")
 router.HandleFunc("/student/v1/students/{studentId}", deleteStudent).Methods("DELETE")
 router.HandleFunc("/student/v1/students/forcedelete/{studentId}", forceDeleteStudent).Methods("DELETE")
 log.Fatal(http.ListenAndServe(":8080", router))
}

func createStudent(w http.ResponseWriter, r *http.Request) {
 var student Student
 err := json.NewDecoder(r.Body).Decode(&student)
 if err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }
 student.ID = gocql.TimeUUID()
 query := "INSERT INTO students (id, name, age, class, subject, deleted) VALUES (?, ?, ?, ?, ?, ?)"
 err = session.Query(query, student.ID, student.Name, student.Age, student.Class, student.Subject, false).Exec()
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 response := map[string]interface{}{
  "enrollmentNumber": student.ID,
 }
 json.NewEncoder(w).Encode(response)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
 query := "SELECT id, name, age, class, subject FROM students WHERE deleted = ?  ALLOW FILTERING"
 iter := session.Query(query, false).Iter()
 var students []Student
 var student Student
 for iter.Scan(&student.ID, &student.Name, &student.Age, &student.Class, &student.Subject) {
  students = append(students, student)
 }
 if err := iter.Close(); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 studentID, _ := gocql.ParseUUID(params["studentId"])
 query := "SELECT id, name, age, class, subject FROM students WHERE id = ? AND deleted = ? LIMIT 1  ALLOW FILTERING"
 iter := session.Query(query, studentID, false).Iter()
 var student Student
 if iter.Scan(&student.ID, &student.Name, &student.Age, &student.Class, &student.Subject) {
  json.NewEncoder(w).Encode(student)
 } else {
  http.NotFound(w, r)
 }
 if err := iter.Close(); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
}


func deleteStudent(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 studentID,_ := gocql.ParseUUID(params["studentId"])
 query := "UPDATE students SET deleted = ? WHERE id = ?"
 err := session.Query(query, true, studentID).Exec()
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 fmt.Fprintf(w, "Student with ID %s has been deleted", studentID)
}

func forceDeleteStudent(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 studentID,_ := gocql.ParseUUID(params["studentId"])
 query := "DELETE FROM students WHERE id = ?"
 err := session.Query(query, studentID).Exec()
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 fmt.Fprintf(w, "Student with ID %s has been deleted permenantly", studentID)
}

func getArchivedStudents(w http.ResponseWriter, r *http.Request)  {
 query := "SELECT id, name, age, class,  subject, deleted FROM students WHERE deleted = ? ALLOW FILTERING"
 iter := session.Query(query, true).Iter()
 var students []Student
 var student Student
 fmt.Println("outside loop")
 for iter.Scan(&student.ID, &student.Name, &student.Age, &student.Class, &student.Subject, &student.Deleted) {
  fmt.Println("inside loop")
  students = append(students, student)
 }
 if err := iter.Close(); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  fmt.Println("Error while getting Archived students")
  return
 }
 json.NewEncoder(w).Encode(students)	
}
