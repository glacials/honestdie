package handler

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/websocket"
)

type ConnectionStatus struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type RollRequest struct {
	Type       string   `json:"type"`
	Candidates []string `json:"candidates"`
}

type RollCountdown struct {
	Type        string `json:"type"`
	SecondsLeft uint   `json:"secondsLeft"`
}

type Roll struct {
	Type       string   `json:"type"`
	Candidates []string `json:"candidates"`
	Winner     string   `json:"winner"`
}

var conns = make(map[*websocket.Conn]struct{})

func Websocket(conn *websocket.Conn) {
	cs := ConnectionStatus{
		Type:   "status",
		Status: "connected",
	}

	if err := websocket.JSON.Send(conn, &cs); err != nil {
		log.Printf("can't send connection status frame: %s", err)
		conn.Close()
		return
	}

	rr := make(chan RollRequest)

	conns[conn] = struct{}{}

	go standby(rr)
	listen(conn, rr)
}

func standby(c chan RollRequest) {
	for {
		select {
		case rr := <-c:
			for i := uint(3); i > 0; i -= 1 {
				for conn := range conns {
					if err := websocket.JSON.Send(conn, RollCountdown{
						Type:        "countdown",
						SecondsLeft: i,
					}); err != nil {
						log.Printf("can't send countdown frame %d: %s", i, err)
						conn.Close()
						delete(conns, conn)
					}
				}
				time.Sleep(1 * time.Second)
			}

			numCandidates := len(rr.Candidates)
			winner := rand.Intn(numCandidates)

			r := Roll{
				Type:       "roll",
				Candidates: rr.Candidates,
				Winner:     rr.Candidates[winner],
			}

			for conn := range conns {
				if err := websocket.JSON.Send(conn, r); err != nil {
					log.Printf("can't send roll frame: %s", err)
					conn.Close()
					delete(conns, conn)
				}
			}

			// cooldown
			time.Sleep(1 * time.Second)
		}
	}
}

func listen(conn *websocket.Conn, c chan RollRequest) {
	for {
		r := RollRequest{
			Type: "roll",
		}

		if err := websocket.JSON.Receive(conn, &r); err != nil {
			log.Printf("can't receive frame: %s", err)
			conn.Close()
			return
		}

		c <- r
	}
}
