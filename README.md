
# Golang Application for Student CRUD Operations
This is a Golang application that handles CRUD operations for Student using REST APIs. It uses Cassandra DB and the gocql library to connect to the database.
## Installation
1. Clone the repository:
   
   git clone https://github.com/sachinpatil55/student-go-example
   
2. Install the required dependencies:
   
   go mod download
   
   
4. Build the application:
   
   go build
   
5. Run the application:
   
  go run main.go
   
## API Endpoints
The application exposes the following REST APIs:
### Add a Student
- Method: POST
- Endpoint: /student/v1/students
- Request Body:
  json
  {​​
    "name": "Sachin",
    "age": 18,
    "class": "10th",
    "subject": "Math"
  }​​
  
- Response:
  json
  {​​
    "enrollmentNumber": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
  }​​
  
### Get All Students
- Method: GET
- Endpoint: /student/v1/students
- Response:
  json
  [
    {​​
      "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
      "name": "Sachin ",
      "age": 18,
      "class": "10th",
      "subject": "Math"
      "deleted": false"
    }​​,
    {​​
      "id": "f47ac10b-58cc-4372-a567-0e02b2c3d480",
      "name": "Joe",
      "age": 17,
      "class": "9th",
      "subject": "Science",
      "deleted": false"
    }​​
  ]
  
### Get One Student
- Method: GET
- Endpoint: /student/v1/students/{​​studentId}​​
- Response:
  json
  {​​
    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "name": "Sachin",
    "age": 18,
    "class": "10th",
    "subject": "Math",
    "deleted": false"
  }​​
  
### Delete a Student
- Method: DELETE
- Endpoint: /student/v1/students/{​​studentId}​​
- Response:
  json
  {​​
    "message": "Student deleted successfully"
  }​​
### force delete student from database 
- Method: DELETE
- Endpoint: /student/v1/students/forcedelete/{​​studentId}​​
- Response:
  json
  {​​
    "message": "Student has been deleted permenantly"
  }​​
### Get All Archived Students
- Method: GET
- Endpoint: /student/v1/students/archived
- Response:
  json
  [
    {​​
      "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
      "name": "Sachin ",
      "age": 18,
      "class": "10th",
      "subject": "Math"
      "deleted": true"
    }​​,
    {​​
      "id": "f47ac10b-58cc-4372-a567-0e02b2c3d480",
      "name": "Joe",
      "age": 17,
      "class": "9th",
      "subject": "Science",
      "deleted": true"
    }​​
  ]
  
