// +build integration

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/m3db/m3/src/dbnode/storage/index"
	"github.com/m3db/m3/src/m3ninx/idx"
	"github.com/m3db/m3/src/x/ident"
)

func TestErrorMsg(t *testing.T) {
	testOpts := NewTestOptions(t)
	testSetup := newTestSetupWithCommitLogAndFilesystemBootstrapper(t, testOpts)
	defer testSetup.Close()

	require.NoError(t, testSetup.StartServer())
	defer require.NoError(t, testSetup.StopServer())

	var (
		storageOpts = testSetup.StorageOpts()
		nowFn       = storageOpts.ClockOptions().NowFn()
		end         = nowFn()
		start       = end.Add(-time.Hour)
		query       = index.Query{Query: idx.NewTermQuery([]byte("tag"), []byte("value"))}
	)

	session, err := testSetup.M3DBClient().DefaultSession()
	require.NoError(t, err)

	_, _, err = session.FetchTagged(
		ident.BytesID("???"), query, index.QueryOptions{StartInclusive: start, EndExclusive: end})
	require.NoError(t, err)
}