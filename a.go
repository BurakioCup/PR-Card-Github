package main

import (
	"os/exec"

	"github.com/canhlinh/svg2png"
	"github.com/sirupsen/logrus"
)

func main() {
	chrome := svg2png.NewChrome().SetHeight(250).SetWith(280)
	filepath := "Soccerball_mask_transparent_background.png"
	if err := chrome.Screenshoot("https://github-readme-stats.vercel.app/api/top-langs/?username=take-2405", filepath); err != nil {

		//if err := chrome.Screenshoot("https://initial-practice.s3-ap-northeast-1.amazonaws.com/second-sprintReview/save.svg", filepath); err != nil {
		logrus.Panic(err)
	}

	exec.Command("open", filepath).Run()
}