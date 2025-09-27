package utils

import (
	"encoding/json"
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/utils/cerrs"
	"log"
)

type Alert struct {
	Message string `json:"message,omitempty"`
	Title   string `json:"title,omitempty"`
	Icon    string `json:"icon,omitempty"`
	Code    uint   `json:"code,omitempty"`
	Scope   string `json:"scope,omitempty"`
}

type Response[R any] struct {
	Error      cerrs.CustomErrorInterface `json:"error,omitempty"`
	StatusCode int                        `json:"status_code" default:"200"`

	Data    R                            `json:"data,omitempty"`
	Results []R                          `json:"results,omitempty"`
	Alert   *Alert                       `json:"alert,omitempty"`
	TraceID string                       `json:"trace_id,omitempty"`
	Success bool                         `json:"success" default:"true"`
	Errors  []cerrs.CustomErrorInterface `json:"errors,omitempty"`
}

func (r Response[R]) ToMapWithCustomContext(ctx *customctx.CustomContext) map[string]interface{} {

	if ctx == nil {
		return r.ToMap()
	}

	fields := ctx.Context().Value("fields").(logger.LogFields)

	r.TraceID = fields.TraceID

	res := r.ToMap()

	if len(ctx.Errors()) > 0 {

		if logger.LoggerConfig.ENVIRONMENT != "production" {

			res["errors"] = ctx.Errors()
			delete(res, "data")
		}

	}

	if len(r.Results) > 0 {
		delete(res, "data")
	}

	return res
}

func (r Response[R]) ToMap() map[string]interface{} {

	if r.StatusCode == 0 {
		r.StatusCode = 200
	}

	data, err := json.Marshal(r)
	if err != nil {
		log.Printf("error marshaling OfferEntity: %v", err)
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Printf("error unmarshaling to map: %v", err)
		return nil
	}

	return result
}
