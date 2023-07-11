package draw

import "image"

type TwitchUser struct {
	ProfileImage image.Image
	Username     string
	StreamerType string
}

type Streamer struct {
	*TwitchUser
	Oldest *TwitchUser
	Date   string
}
