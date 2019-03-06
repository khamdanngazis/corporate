package services

import (
	"context"

	dt "../datastruct"
	lib "../lib"
	"../logging"
)

// CorporateCreateTicketServices provides operations for endpoint
type CorporateCreateTicketServices interface {
	CorporateCreateTicket(context.Context, dt.CorporateCreateTicketJSONRequest, lib.DbConnection) dt.CorporateCreateTicketJSONResponse
}

// CorporateCreateTicketService is a concrete implementation of QueueServices
type CorporateCreateTicketService struct{}

// CorporateCreateTicket service
func (CorporateCreateTicketService) CorporateCreateTicket(ctx context.Context, req dt.CorporateCreateTicketJSONRequest, dbConn lib.DbConnection) dt.CorporateCreateTicketJSONResponse {

	logging.Logf(" Request ", req)

	return dt.CorporateCreateTicketJSONResponse{
		ResponseCode: dt.ErrSuccess,
		ResponseDesc: dt.DescSuccess,
	}

}
