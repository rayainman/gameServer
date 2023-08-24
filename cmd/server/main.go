package main

import (
	"fmt"
	"log"
	"net"

	"gameserver/internal/game/game_logic"
	"gameserver/internal/game/handlers"
	"gameserver/internal/game/model/player"
	"gameserver/internal/game/model/room"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	id, ok := room.CreateRoom()

	if !ok {
		log.Fatalf("failed to create room: %v", err)
	}
	room1 := room.GetRoom(id)
	for i := 0; i < 100; i++ {
		// 創100個玩家
		id := fmt.Sprintf("ID:%d", i)
		p := player.NewAudience(id, fmt.Sprintf("player%d", i), nil)
		room1.AddAudience(p)
	}

	room1.JoinSeat(room1.Audiences)

	// 實際流程是
	// CreatRoom 獲得房間ID

	// room.CreateRoom("2")
	// room2 := room.GetRoom("2")

	// var playLit2 []*player.Player
	// for i := 0; i < 100; i++ {
	// 	// 創100個玩家
	// 	p := player.NewPlayer(i, fmt.Sprintf("player%d", i), nil)
	// 	playLit2 = append(playLit2, p)
	// }

	// room2.AddPlayer(playLit2...)

	go game_logic.StartGame(room1)
	// time.Sleep(1 * time.Second)
	// go game_logic.StartGame(room2)

	s := grpc.NewServer()
	handlers.CreateHandler(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
