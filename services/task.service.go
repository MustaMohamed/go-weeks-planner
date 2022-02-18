package services

import (
	"context"
	"github.com/mustamohamed/weekplanner/errors"
	"github.com/mustamohamed/weekplanner/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskService interface {
	Create(task models.Task) (*models.Task, error)
	Update(task models.Task) (*models.Task, error)
	Delete(task models.Task) error
	DeleteById(id uint) error
	GetAll() ([]models.Task, error)
	Get(task models.Task) (*models.Task, error)
	GetById(id uint) (*models.Task, error)
}

type TaskService struct {
	db *gorm.DB
}

func (tskS *TaskService) Create(task models.Task) (*models.Task, error) {
	result := tskS.db.Create(&task)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewCannotInsertRecordError(result.Error)
	}
	return &task, nil
}

func (tskS *TaskService) Update(task models.Task) (*models.Task, error) {
	tsk := &models.Task{}
	tsk.ID = task.ID
	result := tskS.db.Model(tsk).Updates(task)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewCannotUpdateRecordError(result.Error)
	}
	return tsk, nil
}

func (tskS *TaskService) Delete(task models.Task) error {
	result := tskS.db.Delete(task)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(result.Error)
	}
	return nil
}

func (tskS *TaskService) DeleteById(id uint) error {
	result := tskS.db.Delete(models.Task{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(result.Error)
	}
	return nil
}

func (tskS *TaskService) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := tskS.db.Preload(clause.Associations).Find(&tasks)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return tasks, nil
}

func (tskS *TaskService) Get(task models.Task) (*models.Task, error) {
	tsk := &models.Task{}
	result := tskS.db.Where(task).Preload(clause.Associations).First(tsk)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return tsk, nil
}

func (tskS *TaskService) GetById(id uint) (*models.Task, error) {
	tsk := &models.Task{}
	result := tskS.db.Preload(clause.Associations).First(tsk, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(result.Error)
	}
	return tsk, nil
}

func NewTaskService(db *gorm.DB, ctx context.Context) ITaskService {
	return &TaskService{
		db: db.WithContext(ctx),
	}
}
