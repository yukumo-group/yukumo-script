package chorus

import (
	"time"

	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// Task defines the chorus task for multiple characters
type Task struct {
	ID           string            `json:"id"`
	TaskName     string            `json:"taskName"`
	PhontList    *[]string         `json:"phontList"`
	ResultFile   *string           `json:"resultFile"`
	CreatedTime  time.Time         `json:"createdTime"`
	EditTime     time.Time         `json:"editTime"`
	Text         string            `json:"text"`
	TaskLanguage language.Language `json:"taskLanguage"`
}
