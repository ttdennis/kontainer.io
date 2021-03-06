package container

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/kontainerooo/kontainer.ooo/pkg/container/pb"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

// MakeWebsocketService makes a set of container Endpoints available as a websocket Service
func MakeWebsocketService(endpoints Endpoints) *ws.ServiceDescription {
	service, _ := ws.NewServiceDescription("containerService", ws.ProtoIDFromString("CNT"))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"CreateContainer",
		ws.ProtoIDFromString("CRT"),
		endpoints.CreateContainerEndpoint,
		DecodeWSCreateContainerRequest,
		EncodeGRPCCreateContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveContainer",
		ws.ProtoIDFromString("REM"),
		endpoints.RemoveContainerEndpoint,
		DecodeWSRemoveContainerRequest,
		EncodeGRPCRemoveContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"Instances",
		ws.ProtoIDFromString("ALL"),
		endpoints.InstancesEndpoint,
		DecodeWSInstancesRequest,
		EncodeGRPCInstancesResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"StopContainer",
		ws.ProtoIDFromString("STO"),
		endpoints.StopContainerEndpoint,
		DecodeWSStopContainerRequest,
		EncodeGRPCStopContainerResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"Execute",
		ws.ProtoIDFromString("EXE"),
		endpoints.ExecuteEndpoint,
		DecodeWSExecuteRequest,
		EncodeGRPCExecuteResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetEnv",
		ws.ProtoIDFromString("GEV"),
		endpoints.GetEnvEndpoint,
		DecodeWSGetEnvRequest,
		EncodeGRPCGetEnvResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"SetEnv",
		ws.ProtoIDFromString("SEV"),
		endpoints.SetEnvEndpoint,
		DecodeWSSetEnvRequest,
		EncodeGRPCSetEnvResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"IDForName",
		ws.ProtoIDFromString("IFN"),
		endpoints.IDForNameEndpoint,
		DecodeWSIDForNameRequest,
		EncodeGRPCIDForNameResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetContainerKMI",
		ws.ProtoIDFromString("GCK"),
		endpoints.GetContainerKMIEndpoint,
		DecodeWSGetContainerKMIRequest,
		EncodeGRPCGetContainerKMIResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"SetLink",
		ws.ProtoIDFromString("SLI"),
		endpoints.SetLinkEndpoint,
		DecodeWSSetLinkRequest,
		EncodeGRPCSetLinkResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"RemoveLink",
		ws.ProtoIDFromString("RLI"),
		endpoints.RemoveLinkEndpoint,
		DecodeWSRemoveLinkRequest,
		EncodeGRPCRemoveLinkResponse,
	))

	service.AddEndpoint(ws.NewServiceEndpoint(
		"GetLinks",
		ws.ProtoIDFromString("GLI"),
		endpoints.GetLinksEndpoint,
		DecodeWSGetLinksRequest,
		EncodeGRPCGetLinksResponse,
	))

	return service
}

// DecodeWSCreateContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS CreateContainer request to a container.proto-domain createcontainer request.
func DecodeWSCreateContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.CreateContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCCreateContainerRequest(ctx, req)
}

// DecodeWSRemoveContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveContainer request to a container.proto-domain removecontainer request.
func DecodeWSRemoveContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveContainerRequest(ctx, req)
}

// DecodeWSInstancesRequest is a websocket.DecodeRequestFunc that converts a
// WS Instances request to a container.proto-domain instances request.
func DecodeWSInstancesRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.InstancesRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCInstancesRequest(ctx, req)
}

// DecodeWSStopContainerRequest is a websocket.DecodeRequestFunc that converts a
// WS StopContainer request to a container.proto-domain stopcontainer request.
func DecodeWSStopContainerRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.StopContainerRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCStopContainerRequest(ctx, req)
}

// DecodeWSExecuteRequest is a websocket.DecodeRequestFunc that converts a
// WS Execute request to a container.proto-domain execute request.
func DecodeWSExecuteRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.ExecuteRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCExecuteRequest(ctx, req)
}

// DecodeWSGetEnvRequest is a websocket.DecodeRequestFunc that converts a
// WS GetEnv request to a container.proto-domain getenv request.
func DecodeWSGetEnvRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetEnvRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetEnvRequest(ctx, req)
}

// DecodeWSSetEnvRequest is a websocket.DecodeRequestFunc that converts a
// WS SetEnv request to a container.proto-domain setenv request.
func DecodeWSSetEnvRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.SetEnvRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCSetEnvRequest(ctx, req)
}

// DecodeWSIDForNameRequest is a websocket.DecodeRequestFunc that converts a
// WS IDForName request to a container.proto-domain idforname request.
func DecodeWSIDForNameRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.IDForNameRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCIDForNameRequest(ctx, req)
}

// DecodeWSGetContainerKMIRequest is a websocket.DecodeRequestFunc that converts a
// WS GetContainerKMI request to a container.proto-domain getcontainerkmi request.
func DecodeWSGetContainerKMIRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetContainerKMIRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetContainerKMIRequest(ctx, req)
}

// DecodeWSSetLinkRequest is a websocket.DecodeRequestFunc that converts a
// WS SetLink request to a container.proto-domain setlink request.
func DecodeWSSetLinkRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.SetLinkRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCSetLinkRequest(ctx, req)
}

// DecodeWSRemoveLinkRequest is a websocket.DecodeRequestFunc that converts a
// WS RemoveLink request to a container.proto-domain removelink request.
func DecodeWSRemoveLinkRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.RemoveLinkRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCRemoveLinkRequest(ctx, req)
}

// DecodeWSGetLinksRequest is a websocket.DecodeRequestFunc that converts a
// WS GetLinks request to a container.proto-domain getlinks request.
func DecodeWSGetLinksRequest(ctx context.Context, data interface{}) (interface{}, error) {
	req := &pb.GetLinksRequest{}
	err := proto.Unmarshal(data.([]byte), req)
	if err != nil {
		return nil, err
	}

	return DecodeGRPCGetLinksRequest(ctx, req)
}
