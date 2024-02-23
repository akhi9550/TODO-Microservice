package services

import (
	"context"

	"github.com/akhi9550/todo-service/pkg/db"
	"github.com/akhi9550/todo-service/pkg/models"
	"github.com/akhi9550/todo-service/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedAuthServiceServer
}

func (s *Server) CreateTodo(ctx context.Context, req *pb.AddTodoRequest) (*pb.AddTodoResponse, error) {
	var Data models.Data
	err := s.H.DB.Raw("INSERT INTO data (title,description) VALUES($1,$2) RETURNING id, title, description", req.Title, req.Description).Scan(&Data).Error
	if err != nil {
		return &pb.AddTodoResponse{}, err
	}
	return &pb.AddTodoResponse{
		ID:          Data.ID,
		Title:       Data.Title,
		Description: Data.Description,
	}, nil
}
func (s *Server) ListTodo(ctx context.Context, details *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	page := int(details.Page)
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 5
	var todoData []models.Data
	result := s.H.DB.Raw("SELECT id, title, description FROM data LIMIT $1 OFFSET $2", 5, offset).Scan(&todoData)
	if result.Error != nil {
		return &pb.ListTodoResponse{
			ListTodos: []*pb.TodoDetails{},
		}, result.Error
	}
	var TodoDetails pb.ListTodoResponse
	for _, v := range todoData {
		TodoDetails.ListTodos = append(TodoDetails.ListTodos, &pb.TodoDetails{
			ID:          int64(v.ID),
			Title:       v.Title,
			Description: v.Description,
		})
	}
	return &TodoDetails, nil
}
func (s *Server) GetTodoByID(ctx context.Context, details *pb.TodoIDRequest) (*pb.TodoItemResponse, error) {
	var todoData models.Data
	err := s.H.DB.Raw("SELECT id, title, description FROM data WHERE id = $1", details.ID).Scan(&todoData).Error
	if err != nil {
		return &pb.TodoItemResponse{},
			err
	}
	return &pb.TodoItemResponse{
		ID:          todoData.ID,
		Title:       todoData.Title,
		Description: todoData.Description,
	}, nil
}
func (s *Server) DeleteTodo(ctx context.Context, details *pb.TodoIDRequest) (*pb.DeleteTodoResponse, error) {
	err := s.H.DB.Exec(" DELETE FROM data WHERE id = $1", details.ID).Error
	if err != nil {
		return &pb.DeleteTodoResponse{},
			err
	}
	return &pb.DeleteTodoResponse{
		Status: "Delete Successfully",
	}, nil
}
func (s *Server) UpdateTodo(ctx context.Context, details *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	err := s.H.DB.Exec("UPDATE data SET title = $1 WHERE id = $2", details.Title, details.ID).Error
	if err != nil {
		return &pb.UpdateTodoResponse{},
			err
	}
	err = s.H.DB.Exec("UPDATE data SET description = $1 WHERE id = $2", details.Description, details.ID).Error
	if err != nil {
		return &pb.UpdateTodoResponse{},
			err
	}
	var todoData models.Data
	if err := s.H.DB.Raw("SELECT id, title, description FROM data WHERE id = $1", details.ID).Scan(&todoData).Error; err != nil {
		return &pb.UpdateTodoResponse{}, err
	}
	return &pb.UpdateTodoResponse{
		ID:          todoData.ID,
		Title:       todoData.Title,
		Description: todoData.Description,
	}, nil
}
