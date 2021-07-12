package server

import "fmt"
import "errors"
import "strings"
import "github.com/simba-fs/go-chat/internal/room"
import "github.com/simba-fs/go-chat/internal/cmdParser"

var cl cmdParser.CmdList

var (
	ErrNoConnection = errors.New("no connection")
	ErrNoRoom = errors.New("no default room")
)

// func wrapper(f func()) func(raw string, cmds []string, exec cmdParser.FuncExec)(string, error){
//
// }

func init(){
	c := []cmdParser.Cmd{
		cmdParser.New("msg", "echo message", func(raw string, cmds []string, exec cmdParser.FuncExec, arg ...interface{})(string, error){
			msg := strings.Join(cmds[1:], " ")

			// get conn
			conn, ok := arg[0].(*room.Connection)
			if !ok {
				return "", ErrNoConnection
			}

			// get room
			// curtRoom means current room
			curtRoom, ok := arg[1].(*room.Room)
			if !ok {
				return "", ErrNoRoom
			}

			fmt.Printf("%s#%s: %s\n", conn.Name, conn.Room.Name, msg)

			curtRoom.Broadcast("msg", fmt.Sprintf("%s: %s", conn.Name, msg))

			// err := conn.WriteMessage(websocket.TextMessage, []byte(raw))
			return "", nil
		}),
	}

	cl = cmdParser.CmdList{
		Cmds: c,
		Help: "help",
		Helper: cmdParser.Helper,
	}
}

func exec(cmd []byte, conn *room.Connection, room *room.Room){
	cl.Exec(string(cmd), conn, room)
}
