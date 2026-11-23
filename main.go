package main

import 
 
	"fmt"


type Task struct{
	ID int32
	Title string
	Done bool
}

type Tasklist struct {
	tasks []Task
}

func (t *Tasklist) create(id int32, title string, done bool) {
	newTask :=  Task{
		ID : id,
		Title : title,
		Done : done,
	}

	t.tasks = append(t.tasks, newTask)

}
func  (t *Tasklist) getall() []Task{
	return t.tasks
}

// func (t *Task) getbyid(id int32) (*Task, error) {}
// func (t *Task) update(id int32) (*Task, error){}
// func (t *Task) delete(id int32) (*Task, error){}

func main(){
	// tasks := Task{
	// 	ID: 1, Title: "learn python", Done: false,
	// 	
	// }
	tasks := Tasklist{}
	tasks.create(2, "learn golang", true)
	tasks.create(2, "learn cpp", true)
	tasks.create(2, "learn c", true)

	

	for _, t := range tasks.getall() {
		fmt.Println(t.Title)
	}


}