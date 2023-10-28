package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAccountService,
	NewPostService,
	NewViewService,
	NewRelationService,
)
