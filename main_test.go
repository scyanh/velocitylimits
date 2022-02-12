package main

import (
	"github.com/scyanh/velocitylimits/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOutput(t *testing.T) {
	loadedJsonExpected := utils.ReadFile(utils.ExpectedOutputPath)
	loadedJsonGenerated := utils.ReadFile(utils.OutputPath)

	require.Equal(t, loadedJsonExpected, loadedJsonGenerated)
}