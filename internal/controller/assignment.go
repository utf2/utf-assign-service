package controller

import transfer "github.com/utf2/utf-assign-service/internal/controller/model"

type AssignmentControllerAPI interface {
	AssignForm(transfer.FormAssignRequest) transfer.FormAssignResponse
	SearchByFormID(transfer.SearchAssignmentsByFormIDRequest) transfer.SearchAssignmentsByFormIDResponse
	SearchByGroupID(transfer.SearchAssignmentsByGroupIDRequest) transfer.SearchAssignmentsByGroupIDResponse
}
