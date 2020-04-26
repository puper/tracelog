package tracelog

import "time"

type LogEntry struct {
	Module  string        `json:"module,omitempty"`
	Caller  string        `json:"caller,omitempty"`
	Arg     interface{}   `json:"arg,omitempty"`
	Reply   interface{}   `json:"reply,omitempty"`
	Error   interface{}   `json:"error,omitempty"`
	Env     []interface{} `json:"env,omitempty"`
	StartAt *time.Time    `json:"startAt,omitempty"`
	StopAt  *time.Time    `json:"stopAt,omitempty"`
	Slow    bool          `json:"slow,omitempty"`
	Cost    time.Duration `json:"cost,omitempty"`
}
