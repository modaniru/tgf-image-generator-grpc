package utils

import (
	"errors"
	"image"
	"image/color"
	"net/http"

	"github.com/fogleman/gg"
	pkg "github.com/modaniru/image-generator/pkg/proto"
)

// TODO check width and height bound
func DrawHeaderSquareWithShadow(c *gg.Context, xPos, yPos, width, height float64, title string) error {
	c.SetHexColor("#0000007D")
	c.DrawRectangle(xPos+10, yPos+10, width, height)
	c.Fill()
	c.SetColor(color.RGBA{240, 241, 246, 255})
	c.DrawRectangle(xPos, yPos, width, height)
	c.Fill()
	err := c.LoadFontFace("fonts\\B612-Bold.ttf", float64(CalculateFontSize(width-20, height, len(title))))
	if err != nil {
		return err
	}
	c.SetHexColor("#000")
	c.DrawStringAnchored(title, xPos+width/2, yPos+height/2, 0.5, 0.5)
	c.Fill()
	return nil
}

func FillWithGradient(c *gg.Context, x, y, w, h float64, colors ...color.Color) error {
	if len(colors) < 2 {
		return errors.New("colors length must be greater that 1")
	}
	c.DrawRectangle(x, y, w, h)

	c.SetFillStyle(CreateGradient(x, y, w, h, colors...))
	c.Fill()
	return nil
}

func CreateGradient(x, y, w, h float64, colors ...color.Color) gg.Gradient {
	gradient := gg.NewLinearGradient(x, y, w+x, h+y)
	i := 0.0
	coef := 1.0 / float64(len(colors)-1)
	for _, c := range colors {
		gradient.AddColorStop(i, c)
		i += coef
	}
	return gradient
}

// todo create enum of static images
func DrawFragment(c *gg.Context, x, y, w, h float64, streamer *pkg.ResponseUser, oldest *pkg.ResponseUser, date string, images map[string]image.Image) error {
	date = string(date[:10])
	c.Identity()
	c.SetHexColor("#0000007D")
	c.DrawRectangle(x+5, y+5, w, h)
	c.Fill()

	c.SetHexColor("#DCDCDC")
	c.DrawRectangle(x, y, w, h)
	c.Fill()
	c.Push()
	{
		c.Identity()
		image := images[oldest.DisplayName]
		c.Translate(50, -80)
		c.Scale(0.2, 0.2)
		c.SetHexColor("#fff")
		c.DrawCircle(x*5+w*5/2, y*5+h*5/2, float64(image.Bounds().Dx())/2)
		c.Fill()
		c.DrawCircle(x*5+w*5/2, y*5+h*5/2, float64(image.Bounds().Dx())/2)
		c.Clip()
		c.AsMask()
		c.DrawImageAnchored(image, int(x*5+w*5/2), int(y*5+h*5/2), 0.5, 0.5)
		c.ResetClip()
	}
	c.Pop()

	//Переместить сюда бг и маску?
	c.Push()
	{
		c.Identity()
		image := images[streamer.DisplayName]
		c.Translate(0, -20)
		c.Scale(0.45, 0.45)
		c.SetHexColor("#fff")
		c.DrawCircle(x*(1/0.45)+w*(1/0.45)/2, y*(1/0.45)+h*(1/0.45)/2, float64(image.Bounds().Dx())/2)
		c.Fill()
		c.DrawCircle(x*(1/0.45)+w*(1/0.45)/2, y*(1/0.45)+h*(1/0.45)/2, float64(image.Bounds().Dx())/2)
		c.Clip()
		c.AsMask()
		c.DrawImageAnchored(image, int(x*(1/0.45)+w*(1/0.45)/2), int(y*(1/0.45)+h*(1/0.45)/2), 0.5, 0.5)
		c.ResetClip()
	}
	c.Pop()
	c.Push()
	{
		c.Identity()
		crown, err := gg.LoadPNG("images/crown.png")
		if err != nil {
			return err
		}
		//c.RotateAbout(gg.Radians(30), x+w/2, y+h/2)
		c.Translate(50, -115)
		c.Scale(0.1, 0.1)
		c.DrawImageAnchored(crown, int(x*10+w*5), int(y*10+h*5), 0.5, 0.5)
	}
	c.Pop()
	err := c.LoadFontFace("fonts\\B612-Bold.ttf", float64(CalculateFontSize(w-40, 20, len(date))))
	if err != nil {
		return err
	}
	c.SetHexColor("#636363")
	c.DrawString(date, x+w*0.05, y+h*0.1)
	err = c.LoadFontFace("fonts\\B612-Bold.ttf", float64(CalculateFontSize(w-20, 40, len(streamer.DisplayName))))
	if err != nil {
		return err
	}
	c.SetHexColor("#000")
	c.DrawStringAnchored(streamer.DisplayName, x+w/2, y+h*0.8, 0.5, 0.5)
	return nil
}

type Response struct {
	Image image.Image
	Name  string
}

// go routine
func GetImageFromURI(link, name string, channel chan *Response) {
	resp, err := http.Get(link)
	if err != nil {
		channel <- nil
		return
	}
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		channel <- nil
		return
	}
	channel <- &Response{Name: name, Image: image}
}
