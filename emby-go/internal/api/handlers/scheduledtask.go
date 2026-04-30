package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/scheduled"
)

// ScheduledTaskHandler handles scheduled task-related API endpoints.
type ScheduledTaskHandler struct {
	scheduledSvc *scheduled.Manager
}

// NewScheduledTaskHandler creates a new scheduled task handler.
func NewScheduledTaskHandler(scheduledSvc *scheduled.Manager) *ScheduledTaskHandler {
	return &ScheduledTaskHandler{scheduledSvc: scheduledSvc}
}

// GetAllTasks handles GET /ScheduledTasks
func (h *ScheduledTaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetRunningTasks handles GET /ScheduledTasks/Running
func (h *ScheduledTaskHandler) GetRunningTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTask handles GET /ScheduledTasks/{id}
func (h *ScheduledTaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	task := map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// ExecuteTask handles POST /ScheduledTasks/{id}/Execute
func (h *ScheduledTaskHandler) ExecuteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// CancelTask handles POST /ScheduledTasks/{id}/Cancel
func (h *ScheduledTaskHandler) CancelTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// GetTaskCount handles GET /ScheduledTasks/Count
func (h *ScheduledTaskHandler) GetTaskCount(w http.ResponseWriter, r *http.Request) {
	count := map[string]int{"count": 0}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// GetRunningTaskCount handles GET /ScheduledTasks/RunningCount
func (h *ScheduledTaskHandler) GetRunningTaskCount(w http.ResponseWriter, r *http.Request) {
	count := map[string]int{"count": 0}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}