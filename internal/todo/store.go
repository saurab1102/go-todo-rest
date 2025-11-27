package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Store struct {
	Path string
}

func NewStore(path string) *Store {
	return &Store{Path:path}
}

func (s *Store) load() ([]Todo,error){
	f,err := os.Open(s.Path)
	if errors.Is(err,os.ErrNotExist) {
		return []Todo{} , nil
	}
	if err!=nil {
		return nil,err
	}
	defer f.Close()

	var out []Todo
	if err:=json.NewDecoder(f).Decode(&out); err!=nil {
		return nil,err
	}

	return out, nil
}

func (s *Store) save(todos []Todo) error {
	f, err :=os.Create(s.Path)
	if err!=nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(todos)
}

func (s *Store) Add(text string) error {
	todos,err:= s.load();
	if err!=nil {
		return err
	}
	next:=1
	if len(todos)>0 {
		next = todos[len(todos)-1].ID+1
	}

	t:=Todo {
		ID: next,
		Text: text,
		CreatedAt: time.Now(),
	}

	todos = append(todos,t)
	return s.save(todos)
}

func (s *Store) List() ([]Todo,error) {
	return s.load();
}

func (s *Store) MarkDone(id int) error {
	todos, err := s.load()
	if err!=nil {
		return err
	}

	for i:= range todos {
		if todos[i].ID==id {
			todos[i].Done = true
			todos[i].DoneAt = time.Now()
			return s.save(todos)
		}
	}

	return errors.New("Task not found")
}

func (s *Store) Delete(id int) error {
	todos, err := s.load()
	if err!=nil {
		return err
	}

	out := make([]Todo,0,len(todos))
	found := false

	for _,t := range todos {
		if t.ID==id {
			found=true
			continue
		}
		out = append(out,t)
	}

	if !found {
		return errors.New("Task not found")
	}

	return s.save(out)
}
