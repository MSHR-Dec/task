package application

import (
	"database/sql"
	"time"

	"github.com/MSHR-Dec/task/go_task/application/dto"
	"github.com/MSHR-Dec/task/go_task/domain/model"
	"github.com/MSHR-Dec/task/go_task/domain/repository"
)

type TaskApplication interface {
	Add(input dto.TaskAddInput) (dto.TaskAddOutput, error)
	Update(input dto.TaskEditInput) error
	ListByUserID(input dto.TaskListInput) (dto.TaskListOutput, error)
}

type TaskInteractor struct {
	taskRepository repository.TaskRepository
}

func NewTaskInteractor(taskRepository repository.TaskRepository) TaskInteractor {
	return TaskInteractor{
		taskRepository: taskRepository,
	}
}

func (i TaskInteractor) Add(input dto.TaskAddInput) (dto.TaskAddOutput, error) {
	now := time.Now()
	task, err := model.NewTask(input.Name, input.StartAt, input.EndAt, uint(input.UserID), now)
	if err != nil {
		return dto.TaskAddOutput{}, err
	}

	id, err := i.taskRepository.Save(task)
	if err != nil {
		return dto.TaskAddOutput{}, err
	}

	return dto.TaskAddOutput{
		ID: id,
	}, nil
}

func (i TaskInteractor) Update(input dto.TaskEditInput) error {
	task, err := i.taskRepository.FindByID(uint(input.ID))
	if err != nil {
		return err
	}

	if input.Name != "" {
		name, err := model.NewTaskName(input.Name)
		if err != nil {
			return err
		}

		task.Name = name
	}

	if input.Status != "" {
		task.Status = model.NewTaskStatus(input.Status)
	}

	if input.StartAt != "" {
		task.StartAt = model.NewScheduleAt(input.StartAt)
	}

	if input.EndAt != "" {
		task.EndAt = model.NewScheduleAt(input.EndAt)
	}

	now := time.Now()
	task.ModifiedAt = sql.NullTime{
		Time:  now,
		Valid: true,
	}

	if err = i.taskRepository.Update(task); err != nil {
		return err
	}

	return nil
}

func (i TaskInteractor) ListByUserID(input dto.TaskListInput) (dto.TaskListOutput, error) {
	tasks, err := i.taskRepository.ListByUserID(uint(input.UserID))
	if err != nil {
		return dto.TaskListOutput{}, err
	}

	output := make([]dto.TaskOutput, len(tasks))
	for i, task := range tasks {
		output[i] = dto.TaskOutput{
			ID:         task.ID,
			Name:       task.Name.String(),
			Status:     task.Status.String(),
			StartAt:    task.StartAt.Time(),
			EndAt:      task.EndAt.Time(),
			CreatedAt:  task.CreatedAt,
			ModifiedAt: task.ModifiedAt.Time,
		}
	}

	return dto.TaskListOutput{
		Tasks: output,
	}, nil
}
