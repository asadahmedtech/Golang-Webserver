package models

import("time")

type ToDo struct{
	ID string `json:"_id"`
	Task string `json:"task"`
	CreatedAt time.Time `json:"createdat"`
	Username string `json:"username"`
}