package keeper

import (
	"github.com/nodemadic/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
