package main

import "github.com/skip2/go-qrcode"

//二维码

func main()  {
	qrcode.WriteFile("sb何新月",qrcode.Medium,250,"./demo1.png")
}

