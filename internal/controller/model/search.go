package transfer

import "github.com/google/uuid"

type AssignmentDTO struct {
	ID      uuid.UUID
	FormID  uuid.UUID
	GroupID uuid.UUID
}

type SearchAssignmentsByFormIDRequest struct {
	FormID uuid.UUID
}

type SearchAssignmentsByFormIDResponse struct {
	Assignments []AssignmentDTO
}

type SearchAssignmentsByGroupIDRequest struct {
	GroupID uuid.UUID
}

type SearchAssignmentsByGroupIDResponse struct {
	Assignments []AssignmentDTO
}
