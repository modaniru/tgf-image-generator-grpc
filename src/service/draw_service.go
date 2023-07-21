package service

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/modaniru/image-generator/src/draw"
	"github.com/modaniru/image-generator/src/utils"
)

const (
	imageWidth = 1120.0

	headerPadding = 20.0
	headerHeight  = 240.0

	inputCardHeight = 80.0
	inputCardMargin = 20.0
	inputCardWifth  = 200.0

	streamerCardHeight = 220.0
	streamerCardWidth  = 200.0
	streamerCardMargin = 20.0

	countInRows = 5
)

type DrawerService struct {
	static       map[string]image.Image
	defaultStyle *draw.ImageStyle
}

func NewDrawerService(static map[string]image.Image, style *draw.ImageStyle) *DrawerService {
	return &DrawerService{static: static, defaultStyle: style}
}

func (d *DrawerService) DrawImage(inputedUsers []*draw.TwitchUser, streamers []*draw.Streamer) (image.Image, error) {
	height := float64(headerHeight)
	inputedUsersRows := utils.CalculateRowsCount(len(inputedUsers), countInRows)
	streamersRows := utils.CalculateRowsCount(len(streamers), countInRows)
	height += float64(inputedUsersRows) * (inputCardHeight + inputCardMargin)
	height += float64(streamersRows) * (streamerCardHeight + streamerCardMargin)
	context := gg.NewContext(imageWidth, int(height))
	drawer := draw.NewDrawer(context, d.static, d.defaultStyle)
	drawer.DrawRectangleGradient(0, 0, imageWidth, height, d.defaultStyle.Background)

	{
		err := drawer.DrawHeader(headerPadding, headerPadding, imageWidth-headerPadding-headerPadding, headerHeight-headerPadding-headerPadding, "Twitch General Follows")
		if err != nil {
			return nil, err
		}
		count := 0
		bottom := headerHeight
		for i := 0; i < inputedUsersRows; i++ {
			left := inputCardMargin
			for j := 0; j < countInRows && count < len(inputedUsers); j++ {
				err = drawer.DrawInputedUsers(left, bottom, inputCardWifth, inputCardHeight, inputedUsers[count])
				if err != nil {
					return nil, err
				}
				left += inputCardWifth + inputCardMargin
				count++
			}
			bottom += inputCardHeight + inputCardMargin
		}
		count = 0
		for i := 0; i < streamersRows; i++ {
			left := streamerCardMargin
			for j := 0; j < countInRows && count < len(streamers); j++ {
				err = drawer.DrawFragment(left, bottom, streamerCardWidth, streamerCardHeight, streamers[count])
				if err != nil {
					return nil, err
				}
				left += streamerCardWidth + streamerCardMargin
				count++
			}
			bottom += streamerCardHeight + streamerCardMargin
		}
	}
	return context.Image(), nil
}
