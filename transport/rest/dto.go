package rest

import "errors"

type fetchResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ItemPatchRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (update *ItemPatchRequest) Validate() error {
	if update.Title == "" && update.Description == "" {
		return errors.New("update structure has no values")
	}
	return nil
}
