package engine

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ankaa-labs/ankaa/utils/pointer"
)

func TestCatcherValidate(t *testing.T) {
	type testCase struct {
		desp string
		c    *Catcher
		err  string
	}
	testCases := []testCase{
		{
			desp: "normal one",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp: "normal wildcard",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesAll},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp: "normal multi",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, "tt"},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: ``,
		},
		{
			desp: "empty ErrorEquals",
			c: &Catcher{
				ErrorEquals: []string{},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: `ErrorEquals cannot be empty`,
		},
		{
			desp: "empty ErrorEqual item",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, ""},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: `ErrorEqual cannot be empty`,
		},
		{
			desp: "not pre defined",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed + "1", ErrorCodeStatesHeartbeatTimeout, "tt"},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: `ErrorEqual is not pre defined`,
		},
		{
			desp: "multi with ALL",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesAll, "tt"},
				Next:        "n",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: `States\.ALL must be be only element in ErrorEquals`,
		},
		{
			desp: "empty next",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, "tt"},
				Next:        "Next cannot be empty",
				ResultPath:  pointer.StringPtr("r"),
			},
			err: ``,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desp, func(t *testing.T) {
			g := NewWithT(t)
			err := tc.c.Validate()

			if tc.err != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(MatchRegexp(tc.err))
				return
			}

			g.Expect(err).ToNot(HaveOccurred())
		})
	}
}

func TestCatcherIsAnyErrorWildcard(t *testing.T) {
	type testCase struct {
		desp   string
		c      *Catcher
		expect bool
	}
	testCases := []testCase{
		{
			desp: "not only 1 ErrorEquals",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesAll},
			},
			expect: false,
		},
		{
			desp: "only 1 but not ALL",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesBranchFailed},
			},
			expect: false,
		},
		{
			desp: "only 1 and is ALL",
			c: &Catcher{
				ErrorEquals: []string{ErrorCodeStatesAll},
			},
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desp, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(tc.c.IsAnyErrorWildcard()).To(Equal(tc.expect))
		})
	}
}

func TestCatchersValidate(t *testing.T) {
	type testCase struct {
		desp string
		c    Catchers
		err  string
	}
	testCases := []testCase{
		{
			desp: "empty catchers",
			c:    Catchers([]Catcher{}),
			err:  ``,
		},
		{
			desp: "multi catchers",
			c: Catchers([]Catcher{
				{
					ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, "ERR"},
					Next:        "n",
				},
				{
					ErrorEquals: []string{ErrorCodeStatesAll},
					Next:        "n",
				},
			}),
			err: ``,
		},
		{
			desp: "catcher valid failed",
			c: Catchers([]Catcher{
				{
					ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, "ERR"},
					Next:        "",
				},
				{
					ErrorEquals: []string{ErrorCodeStatesAll},
					Next:        "n",
				},
			}),
			err: `Next cannot be empty`,
		},
		{
			desp: "ALL is not last",
			c: Catchers([]Catcher{
				{
					ErrorEquals: []string{ErrorCodeStatesAll},
					Next:        "n",
				},
				{
					ErrorEquals: []string{ErrorCodeStatesBranchFailed, ErrorCodeStatesHeartbeatTimeout, "ERR"},
					Next:        "n",
				},
			}),
			err: `States\.ALL must be the last catcher`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desp, func(t *testing.T) {
			g := NewWithT(t)
			err := tc.c.Validate()

			if tc.err != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(MatchRegexp(tc.err))
				return
			}

			g.Expect(err).ToNot(HaveOccurred())
		})
	}
}
