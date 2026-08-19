package chorus

import (
	"time"
)

// Task defines the chorus task for multiple characters
type Task struct {
	ID          string    `json:"id"`
	TaskName    string    `json:"taskName"`
	PhontList   *[]string `json:"phontList"`
	ResultFile  *string   `json:"resultFile"`
	CreatedTime time.Time `json:"createdTime"`
}
