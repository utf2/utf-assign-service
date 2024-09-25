package transfer

import "github.com/google/uuid"

type FormAssignRequest struct {
	FormID   uuid.UUID
	GroupIDs []uuid.UUID
}

type FormAssignResponse struct {
	Success bool
}
