package services

import (
	"context"
	"fmt"
	"github.com/mustamohamed/weekplanner/errors"
	"github.com/mustamohamed/weekplanner/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPlanService interface {
	Create(plan models.Plan) (*models.Plan, error)
	Update(plan models.Plan) (*models.Plan, error)
	Delete(plan models.Plan) error
	DeleteById(id uint) error
	GetAll() ([]models.Plan, error)
	Get(plan models.Plan) (*models.Plan, error)
	GetById(id uint) (*models.Plan, error)
}

type PlanService struct {
	db *gorm.DB
}

func (plnS *PlanService) Create(plan models.Plan) (*models.Plan, error) {
	result := plnS.db.Create(&plan)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NewCannotInsertRecordError(fmt.Errorf("cannot add item"))
	}
	return &plan, nil
}

func (plnS *PlanService) Update(plan models.Plan) (*models.Plan, error) {
	pln := &models.Plan{}
	pln.ID = plan.ID
	result := plnS.db.Model(pln).Updates(plan)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NewCannotUpdateRecordError(fmt.Errorf("cannot update item"))
	}
	return pln, nil
}

func (plnS *PlanService) Delete(plan models.Plan) error {
	result := plnS.db.Delete(plan)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(fmt.Errorf("cannot delete item"))
	}
	return nil
}

func (plnS *PlanService) DeleteById(id uint) error {
	result := plnS.db.Delete(models.Plan{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.NewCannotDeleteRecordError(fmt.Errorf("cannot delete item"))
	}
	return nil
}

func (plnS *PlanService) GetAll() ([]models.Plan, error) {
	var plans []models.Plan
	result := plnS.db.Preload(clause.Associations).Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(fmt.Errorf("items not found"))
	}
	return plans, nil
}

func (plnS *PlanService) Get(plan models.Plan) (*models.Plan, error) {
	pln := &models.Plan{}
	result := plnS.db.Where(plan).Preload(clause.Associations).First(pln)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(fmt.Errorf("item not found"))
	}
	return pln, nil
}

func (plnS *PlanService) GetById(id uint) (*models.Plan, error) {
	pln := &models.Plan{}
	result := plnS.db.Preload(clause.Associations).First(pln, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.NewNotFoundRecordError(fmt.Errorf("item not found"))
	}
	return pln, nil
}

func NewPlanService(db *gorm.DB, ctx context.Context) IPlanService {
	return &PlanService{
		db: db.WithContext(ctx),
	}
}
