package ppmshipment

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type PPMShipmentSuite struct {
	*testingsuite.PopTestSuite
}

func TestPPMShipmentServiceSuite(t *testing.T) {
	ts := &PPMShipmentSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage(), testingsuite.WithPerTestTransaction()),
	}
	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}
