// SPDX-FileCopyrightText: 2022 Free Mobile
// SPDX-License-Identifier: AGPL-3.0-only

//go:build !release

package geoip

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"akvorado/common/daemon"
	"akvorado/common/helpers"
	"akvorado/common/reporter"
)

// NewMock creates a GeoIP component usable for testing. It is already
// started. It panics if there is an issue. Data of both databases are
// available here:
//   - https://github.com/maxmind/MaxMind-DB/blob/main/source-data/GeoLite2-ASN-Test.json
//   - https://github.com/maxmind/MaxMind-DB/blob/main/source-data/GeoLite2-Country-Test.json
func NewMock(t *testing.T, r *reporter.Reporter) *Component {
	t.Helper()
	config := DefaultConfiguration()
	_, src, _, _ := runtime.Caller(0)
	config.GeoDatabase = filepath.Join(path.Dir(src), "testdata", "GeoLite2-Country-Test.mmdb")
	config.ASNDatabase = filepath.Join(path.Dir(src), "testdata", "GeoLite2-ASN-Test.mmdb")
	c, err := New(r, config, Dependencies{Daemon: daemon.NewMock(t)})
	if err != nil {
		t.Fatalf("New() error:\n%+s", err)
	}
	helpers.StartStop(t, c)
	return c
}
