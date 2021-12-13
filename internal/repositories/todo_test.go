package repositories

import (
	"reflect"
	"testing"

	"github.com/bagastri07/Todolist-gRPC-Gateway/model"
	todo "github.com/bagastri07/Todolist-gRPC-Gateway/protobuf/go"
)

func TestTodoRepository_ReadTodoByID(t *testing.T) {
	type args struct {
		TodoID int32
	}
	tests := []struct {
		name    string
		r       *TodoRepository
		args    args
		want    *model.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadTodoByID(tt.args.TodoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.ReadTodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.ReadTodoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_CreateToDo(t *testing.T) {
	type args struct {
		req *todo.CreateRequest
	}
	tests := []struct {
		name    string
		r       *TodoRepository
		args    args
		want    *model.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateToDo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.CreateToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.CreateToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_UpdateToDo(t *testing.T) {
	type args struct {
		req *todo.UpdateRequest
	}
	tests := []struct {
		name    string
		r       *TodoRepository
		args    args
		want    *todo.UpdateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UpdateToDo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.UpdateToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.UpdateToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_ReadAll(t *testing.T) {
	type args struct {
		req *todo.ReadAllRequest
	}
	tests := []struct {
		name    string
		r       *TodoRepository
		args    args
		want    *todo.ReadAllResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadAll(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	type args struct {
		req *todo.DeleteRequest
	}
	tests := []struct {
		name    string
		r       *TodoRepository
		args    args
		want    *todo.DeleteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.DeleteTodo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
