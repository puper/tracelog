package tracelog

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

func New(module string) *TraceLog {
	now := time.Now()
	return &TraceLog{
		logEntry: &LogEntry{
			Module:  module,
			StartAt: &now,
		},
		modules:  map[string]*TraceLog{},
		children: []*TraceLog{},
	}
}

type TraceLog struct {
	logEntry *LogEntry
	modules  map[string]*TraceLog
	children []*TraceLog
}

func (this *TraceLog) Module(module string) *TraceLog {
	if _, ok := this.modules[module]; !ok {
		this.modules[module] = New(module)
		this.children = append(this.children, this.modules[module])
	}
	return this.modules[module]
}

func (this *TraceLog) Arg(arg interface{}) *TraceLog {
	this.logEntry.Arg = arg
	return this
}

func (this *TraceLog) Reply(reply interface{}) *TraceLog {
	this.logEntry.Reply = reply
	return this
}

func (this *TraceLog) Env(vs ...interface{}) *TraceLog {
	this.logEntry.Env = append(this.logEntry.Env, vs)
	return this
}

func (this *TraceLog) Error(err interface{}) *TraceLog {
	if err != nil {
		this.logEntry.Error = err
	}
	return this
}

func (this *TraceLog) Stop(maxCost ...time.Duration) *TraceLog {
	now := time.Now()
	this.logEntry.StopAt = &now
	this.logEntry.Cost = this.logEntry.StopAt.Sub(*this.logEntry.StartAt)
	if len(maxCost) > 0 && maxCost[0] <= this.logEntry.Cost {
		this.logEntry.Slow = true
	}
	return this
}

func (this *TraceLog) Caller() *TraceLog {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		this.logEntry.Caller = fmt.Sprintf("%v:%v:%v", file, line, runtime.FuncForPC(pc).Name())
	}
	return this
}

func (this *TraceLog) Json() json.RawMessage {
	l := this.logEntries()
	b, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return b
}

type LogEntries struct {
	LogEntry *LogEntry     `json:"log,omitempty"`
	Children []*LogEntries `json:"children,omitempty"`
}

func (this *TraceLog) logEntries() *LogEntries {
	reply := &LogEntries{
		LogEntry: this.logEntry,
		Children: []*LogEntries{},
	}
	for _, v := range this.children {
		reply.Children = append(
			reply.Children,
			v.logEntries(),
		)
	}
	return reply
}
