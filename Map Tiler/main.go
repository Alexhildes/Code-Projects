package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
	"sync"

	"time"

	"gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/urfave/cli.v2"
)

func main() {

	app := &cli.App{
		Name:  "mapslizle",
		Usage: "Slice a giant image into tiles for mapping applications",
		Authors: []*cli.Author{&cli.Author{
			Name:  "James Kemp",
			Email: "james.kemp@schneider-electric.com",
		}},
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dst",
				Value: "D:\\00Floor\\Tiles",
				Usage: "destination directory",
			},
			&cli.StringFlag{
				Name:  "src",
				Value: "",
				Usage: "source file",
			},
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Map or floor name",
			},
			&cli.UintFlag{
				Name:  "sq",
				Value: 250,
				Usage: "Square size in pixels",
			},
			&cli.UintFlag{
				Name:  "zoom",
				Value: 0,
				Usage: "Zoom level",
			},
			&cli.BoolFlag{
				Name:    "originBottom",
				Aliases: []string{"ob"},
				Value:   true,
				Usage:   "Map starts at bottom left and goes up",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "slice",
				Action: func(c *cli.Context) error {
					s := Slicer{
						SqSize:   int(c.Uint("sq")),
						Dst:      c.String("dst"),
						Src:      c.String("src"),
						Name:     c.String("name"),
						Zoom:     int(c.Uint("zoom")),
						BottomUp: c.Bool("originBottom"),
					}
					return s.Slice()
				},
			},
		},
	}

	app.Run(os.Args)
}

type Slicer struct {
	SqSize   int
	BottomUp bool
	Dst      string
	Src      string
	Name     string
	Zoom     int
}

func (s Slicer) Slice() (err error) {
	if s.SqSize < 1 {
		err = errors.New("Square size must be larger than 1")
		return
	}

	if len(s.Name) < 1 {
		err = errors.New("Name must be provided")
		return
	}

	srcImgFile, err := os.Open(s.Src)
	if err != nil {
		return
	}
	defer srcImgFile.Close()

	stat, err := srcImgFile.Stat()
	if err != nil {
		return
	}

	fmt.Println("Decoding")
	loadBar := pb.New64(stat.Size()).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 20)
	loadBar.Start()
	reader := loadBar.NewProxyReader(srcImgFile)

	srcImg, _, err := image.Decode(reader)
	loadBar.Finish()
	reader.Close()
	if err != nil {
		return
	}

	fmt.Println("\nSlicing")
	bar := pb.StartNew(srcImg.Bounds().Max.X / s.SqSize)
	//ysteps := int(math.Ceil(float64(srcImg.Bounds().Max.Y) / float64(s.SqSize)))
	ysteps := srcImg.Bounds().Max.Y / s.SqSize
	fmt.Println("\nHeight:", srcImg.Bounds().Max.Y)
	fmt.Println("Tiles :", ysteps)

	for x := 0; x*s.SqSize < srcImg.Bounds().Max.X; x += 1 {
		wg := sync.WaitGroup{}
		for y := 0; y <= ysteps; y += 1 {
			wg.Add(1)
			go func(x, y int) {
				if s.BottomUp {
					err := s.Encode(srcImg, x, y, strconv.Itoa(ysteps-y)+".png")
					if err != nil {
						log.Println(err)
					}
				} else {
					err := s.Encode(srcImg, x, y, strconv.Itoa(y)+".png")
					if err != nil {
						log.Println(err)
					}
				}

				wg.Done()
			}(x, y)
		}

		wg.Wait()
		bar.Increment()
	}
	return
}

func (s Slicer) Encode(srcImg image.Image, x int, y int, tilename string) error {
	var dstImgFile *os.File
	dstImg := image.NewRGBA(image.Rect(0, 0, s.SqSize, s.SqSize))
	draw.Draw(dstImg, dstImg.Bounds(), srcImg, image.Pt(x*s.SqSize, y*s.SqSize), draw.Src)
	dstImgFileName := fmt.Sprintf(s.Dst+"\\%s\\%d\\%d", s.Name, s.Zoom, x)
	os.MkdirAll(dstImgFileName, 0666)
	dstImgFile, err := os.Create(dstImgFileName + "\\" + tilename)
	if err != nil {
		return err
	}
	return png.Encode(dstImgFile, dstImg)
}
