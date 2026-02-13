package service

import (
	"net/http"

	"cirius-go/neodb/internal/common/errors"
)

var ErrNotAMember = errors.NewI18nStatusError(http.StatusBadRequest, errors.NewI18nMsg("error:not_a_member"))
