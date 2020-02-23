// Copyright © Go Opus Authors (see AUTHORS file)
//
// License for use of this code is detailed in the LICENSE file

package opus

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if ver := Version(); !strings.HasPrefix(ver, "libopus") {
		t.Errorf("Unexpected linked libopus version: " + ver)
	}
}

func TestOpusErrstr(t *testing.T) {
	// I scooped this -1 up from opus_defines.h, it's OPUS_BAD_ARG. Not pretty,
	// but it's better than not testing at all. Again, accessing #defines from
	// CGO is not possible.
	if ErrBadArg.Error() != "opus: invalid argument" {
		t.Errorf("Expected \"invalid argument\" error message for error code %d: %v",
			ErrBadArg, ErrBadArg)
	}
}
