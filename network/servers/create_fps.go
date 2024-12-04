package servers

import (
	pb "awesomeProject/protoFiles/files"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Invoice(_ context.Context, req *pb.GenerateInvoiceRequestFPSPostman) (*pb.GenerateInvoiceResponseFPSPostman, error) {
	println(req.OrderType)
	println(req.OrderType.Type())
	println(req.OrderType.Enum())
	println(req.OrderType.Descriptor())
	println(req.OrderType.EnumDescriptor())
	println(req.OrderType.String())
	println(req.OrderType.Type())
	return nil, status.Error(codes.Unimplemented, "")
	//token := awesomeProject.GetToken()
	//if token!="" {
	//
	//}
	//if err := validateInvoiceFPSRequest(req); err != nil {
	//
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}
	//return &pb.GenerateInvoiceResponseFPSPostman{}, nil
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
