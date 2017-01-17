package main

import (
	"fmt"
	"os"
	"time"

	"github.com/goldmoment/dataloader"
	"github.com/goldmoment/manager"
	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am runnning task.")

	pts := dbl.GetPictureTimeouts(time.Now())
	fmt.Println("Total: ", len(pts))

	for _, pt := range pts {
		err := os.Remove(".." + pt.Path)
		if err != nil {
			fmt.Println(err.Error())
		}

		dbl.RemovePictureTimeout(pt)
		dbl.RemovePicture(pt)
	}

}

func main() {

	db.InitDatabase()

	gocron.Every(10).Minutes().Do(task)
	<-gocron.Start()
}
