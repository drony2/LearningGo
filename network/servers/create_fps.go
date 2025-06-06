package servers

import (
	pb "awesomeProject/protoFiles/files"
	"context"
	"errors"
)

func (s *Server) InvoiceFPS(_ context.Context, req *pb.GenerateInvoiceRequestFPSPostman) *pb.GenerateInvoiceResponseFPSPostman {
	return &pb.GenerateInvoiceResponseFPSPostman{
		Uuid:           "1",
		Card:           "1",
		Sum:            1,
		CardHolderName: "12",
		BankName:       "sd",
	}
}

func validateInvoiceFPSRequest(in *pb.GenerateInvoiceRequestFPSPostman) error {
	if in.Sum < 1000 {
		return errors.New("sum must be greater than 1000")
	}
	//if in.OrderType. {
	//
	//}
	return nil
}
