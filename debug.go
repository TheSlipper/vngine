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
	logLines []string
	fps           int
	fpsAcc        int
	tick          <-chan time.Time
}

// DebugLog prints out the given string with the timestamp of submission in the debug UI.
func DebugLog(log string) {
	//dbg.dbgLog.Dot = pixel.V(0, 0)
	//fmt.Fprintln(dbg.dbgLog)
	//fmt.Fprintf(dbg.dbgLog, "%s: %s", time.Now().String(), log)
	if len(dbg.logLines) == 20 {
		dbg.logLines = dbg.logLines[:19]
	}
	dbg.logLines = append([]string{time.Now().String() + ": " + log}, dbg.logLines...)
	dbg.dbgLog.Clear()
	for _, line := range dbg.logLines {
		//dbg.dbgLog.Dot.X -= dbg.dbgLog.BoundsOf(line).W() / 2
		_, _ = fmt.Fprintln(dbg.dbgLog, line)
	}
}
