//////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// LICENCE
// VNgine - a simple robust visual novel engine.
// Copyright© 2020 Kornel Domeradzki
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
package vngine

import (
	"fmt"
	"github.com/faiface/pixel/text"
	"runtime"
	"strings"
	"time"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains all of the necessary debug data and utilities.

// dbg is a pointer to the DebugData singleton.
var dbg *DebugData

// DebugData contains the information on the debugging session.
type DebugData struct {
	dbgDataFormat string
	dbgData       *text.Text
	dbgLog        *text.Text
	logLines      []string
	fps           int
	fpsAcc        int
	tick          <-chan time.Time
}

// DebugLog prints out the given string with the timestamp of submission in the debug UI.
func DebugLog(log string) {
	if dbg == nil {
		return
	}
	if len(dbg.logLines) == 20 {
		dbg.logLines = dbg.logLines[:19]
	}
	dbg.logLines = append([]string{time.Now().String() + ": " + log}, dbg.logLines...)
	dbg.dbgLog.Clear()
	for _, line := range dbg.logLines {
		_, _ = fmt.Fprintln(dbg.dbgLog, line)
	}
}

// GetMemUsage returns a string with all the memory usage data.
func GetMemUsage() string {
	bToMb := func(b uint64) uint64 { return b / 1024 / 1024 }
	var m runtime.MemStats
	var sb strings.Builder
	runtime.ReadMemStats(&m)

	_, _ = fmt.Fprintf(&sb, "Alloc = %v MiB\r\nTotal Alloc = %v MiB\r\nSys = %v MiB\r\nNumGC = %v\r\n",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
	return sb.String()
}
