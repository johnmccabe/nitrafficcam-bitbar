package main

import (
	"image"
	_ "image/jpeg"
	"net/http"
	"os"

	bitbar "github.com/johnmccabe/go-bitbar"
)

func main() {
	b := bitbar.New()

	s := bitbar.Style{
		Color: "black",
		Size:  18,
	}

	img, err := getImage("https://video.zidivo.com/live10/2B/preview.jpg")
	if err != nil {
		b.StatusLine("❗️")
		errorMenu := b.NewSubMenu()
		errorMenu.Line(err.Error())
		b.Render()
		os.Exit(1)
	}

	b.StatusLine("🚙")

	menu := b.NewSubMenu()
	menu.Line("M2 - Sandyknowes").Style(s)
	menu.Image(img).Href("http://trafficwatchni.com/cameras?p_p_id=cctv_WAR_twniportlet_INSTANCE_37bihh0Nomul&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_resource_id=stream&p_p_cacheability=cacheLevelPage&p_p_col_id=column-1&p_p_col_pos=1&p_p_col_count=2&_cctv_WAR_twniportlet_INSTANCE_37bihh0Nomul_id=6")

	b.Render()
}

func getImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
