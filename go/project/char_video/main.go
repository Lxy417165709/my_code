package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gdamore/tcell"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Lxy417165709/gcm/char_photo"
	_ "github.com/disintegration/imaging"
	_ "github.com/gdamore/tcell"
	"log"
)

func main() {
	// 1. 初始化。
	var charSet = []byte(` .,:;i1tfLCG08@`)
	frameNoToPixelMatrixMap := map[int]*char_photo.PixelMatrix{}
	imgPathFormat := "C:/Users/李学悦/Desktop/all/github/Lxy417165709/test/photo/mykk/frame_%05d.bmp"
	bgFrameNo, endFrameNo := 1, 1774

	// 2. 处理。
	for frameNo := bgFrameNo; frameNo <= endFrameNo; frameNo++ {
		imgPath := fmt.Sprintf(imgPathFormat, frameNo)
		pixelMatrix, err := char_photo.BuildPixelMatrixByImgPath(imgPath, 100, 46, charSet)
		if err != nil {
			log.Fatalf("BuildPixelMatrixByImgPath fail. err=[%+v].", err)
			return
		}
		frameNoToPixelMatrixMap[frameNo] = pixelMatrix
		fmt.Printf("\r进度: %d/%d", frameNo-bgFrameNo+1, endFrameNo-bgFrameNo+1)
	}

	// 3. 输出。
	//for frameNo := bgFrameNo; frameNo <= endFrameNo; frameNo++ {
	//	go fmt.Println("\n" + frameNoToPixelMatrixMap[frameNo].ToCharPhotoColorful() + "\n")
	//	time.Sleep(33 * time.Millisecond)
	//}

	//if err := ui.Init(); err != nil {	// 这里有点奇怪，为什么颜色会偏向白色呢
	//	log.Fatalf("failed to initialize termui: %v", err)
	//}
	//defer ui.Close()
	//p := widgets.NewParagraph()
	//p.SetRect(0, 0, 120, 50)
	//go playMusic()
	//for frameNo := bgFrameNo; frameNo <= endFrameNo; frameNo++ {
	//	go func(frameNo int) {
	//		m := frameNoToPixelMatrixMap[frameNo]
	//		charLines := make([]string, 0)
	//		for y := 0; y < len(m.Matrix); y++ {
	//			var charLine string
	//			for x := 0; x < len(m.Matrix[y]); x++ {
	//				pixel := m.Matrix[y][x]
	//				charLine += fmt.Sprintf("[%s](fg:%s,mod:bold)", string(pixel.Char), rgbToAnsi8(pixel.Color))
	//			}
	//			charLines = append(charLines, charLine)
	//		}
	//		p.Text = strings.Join(charLines, "\n")
	//		ui.Render(p)
	//	}(frameNo)
	//	time.Sleep(33333 * time.Microsecond)
	//}

	//if err := ui.Init(); err != nil { // 这样的话，和之前的效果是一样的。
	//	log.Fatalf("failed to initialize termui: %v", err)
	//}
	//defer ui.Close()
	//p := widgets.NewParagraph()
	//p.SetRect(0, 0, 120, 50)
	//go playMusic()
	//for frameNo := bgFrameNo; frameNo <= endFrameNo; frameNo++ {
	//	go func(frameNo int) {
	//		m := frameNoToPixelMatrixMap[frameNo]
	//		charLines := make([]string, 0)
	//		for y := 0; y < len(m.Matrix); y++ {
	//			var charLine string
	//			for x := 0; x < len(m.Matrix[y]); x++ {
	//				pixel := m.Matrix[y][x]
	//				s := strconv.Quote(char_photo.DecorateWithColor(pixel.Char, pixel.Color))
	//				s = s[1 : len(s)-1]
	//				text, color, r, g, b, err := getTextAndColor8(s)
	//				if err != nil {
	//					panic(err)
	//				}
	//				if r != int(pixel.Color.R) || g != int(pixel.Color.G) || b != int(pixel.Color.B) {
	//					panic(fmt.Sprintf("存在不一致的颜色,frameNo=[%+v],(%+v,%+v), (%+v,%+v,%+v) -> (%+v,%+v,%+v)",
	//						frameNo, x, y, pixel.Color.R, pixel.Color.G, pixel.Color.B, r, g, b))
	//				}
	//				charLine += fmt.Sprintf("[%s](fg:%s,mod:bold)", text, color)
	//			}
	//			charLines = append(charLines, charLine)
	//		}
	//		p.Text = strings.Join(charLines, "\n")
	//		ui.Render(p)
	//	}(frameNo)
	//	time.Sleep(33333 * time.Microsecond)
	//}

	// 下面是可行版本！！！！！！ 但是速度有点慢。
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.EnableMouse()
	//if err := ui.Init(); err != nil { // 这样的话，和之前的效果是一样的。 (性能太拉了...)
	//	log.Fatalf("failed to initialize termui: %v", err)
	//}
	//defer ui.Close()
	var mutex sync.Mutex
	go playMusic()
	t1s := make([]time.Time, 0)
	t2s := make([]time.Time, 0)
	t3s := make([]time.Time, 0)

	wg := sync.WaitGroup{}
	for frameNo := bgFrameNo; frameNo <= endFrameNo; frameNo += 1 {
		go func(frameNo int, startTime time.Time) {
			wg.Add(1)
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			if time.Now().UnixMicro()-startTime.UnixMicro() >= 10000 {
				return
			}
			m := frameNoToPixelMatrixMap[frameNo]
			t1s = append(t1s, time.Now())
			for y := 0; y < len(m.Matrix); y++ {
				for x := 0; x < len(m.Matrix[y]); x++ {
					pixel := m.Matrix[y][x]
					newStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(pixel.Color.R),
						int32(pixel.Color.G), int32(pixel.Color.B)))
					s.SetContent(x, y, rune(pixel.Char), nil, newStyle)
				}
			}
			t2s = append(t2s, time.Now())
			s.Show()
			t3s = append(t3s, time.Now())
		}(frameNo, time.Now())
		time.Sleep(33333 * time.Microsecond)
	}

	wg.Wait()
	lines := make([]string, 0)
	for i := 0; i < len(t1s); i++ {
		lines = append(lines, fmt.Sprintf("%+v, %+v", t3s[i].Sub(t2s[i]), t2s[i].Sub(t1s[i])))
	}
	path := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\my_code\go\char_video\f.md`
	if err := ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0777); err != nil {
		panic(err)
	}
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func playMusic() {
	// 1. 打开mp3文件
	audioFile, err := os.Open("C:\\Users\\李学悦\\Desktop\\all\\github\\Lxy417165709\\test\\photo\\kk.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	defer audioFile.Close()

	// 对文件进行解码
	audioStreamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		log.Fatal(err)
	}

	defer audioStreamer.Close()
	// SampleRate is the number of samples per second. 采样率
	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// 用于数据同步，当播放完毕的时候，回调函数中通过chan通知主goroutine
	done := make(chan bool)
	// 这里播放音乐
	speaker.Play(beep.Seq(audioStreamer, beep.Callback(func() {
		// 播放完成调用回调函数
		done <- true
	})))

	// 等待播放完成
	<-done
}

//func rgbToAnsi8New(color *char_photo.Color) string {
//	r, g, b := color.R, color.G, color.B
//	return rgbToAnsi8(r, g, b)
//}

func getTextAndColor8(str string) (string, string, int, int, int, error) {
	t, r, g, b, err := parseANSI(str)
	if err != nil {
		return "", "", 0, 0, 0, err
	}
	return t, rgbToAnsi8(r, g, b), r, g, b, nil
}

func rgbToAnsi8(r, g, b int) string {
	if r < 128 && g < 128 && b < 128 {
		return "black" // black
	} else if r >= 128 && g >= 128 && b >= 128 {
		return "white" // white
	} else if r >= 128 && g < 128 && b < 128 {
		return "red" // red
	} else if g >= 128 && r < 128 && b < 128 {
		return "green" // green
	} else if b >= 128 && r < 128 && g < 128 {
		return "blue" // blue
	} else if r >= 128 && g >= 128 && b < 128 {
		return "yellow" // yellow
	} else if r >= 128 && b >= 128 && g < 128 {
		return "magenta" // magenta
	} else if g >= 128 && b >= 128 && r < 128 {
		return "cyan" // cyan
	} else {
		return "black" // unsupported color
	}
}

func parseANSI(s string) (text string, r, g, b int, err error) {
	// 定义正则表达式匹配 ANSI 转义序列
	str := `\\x1b\[(\d+);(\d+);(\d+)m(.+?)`
	re := regexp.MustCompile(str)
	matches := re.FindStringSubmatch(s)
	//fmt.Println(len(matches), s)
	if len(matches) != 5 {
		err = fmt.Errorf("invalid ANSI sequence,s=[%+v]", s)
		return
	}
	// 解析颜色信息
	color, err := strconv.Atoi(matches[3])
	if err != nil {
		log.Println(err.Error())
		return
	}

	r, g, b = color256ToRGB(color)

	// 返回文本和颜色信息
	text = matches[4]
	return
}

func color256ToRGB(color int) (r, g, b int) {
	if color < 0 || color > 255 {
		return 0, 0, 0
	}
	if color < 16 {
		// 16 颜色
		r = ((color >> 2) & 1) * 0xAA
		g = ((color >> 1) & 1) * 0xAA
		b = (color & 1) * 0xAA
	} else if color < 232 {
		// 256 颜色
		color -= 16
		b = (color % 6) * 0x33
		color /= 6
		g = (color % 6) * 0x33
		color /= 6
		r = (color % 6) * 0x33
	} else {
		// 24 灰阶
		gray := (color-232)*0x0A + 0x08
		r, g, b = gray, gray, gray
	}
	return
}
