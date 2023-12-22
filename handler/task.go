package handler

import (
	"log"

	"github.com/Maxime-Hrt/got/model"
	"github.com/Maxime-Hrt/got/view/components"
	"github.com/Maxime-Hrt/got/view/tasks"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	Tasks []model.Task
}

func (h *TaskHandler) ShowTasks(c echo.Context) error {
	return render(c, tasks.ShowTask(h.Tasks))
}

func (h *TaskHandler) HandleTaskAdd(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		log.Printf("Error binding task: %v", err)
		return err
	}

	h.Tasks = append(h.Tasks, task)

	log.Println(h.Tasks)

	return render(c, components.TaskList(h.Tasks))
}
