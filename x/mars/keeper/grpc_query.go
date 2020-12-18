package keeper

import (
	"github.com/ilgooz/mars/x/mars/types"
)

var _ types.QueryServer = Keeper{}
