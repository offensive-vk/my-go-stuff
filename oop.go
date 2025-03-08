package main

import "fmt"

// **Encapsulation**: Define a `Person` struct with private and public fields
type Person struct {
    Name string
    age  int // Private field (lowercase, unexported)
}

// Constructor function for `Person`
func NewPerson(name string, age int) Person {
    return Person{Name: name, age: age}
}

// Method to access the private field (getter)
func (p Person) GetAge() int {
    return p.age
}

// **Inheritance via Composition**: `Employee` "inherits" from `Person`
type Employee struct {
    Person    // Embedding `Person` struct
    Position  string
    Salary    float64
}

// Method for `Employee`
func (e Employee) GetDetails() {
    fmt.Printf("Name: %s, Age: %d, Position: %s, Salary: $%.2f\n",
        e.Name, e.GetAge(), e.Position, e.Salary)
}

// **Abstraction & Polymorphism**: Define an interface
type Worker interface {
    Work()
}

// Implementing interface for `Employee`
func (e Employee) Work() {
    fmt.Println(e.Name, "is working as a", e.Position)
}

// Another struct implementing the `Worker` interface
type Freelancer struct {
    Name    string
    Project string
}

// Implementing the `Worker` interface
func (f Freelancer) Work() {
    fmt.Println(f.Name, "is freelancing on", f.Project)
}

// **Polymorphism in action**
func DoWork(w Worker) {
    w.Work()
}

func main() {
    // Creating an Employee (Composition + Encapsulation)
    emp := Employee{
        Person:   NewPerson("Alice", 30), // Using constructor function
        Position: "Software Engineer",
        Salary:   75000.00,
    }
    emp.GetDetails()

    // Creating a Freelancer (Another implementation of Worker)
    freelancer := Freelancer{Name: "Bob", Project: "Open-source contributions"}

    // **Polymorphism**: Different types (Employee & Freelancer) implementing `Worker` interface
    DoWork(emp)
    DoWork(freelancer)
}
