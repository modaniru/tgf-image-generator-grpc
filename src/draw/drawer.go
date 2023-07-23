package draw

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/modaniru/image-generator/src/utils"
)

type Drawer struct {
	staticImages map[string]image.Image
	c            *gg.Context
	style        *ImageStyle
}

// Конструктор рисовальщика
func NewDrawer(c *gg.Context, staticImages map[string]image.Image, style *ImageStyle) *Drawer {

	return &Drawer{c: c, staticImages: staticImages, style: style}
}

// Рисует header картинки
func (d *Drawer) DrawHeader(x, y, w, h float64, title string) error {
	d.DrawRectangleGradient(x, y, w, h, d.style.HeaderStyle.Background)
	err := d.DrawText(d.c, x+w/2, y+h/2, w, h, 0.5, 0.5, "Twitch General Follows", d.style.HeaderStyle.TextColors)
	if err != nil {
		return err
	}
	return nil
}

// Рисует текст
func (d *Drawer) DrawText(c *gg.Context, x, y, w, h, hor, ver float64, title string, color color.Color) error {
	c.Push()
	{
		err := c.LoadFontFace("fonts/B612-Bold.ttf", float64(utils.CalculateFontSize(w, h, len(title))))
		if err != nil {
			return err
		}
		d.c.SetColor(color)
		d.c.DrawStringAnchored(title, x, y, hor, ver)
	}
	c.Pop()
	return nil
}

// Рисует фрагмент входящих пользователей
func (d *Drawer) DrawInputedUsers(x, y, w, h float64, user *TwitchUser) error {
	d.DrawRectangleGradient(x, y, w, h, d.style.InputedUsersCard.Background)
	d.DrawProfileImage(user.ProfileImage, x+w*0.2, y+h*0.5, 60)
	err := d.DrawText(d.c, x+(w*0.4), y+h*0.5, w*0.6, h, 0, 0.5, user.Username, d.style.InputedUsersCard.TextColors)
	if err != nil {
		return err
	}
	return nil
}

// Рисует фрагмент общих подписок
func (d *Drawer) DrawFragment(x, y, w, h float64, streamer *Streamer) error {

	d.DrawRectangleGradient(x, y, w, h, d.style.GeneralFollowsCard.Background)
	d.DrawProfileImage(streamer.Oldest.ProfileImage, x+w*0.8, y+h*0.2, 60)
	d.DrawProfileImage(streamer.ProfileImage, x+w*0.5, y+h*0.45, 120)
	streamerType := d.staticImages[streamer.StreamerType]
	if streamerType != nil {
		d.DrawProfileImage(streamerType, x+w*0.7, y+h*0.6, 36)
	}
	d.c.Push()
	{
		d.c.RotateAbout(gg.Radians(25), x+w*0.87, y+h*0.06)
		d.DrawImageAnchorPoint(d.staticImages["crown"], x+w*0.87, y+h*0.06, 80, 60)
	}
	d.c.Pop()
	err := d.DrawText(d.c, x+(w*0.05), y+h*0.1, w*0.9, 20, 0, 0.5, streamer.Date[0:10], d.style.GeneralFollowsCard.TextColors)
	if err != nil {
		return err
	}
	err = d.DrawText(d.c, x+w*0.5, y+h*0.8, w*0.9, 50, 0.5, 0.5, streamer.Username, d.style.GeneralFollowsCard.TextColors)
	if err != nil {
		return err
	}
	return nil
}

// Рисует круглую картинку
func (d *Drawer) DrawProfileImage(image image.Image, x, y, side float64) {
	d.c.Push()
	{
		d.c.SetHexColor("#fff")
		d.c.DrawEllipse(x, y, side/2, side/2)
		d.c.Fill()
		d.c.DrawEllipse(x, y, side/2, side/2)
		d.c.Clip()
		d.DrawImageAnchorPoint(image, x, y, side, side)
		d.c.ResetClip()
	}
	d.c.Pop()
}

// Рисует картинку относительно точки
func (d *Drawer) DrawImageAnchorPoint(image image.Image, x, y, w, h float64) {
	d.c.Push()
	{
		coeffWidthSize := w / float64(image.Bounds().Dx())
		coeffHeightSize := h / float64(image.Bounds().Dy())
		d.c.Scale(coeffWidthSize, coeffHeightSize)
		d.c.DrawImageAnchored(image, int(x*(1/coeffWidthSize)), int(y*(1/coeffHeightSize)), 0.5, 0.5)
	}
	d.c.Pop()
}

// Рисует прямоугольник по заданным цветам
func (d *Drawer) DrawRectangleGradient(x, y, w, h float64, colors Colors) {
	d.DrawRectangleShadow(x, y, w, h, 10, 10)
	d.c.Push()
	{
		gradient := gg.NewLinearGradient(x, y, x+w, y+h)
		d.getGradient(&gradient, colors)
		d.c.SetFillStyle(gradient)
		d.c.DrawRectangle(x, y, w, h)
		d.c.Fill()
	}
	d.c.Pop()
}

// Устанавливает градиенту цвета
func (d *Drawer) getGradient(gradient *gg.Gradient, colors Colors) {
	if len(colors) == 1 {
		(*gradient).AddColorStop(0, colors[0])
		(*gradient).AddColorStop(1, colors[0])
	} else {
		coefficient := 1 / float64(len(colors)-1)
		for i, c := range colors {
			(*gradient).AddColorStop(float64(i)*coefficient, c)
		}
	}
}

// Рисует прямоугольную тень по координатам (x + dx; y + dy)
func (d *Drawer) DrawRectangleShadow(x, y, w, h, dx, dy float64) {
	d.c.Push()
	{
		d.c.SetHexColor("#00000072")
		d.c.DrawRectangle(x+dx, y+dy, w, h)
		d.c.Fill()
	}
	d.c.Pop()
}
