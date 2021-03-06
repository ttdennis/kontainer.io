package user

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kontainerooo/kontainer.ooo/pkg/user/pb"
	oldcontext "golang.org/x/net/context"
)

// TODO: fix config proto message, restrict password/salt properties

// MakeGRPCServer makes a set of endpoints available as a gRPC UserServiceServer
func MakeGRPCServer(ctx context.Context, endpoints Endpoints, logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		createUser: grpctransport.NewServer(
			endpoints.CreateUserEndpoint,
			DecodeGRPCCreateUserRequest,
			EncodeGRPCCreateUserResponse,
			options...,
		),
		editUser: grpctransport.NewServer(
			endpoints.EditUserEndpoint,
			DecodeGRPCEditUserRequest,
			EncodeGRPCEditUserResponse,
			options...,
		),
		changeUsername: grpctransport.NewServer(
			endpoints.ChangeUsernameEndpoint,
			DecodeGRPCChangeUsernameRequest,
			EncodeGRPCChangeUsernameResponse,
			options...,
		),
		deleteUser: grpctransport.NewServer(
			endpoints.DeleteUserEndpoint,
			DecodeGRPCDeleteUserRequest,
			EncodeGRPCDeleteUserResponse,
			options...,
		),
		resetPassword: grpctransport.NewServer(
			endpoints.ResetPasswordEndpoint,
			DecodeGRPCResetPasswordRequest,
			EncodeGRPCResetPasswordResponse,
			options...,
		),
		getUser: grpctransport.NewServer(
			endpoints.GetUserEndpoint,
			DecodeGRPCGetUserRequest,
			EncodeGRPCGetUserResponse,
			options...,
		),
		checkLoginCredentials: grpctransport.NewServer(
			endpoints.CheckLoginCredentialsEndpoint,
			DecodeGRPCCheckLoginCredentialsRequest,
			EncodeGRPCCheckLoginCredentialsResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createUser            grpctransport.Handler
	editUser              grpctransport.Handler
	changeUsername        grpctransport.Handler
	deleteUser            grpctransport.Handler
	resetPassword         grpctransport.Handler
	getUser               grpctransport.Handler
	checkLoginCredentials grpctransport.Handler
}

func (s *grpcServer) CreateUser(ctx oldcontext.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, res, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateUserResponse), nil
}

func (s *grpcServer) EditUser(ctx oldcontext.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	_, res, err := s.editUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.EditUserResponse), nil
}

func (s *grpcServer) ChangeUsername(ctx oldcontext.Context, req *pb.ChangeUsernameRequest) (*pb.ChangeUsernameResponse, error) {
	_, res, err := s.changeUsername.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ChangeUsernameResponse), nil
}

func (s *grpcServer) DeleteUser(ctx oldcontext.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, res, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DeleteUserResponse), nil
}

func (s *grpcServer) ResetPassword(ctx oldcontext.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	_, res, err := s.resetPassword.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ResetPasswordResponse), nil
}

func (s *grpcServer) GetUser(ctx oldcontext.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, res, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUserResponse), nil
}

func (s *grpcServer) CheckLoginCredentials(ctx oldcontext.Context, req *pb.CheckLoginCredentialsRequest) (*pb.CheckLoginCredentialsResponse, error) {
	_, res, err := s.checkLoginCredentials.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CheckLoginCredentialsResponse), nil
}

func convertPbAddress(pb *pb.Address) *Address {
	return &Address{
		ID:         uint(pb.ID),
		Postcode:   pb.Postcode,
		City:       pb.City,
		Country:    pb.Country,
		Street:     pb.Street,
		Houseno:    int(pb.Houseno),
		Additional: pb.Additional,
	}
}

// ConvertPbConfig converts a protbuf config into a user.Config
func ConvertPbConfig(pb *pb.Config) (*Config, *Address) {
	cfg := &Config{
		Admin:     pb.Admin,
		Email:     pb.Email,
		Password:  pb.Password,
		Salt:      pb.Salt,
		Image:     pb.Image,
		Phone:     pb.Phone,
		AddressID: uint(pb.AddressID),
	}

	adr := convertPbAddress(pb.Address)
	cfg.Address = *adr

	return cfg, adr
}

// ConvertConfig converts a user-package-domain config struct into its protobuf message equivalent
func ConvertConfig(cfg *Config, password bool) *pb.Config {
	var hash, salt string
	if password {
		hash = cfg.Password
		salt = cfg.Salt
	}

	return &pb.Config{
		Admin:     cfg.Admin,
		Email:     cfg.Email,
		Password:  hash,
		Salt:      salt,
		Phone:     cfg.Phone,
		Image:     cfg.Image,
		AddressID: uint32(cfg.AddressID),
		Address: &pb.Address{
			ID:         uint32(cfg.Address.ID),
			Postcode:   cfg.Address.Postcode,
			City:       cfg.Address.City,
			Country:    cfg.Address.Country,
			Street:     cfg.Address.Street,
			Houseno:    int32(cfg.Address.Houseno),
			Additional: cfg.Address.Additional,
		},
	}
}

func convertUser(usr *User) *pb.User {
	return &pb.User{
		ID:       uint32(usr.ID),
		Username: usr.Username,
		Name:     usr.Name,
		Surname:  usr.Surname,
		Config:   ConvertConfig(&usr.Config, true),
	}
}

// DecodeGRPCCreateUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateUser request to a user-domain createUser request.
func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateUserRequest)

	cfg, adr := ConvertPbConfig(req.Config)

	return CreateUserRequest{
		Username: req.Username,
		Cfg:      cfg,
		Adr:      adr,
	}, nil
}

// DecodeGRPCEditUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC EditUser request to a user-domain editUser request.
func DecodeGRPCEditUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EditUserRequest)
	cfg, _ := ConvertPbConfig(req.Config)
	return EditUserRequest{
		ID:  uint(req.ID),
		Cfg: cfg,
	}, nil
}

// DecodeGRPCChangeUsernameRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ChangeUsername request to a user-domain changeUsername request.
func DecodeGRPCChangeUsernameRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ChangeUsernameRequest)
	return ChangeUsernameRequest{
		ID:       uint(req.ID),
		Username: req.Username,
	}, nil
}

// DecodeGRPCDeleteUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC DeleteUser request to a user-domain deleteUser request.
func DecodeGRPCDeleteUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteUserRequest)
	return DeleteUserRequest{
		ID: uint(req.ID),
	}, nil
}

// DecodeGRPCResetPasswordRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ResetPassword request to a user-domain resetPassword request.
func DecodeGRPCResetPasswordRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ResetPasswordRequest)
	return ResetPasswordRequest{
		Email: req.Email,
	}, nil
}

// DecodeGRPCGetUserRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetUser request to a user-domain getUser request.
func DecodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetUserRequest)
	return GetUserRequest{
		ID: uint(req.ID),
	}, nil
}

// DecodeGRPCCheckLoginCredentialsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CheckLoginCredentialsRequest request to a user-domain CheckLoginCredentials request.
func DecodeGRPCCheckLoginCredentialsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CheckLoginCredentialsRequest)
	return CheckLoginCredentialsRequest{
		Username: req.Username,
		Password: req.Password,
	}, nil
}

// EncodeGRPCCreateUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createUser response to a gRPC CreateUser response.
func EncodeGRPCCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateUserResponse)
	gRPCRes := &pb.CreateUserResponse{
		ID: uint32(res.ID),
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCEditUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain editUser response to a gRPC EditUser response.
func EncodeGRPCEditUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(EditUserResponse)
	gRPCRes := &pb.EditUserResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCChangeUsernameResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain changeUsername response to a gRPC ChangeUsername response.
func EncodeGRPCChangeUsernameResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ChangeUsernameResponse)
	gRPCRes := &pb.ChangeUsernameResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCDeleteUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain deleteUser response to a gRPC DeleteUser response.
func EncodeGRPCDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(DeleteUserResponse)
	gRPCRes := &pb.DeleteUserResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCResetPasswordResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createUser response to a gRPC CreateUser response.
func EncodeGRPCResetPasswordResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(ResetPasswordResponse)
	gRPCRes := &pb.ResetPasswordResponse{}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCGetUserResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getUser response to a gRPC GetUser response.
func EncodeGRPCGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(GetUserResponse)
	gRPCRes := &pb.GetUserResponse{}
	if res.User != nil {
		gRPCRes.User = convertUser(res.User)
	}
	if res.Error != nil {
		gRPCRes.Error = res.Error.Error()
	}
	return gRPCRes, nil
}

// EncodeGRPCCheckLoginCredentialsResponse is a transport/grpc.EncodeRequestFunc that converts a
// user-domain CheckLoginCredentials response to a gRPC CheckLoginCredentials response.
func EncodeGRPCCheckLoginCredentialsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(CheckLoginCredentialsResponse)
	gRPCRes := &pb.CheckLoginCredentialsResponse{
		ID: uint32(res.ID),
	}
	return gRPCRes, nil
}
