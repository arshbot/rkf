package keyboard

import (
	"fmt"
	"log"
	"os"
	"time"
)

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

func InfiniteRainbow() {
	files := []string{"color_center", "color_left", "color_right", "color_extra"}

	open := func(files []string) []*os.File {
		kbfiles := make([]*os.File, 0, len(files))
		for _, file := range files {
			p := fmt.Sprintf("/sys/class/leds/system76::kbd_backlight/%v", file)
			fh, err := os.OpenFile(p, os.O_RDWR, 0755)
			if err != nil {
				log.Fatal("error: %v", err)
				continue
			}
			kbfiles = append(kbfiles, fh)
		}
		return kbfiles
	}

	colors := make([]string, 0, 6)
	// generate range of rainbow values
	for i := 0; i <= 255; i++ {
		c := RGBColor{255, i, 0}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{i, 255, 0}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{0, 255, i}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{0, i, 255}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 0; i <= 255; i++ {
		c := RGBColor{i, 0, 255}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	for i := 255; i >= 0; i-- {
		c := RGBColor{255, 0, i}
		hex := c.GetColorInHex()
		colors = append(colors, hex)
	}

	kbfiles := open(files)
	for {
		for _, c := range colors {
			for _, f := range kbfiles {
				f.WriteString(c)
			}
			time.Sleep(time.Nanosecond)
		}
	}
}
