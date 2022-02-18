package services

import (
	"context"
	"github.com/mustamohamed/weekplanner/errors"
	"github.com/mustamohamed/weekplanner/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITodoService interface {
	Create(todo models.Todo) (*models.Todo, error)
	Update(todo models.Todo) (*models.Todo, error)
	Delete(todo models.Todo) error
	DeleteById(id uint) error
	GetAll() ([]models.Todo, error)
	Get(todo models.Todo) (*models.Todo, error)
	GetById(id uint) (*models.Todo, error)
}

type TodoService struct {
	db *gorm.DB
}

func (tskS *TodoService) Create(todo models.Todo) (*models.Todo, error) {
	result := tskS.db.Create(&todo)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewCannotInsertRecordError(result.Error)
	}
	return &todo, nil
}

func (tskS *TodoService) Update(todo models.Todo) (*models.Todo, error) {
	tdo := &models.Todo{}
	tdo.ID = todo.ID
	result := tskS.db.Model(tdo).Updates(todo)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewCannotUpdateRecordError(result.Error)
	}
	return tdo, nil
}

func (tskS *TodoService) Delete(todo models.Todo) error {
	result := tskS.db.Delete(todo)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(result.Error)
	}
	return nil
}

func (tskS *TodoService) DeleteById(id uint) error {
	result := tskS.db.Delete(models.Todo{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(result.Error)
	}
	return nil
}

func (tskS *TodoService) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := tskS.db.Preload(clause.Associations).Find(&todos)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return todos, nil
}

func (tskS *TodoService) Get(todo models.Todo) (*models.Todo, error) {
	tdo := &models.Todo{}
	result := tskS.db.Where(todo).Preload(clause.Associations).First(tdo)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return tdo, nil
}

func (tskS *TodoService) GetById(id uint) (*models.Todo, error) {
	tdo := &models.Todo{}
	result := tskS.db.Preload(clause.Associations).First(tdo, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return tdo, nil
}

func NewTodoService(db *gorm.DB, ctx context.Context) ITodoService {
	return &TodoService{
		db: db.WithContext(ctx),
	}
}
