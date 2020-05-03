package player

type Player struct {
	id     int
	roomId int
}

var players [15]Player

func New(id int, roomId int) Player {
	p := (Player{id: id, roomId: roomId})
	players[0] = p
	return p
}

func (player Player) GetID() int {
	return player.id
}

func (player Player) GetRoomID() int {
	return player.roomId
}
