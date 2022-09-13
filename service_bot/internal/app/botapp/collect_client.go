package botapp

import (
	ci "github.com/xamust/petbot/service_bot/api"
	"google.golang.org/grpc"
)

type client struct {
	collectBot *BotApp
	conn       *grpc.ClientConn
	infoClient ci.ClubsInfoClient
}

func (c *client) Start() error {
	conn, err := grpc.Dial(c.collectBot.config.PortgRPC, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.infoClient = ci.NewClubsInfoClient(c.conn)
	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
