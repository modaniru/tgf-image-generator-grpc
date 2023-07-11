package draw

import "image/color"

type ImageStyle struct {
	Background         Colors
	HeaderStyle        HeaderStyle
	GeneralFollowsCard GeneralFollowsCard
	InputedUsersCard   InputedUsersCard
}

type HeaderStyle struct {
	Background Colors
	TextColors color.Color
}

type GeneralFollowsCard struct {
	Background      Colors
	ImageBackground Colors
	TextColors      color.Color
}

type InputedUsersCard struct {
	Background      Colors
	ImageBackground Colors
	TextColors      color.Color
}

type Colors []color.Color
