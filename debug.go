package vngine

import (
	"fmt"
	"github.com/faiface/pixel/text"
	"time"
)

// dbg is a pointer to the DebugData singleton.
var dbg *DebugData

// DebugData contains the information on the debugging session.
type DebugData struct {
	dbgDataFormat string
	dbgData       *text.Text
	dbgLog        *text.Text
	fps           int
	fpsAcc        int
	tick          <-chan time.Time
}

// DebugLog prints out the given string with the timestamp of submission in the debug UI.
func DebugLog(log string) {
	dot := dbg.dbgLog.Dot
	lHeight := dbg.dbgLog.LineHeight
	_, _ = fmt.Fprintf(dbg.dbgLog, "%s: %s", time.Now().String(), log)
	if lHeight > 0 {
		dbg.dbgLog.Dot = dot
		_, _ = fmt.Fprintln(dbg.dbgLog)
	}
	dbg.dbgLog.Dot = dot
}
