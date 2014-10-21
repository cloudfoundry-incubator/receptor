package receptor

import "github.com/cloudfoundry-incubator/runtime-schema/models"

type CreateTaskRequest struct {
	TaskGuid   string                  `json:"task_guid"`
	Domain     string                  `json:"domain"`
	Actions    []models.ExecutorAction `json:"actions"`
	Stack      string                  `json:"stack"`
	MemoryMB   int                     `json:"memory_mb"`
	DiskMB     int                     `json:"disk_mb"`
	CpuPercent float64                 `json:"cpu_percent"`
	Log        models.LogConfig        `json:"log"`
	ResultFile string                  `json:"result_file"`
	Annotation string                  `json:"annotation,omitempty"`
}

type TaskResponse CreateTaskRequest
