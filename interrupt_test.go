// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package interrupt

import (
	"testing"

	"github.com/maruel/ut"
)

func TestSet(t *testing.T) {
	ut.AssertEqual(t, false, IsSet())
	select {
	case <-Channel:
		t.Fatal()
	default:
	}

	HandleCtrlC()

	ut.AssertEqual(t, false, IsSet())
	select {
	case <-Channel:
		t.Fatal()
	default:
	}

	Set()

	for i := 0; i < 2; i++ {
		ut.AssertEqual(t, true, IsSet())
		x, ok := <-Channel
		ut.AssertEqual(t, true, x)
		ut.AssertEqual(t, true, ok)
	}
}
