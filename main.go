package main

import (
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/logerror/easylog"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
)

func main() {
	// 定义URL和输出文件名
	url := "https://sspai.com/post/58014"
	outputFile := "screenshot_with_qrcode.png"

	// 创建一个无头Chrome浏览器的上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置截图高度
	screenshotHeight := 800 // 只截取网页的顶部800像素高度
	screenshotWidth := 1920 // 设置截图的宽度

	// 截图的变量
	var buf []byte

	// 执行chromedp任务：打开URL并截取指定区域截图
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.EmulateViewport(int64(screenshotWidth), int64(screenshotHeight)), // 设置视口大小
		chromedp.Sleep(3*time.Second),                                             // 确保页面完全渲染
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		easylog.Error("Error taking screenshot", zap.Error(err))
		return
	}

	// 解码网页截图为图像 (PNG格式)
	img, err := png.Decode(bytes.NewReader(buf))
	if err != nil {
		easylog.Error("Error decoding screenshot:", zap.Error(err))
		return
	}

	// 生成二维码
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		easylog.Error("Error generating QR code", zap.Error(err))
		return
	}
	qrCodeImage := qrCode.Image(128) // 128为二维码的大小

	// 创建一个新的图像，大小与原始截图相同
	combined := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))

	// 绘制截图到新图像
	draw.Draw(combined, img.Bounds(), img, image.Point{}, draw.Src)

	// 二维码的绘制位置在图片的右上角
	qrOffset := image.Pt(img.Bounds().Dx()-qrCodeImage.Bounds().Dx()-10, 10)
	draw.Draw(combined, qrCodeImage.Bounds().Add(qrOffset), qrCodeImage, image.Point{}, draw.Over)

	// 绘制网址文本在二维码的下方
	textOffsetY := qrOffset.Y + qrCodeImage.Bounds().Dy() + 10
	err = addLabel(combined, qrOffset.X, textOffsetY, url, img.Bounds().Dx()-qrOffset.X-10)
	if err != nil {
		easylog.Error("Error adding label:", zap.Error(err))
		return
	}

	// 将合成图像保存为PNG文件
	output, err := os.Create(outputFile)
	if err != nil {
		easylog.Error("Error creating output file:", zap.Error(err))
		return
	}
	defer output.Close()

	err = png.Encode(output, combined)
	if err != nil {
		easylog.Error("Error saving final image:", zap.Error(err))
		return
	}

	easylog.Info("Screenshot with QR code saved to " + outputFile)
}

// addLabel 在给定的图像上绘制文本，确保文本不会超出图像的右边界
func addLabel(img *image.RGBA, x, y int, label string, maxWidth int) error {
	// 加载字体
	fontBytes, err := os.ReadFile("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf") // 确保路径指向你系统中的实际字体文件
	if err != nil {
		return fmt.Errorf("error loading font: %v", err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return fmt.Errorf("error parsing font: %v", err)
	}

	// 初始化 freetype context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(16)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	// 获取字体度量信息
	face := truetype.NewFace(f, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	// 计算文本宽度
	textWidth := calcTextWidth(label, face)

	// 如果文本宽度超过了最大允许宽度，则需要截断
	//if textWidth > maxWidth {
	//	for len(label) > 0 && calcTextWidth(label, face) > maxWidth {
	//		label = label[:len(label)-1] // 逐渐减小文本长度
	//	}
	//	label += "..." // 添加省略号表示文本被截断
	//}

	// 设置绘制位置
	pt := freetype.Pt(x-(textWidth-128), y+int(c.PointToFixed(16)>>6))

	// 绘制文本
	_, err = c.DrawString(label, pt)
	if err != nil {
		return fmt.Errorf("error drawing string: %v", err)
	}

	return nil
}

// calcTextWidth 计算给定文本在指定字体下的宽度
func calcTextWidth(text string, face font.Face) int {
	width := 0
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(x)
		width += awidth.Round()
	}
	return width
}
