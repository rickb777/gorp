// Copyright 2012 James Cooper. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package gorp provides a simple way to marshal Go structs to and from
// SQL databases.  It uses the database/sql package, and should work with any
// compliant database/sql driver.
//
// Source code and project home:
// https://github.com/go-gorp/gorp

package gorp

import (
	"fmt"
	"strings"
)

type GorpLogger interface {
	Printf(format string, v ...interface{})
}

// TraceOn turns on SQL statement logging for this DbMap.  After this is
// called, all SQL statements will be sent to the logger.  If prefix is
// a non-empty string, it will be written to the front of all logged
// strings, which can aid in filtering log lines.
//
// Use TraceOn if you want to spy on the SQL statements that gorp
// generates.
//
// Note that the base log.Logger type satisfies GorpLogger, but adapters can
// easily be written for other logging packages (e.g., the golang-sanctioned
// glog framework).
func (m *DbMap) TraceOn(prefix, suffix string, logger GorpLogger) {
	m.logger = logger

	if strings.TrimSpace(prefix) == "" {
		m.logPrefix = prefix
	} else if !strings.HasSuffix(prefix, " ") {
		m.logPrefix = fmt.Sprintf("%s ", prefix)
	}

	m.logSuffix = suffix
}

// TraceOff turns off tracing. It is idempotent.
func (m *DbMap) TraceOff() {
	m.logger = nil
	m.logPrefix = ""
}
