package domain

import "time"

type RegisterApiReq struct {
	TimeIntervalCheck int64     `json:"time_interval_check" bson:"time_interval_check"`
	Url               string    `json:"url" bson:"url"`
	Method            string    `json:"method" bson:"method"`
	Headers           []Headers `json:"headers" bson:"headers"`
	Body              string    `json:"body" bson:"body"`
}

type CheckedApi struct {
	TimeIntervalCheck int64     `json:"time_interval_check" bson:"time_interval_check"`
	Url               string    `json:"url" bson:"url"`
	Method            string    `json:"method" bson:"method"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}

type Headers struct {
	Authorization string `json:"authorization" bson:"authorization"`
	XAccessToken  string `json:"x_access_token" bson:"XAccessToken"`
}

type Api struct {
	Url               string `json:"url" bson:"url"`
	TimeIntervalCheck int    `json:"time_interval_check" bson:"time_interval_check"`
	Method            string `json:"method" bson:"method"`
}

type Response struct {
	Status     string `json:"status"`
	StatucCode int    `json:"statusCode"`
}

type WebHookMessage struct {
	Message string `json:"message"`
}
