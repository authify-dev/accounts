package utils

import (
	"context"
	"foundation/domain/logger"
)

func GetFieldsOfLogger(ctx context.Context) logger.LogFields {

	entry := logger.FromContext(ctx)
	fields, ok := entry.Data["fields"].(logger.LogFields)
	if !ok {
		fields = logger.LogFields{}
	}

	return fields
}
