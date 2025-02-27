//go:build !fuzzlight

package test_cases_test

import (
	"testing"

	"github.com/consensys/linea-monorepo/prover/maths/common/smartvectors"
	"github.com/consensys/linea-monorepo/prover/protocol/column"
	"github.com/consensys/linea-monorepo/prover/protocol/ifaces"
	"github.com/consensys/linea-monorepo/prover/protocol/wizard"
)

func defineFibo(build *wizard.Builder) {

	// Number of rows
	n := 1 << 3
	p1 := build.RegisterCommit(P1, n) // overshadows P

	// P(X) = P(X/w) + P(X/w^2)
	expr := ifaces.ColumnAsVariable(column.Shift(p1, -1)).
		Add(ifaces.ColumnAsVariable(column.Shift(p1, -2))).
		Sub(ifaces.ColumnAsVariable(p1))

	_ = build.GlobalConstraint(GLOBAL1, expr)
}

func proveFibo(run *wizard.ProverRuntime) {
	x := smartvectors.ForTest(1, 1, 2, 3, 5, 8, 13, 21)
	run.AssignColumn(P1, x)
}

func TestFibo(t *testing.T) {
	checkSolved(t, defineFibo, proveFibo, ALL_BUT_ILC, true)
}
