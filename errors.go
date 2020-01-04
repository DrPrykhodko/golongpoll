package golongpoll

import (
	"github.com/google/jsonapi"
	"net/http"
	"strconv"
)

var (
	InvalidTimeoutArg = &jsonapi.ErrorObject{
		Title:  `Invalid "timeout" argument`,
		Status: strconv.Itoa(http.StatusBadRequest),
		Code:   "invalid_timeout_arg",
	}
	InvalidCategoryArg = &jsonapi.ErrorObject{
		Title:  `Invalid "category" argument`,
		Status: strconv.Itoa(http.StatusBadRequest),
		Code:   "invalid_subscription_category_arg",
	}
	ErrorParsingLastEventTimeArg = &jsonapi.ErrorObject{
		Title:  `Invalid "last_event_time" argument`,
		Status: strconv.Itoa(http.StatusBadRequest),
		Code:   "error_parsing_last_event_time_arg",
	}
	ErrorCreatingNewSubscription = &jsonapi.ErrorObject{
		Title:  "Error creating new subscription",
		Status: strconv.Itoa(http.StatusInternalServerError),
		Code:   "error_creating_new_subscription",
	}
	JsonMarshallerFailed = &jsonapi.ErrorObject{
		Title:  "JSON marshaller failed",
		Status: strconv.Itoa(http.StatusInternalServerError),
		Code:   "json_marshaller_failed",
	}
)
