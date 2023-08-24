package handlers

import (
	"context"
	"log"

	"gameserver/internal/game/model/player"
	"gameserver/internal/game/model/room"
	pb "gameserver/proto/pb/room_manager"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRoomManagerServer
}

func CreateHandler(s *grpc.Server) {
	pb.RegisterRoomManagerServer(s, &server{})
}

// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

func (s *server) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomId := in.GetId()
	if id, ok := room.CreateRoom(); !ok {
		return &pb.CreateRoomResponse{
			Id: id,
			Ok: false,
		}, nil
	}

	log.Printf("Create room: %v", roomId)
	return &pb.CreateRoomResponse{Ok: true}, nil
}

func (s *server) EnterRoom(req *pb.EnterRoomRequest, stream pb.RoomManager_EnterRoomServer) error {
	roomId := req.GetRoomID()

	if room.GetRoom(roomId) == nil {
		return nil
	}

	audiences := req.GetAudiences()
	for _, audience := range audiences {
		userName := audience.GetUserName()
		userId := audience.GetUserID()
		player := player.NewAudience(userId, userName, stream)

		if ok := room.GetRoom(roomId).AddAudience(player); !ok {
			return nil
		}

	}

	log.Printf("Enter room: %v", roomId)
	return nil
}
