package main
import (
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "testing"
 "bytes"
 "github.com/gocql/gocql"
 "github.com/gorilla/mux"
 "io"
 "fmt"
)

type Student1 struct {
	ID      gocql.UUID `json:"id"`
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Class   string     `json:"class"`
	Subject string     `json:"subject"`
	Deleted bool       `json:"deleted"`
}

func TestCreateStudent(t *testing.T) {
 router := mux.NewRouter()
 router.HandleFunc("/student/v1/students", createStudent).Methods("POST")
 studentdata:= Student1{
	Name: "John", 
	Age: 18, 
	Class: "B", 
	Subject: "Math",
 } 
 payload, _ := json.Marshal(&studentdata)
 req, err := http.NewRequest("POST", "/student/v1/students", bytes.NewReader(payload))
 if err != nil {
  t.Fatal(err)
 }
 recorder := httptest.NewRecorder()
 router.ServeHTTP(recorder, req)
 if recorder.Code != http.StatusOK {
  t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
 }
 var response map[string]interface{}
 err = json.Unmarshal(recorder.Body.Bytes(), &response)
 if err != nil {
  t.Fatal(err)
 }
 if _, ok := response["enrollmentNumber"]; !ok {
  t.Error("Expected enrollmentNumber field in response")
 }
}
func TestGetAllStudents(t *testing.T) {
 router := mux.NewRouter()
 router.HandleFunc("/student/v1/students", getAllStudents).Methods("GET")
 req, err := http.NewRequest("GET", "/student/v1/students", nil)
 if err != nil {
  t.Fatal(err)
 }
 recorder := httptest.NewRecorder()
 router.ServeHTTP(recorder, req)
 if recorder.Code != http.StatusOK {
  t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
 }
 var students []Student
 err = json.Unmarshal(recorder.Body.Bytes(), &students)
 if err != nil {
  t.Fatal(err)
 }
}
func TestGetStudent(t *testing.T) {
 router := mux.NewRouter()
 router.HandleFunc("/student/v1/students/{studentId}", getStudent).Methods("GET")
 req, err := http.NewRequest("GET", "/student/v1/students/e85b1da2-4131-11ee-b403-b600937f0704", nil)
 if err != nil {
  t.Fatal(err)
 }
 recorder := httptest.NewRecorder()
 router.ServeHTTP(recorder, req)
 if recorder.Code != http.StatusOK {
  t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
 }
 var student Student
 err = json.Unmarshal(recorder.Body.Bytes(), &student)
 if err != nil {
  t.Fatal(err)
 }
}
func TestDeleteStudent(t *testing.T) {
 router := mux.NewRouter()
 router.HandleFunc("/student/v1/students/{studentId}", deleteStudent).Methods("DELETE")
 req, err := http.NewRequest("DELETE", "/student/v1/students/123", nil)
 if err != nil {
  t.Fatal(err)
 }
 recorder := httptest.NewRecorder()
 router.ServeHTTP(recorder, req)
 
 resp := recorder.Result()
 body, _ := io.ReadAll(resp.Body)
 fmt.Println(string(body))
 if recorder.Code != http.StatusOK {
  t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
 }
 expectedResponse := "Student with ID 123 has been deleted"
 if recorder.Body.String() != expectedResponse {
  t.Errorf("Expected response '%s', but got '%s'", expectedResponse, recorder.Body.String())
 }
}