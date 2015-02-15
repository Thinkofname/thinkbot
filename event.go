package thinkbot

type Event interface{}

type Stop struct{}

type Connected struct{}

type ChannelMessage struct {
	Sender  User
	Channel string
	Message string
	CTCP    bool
}

type PrivateMessage struct {
	Sender  User
	Target  string
	Message string
	CTCP    bool
}
