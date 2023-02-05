package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"

	. "github.com/kkdai/photomgr"
)

type NullWriter int

func printPageResult(p *PTT, count int) {
	for i := 0; i < count; i++ {
		title := p.GetPostTitleByIndex(i);
		likeCounter := p.GetPostStarByIndex(i);
		fmt.Printf("%d:[%d★]%s\n", i, likeCounter, title)
	}

	fmt.Printf("(o: open file in fider, s: top page, n:next, p:prev, quit: quit program)\n")
}

func(NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	log.SetOutput(new(NullWriter))
	ptt := NewPTT()

	user, _ := user.Current()
	ptt.BaseDir = fmt.Sprintf("%v/Pictures/PTT", user.HomeDir)

	var workerNumber int

	rootCmd := &cobra.Command{
		Use: "ptt",
		Short: "Download all the images in given post url",
		Run: func(cmd *cobra.Command, args []string) {
			page := 0
			pagePostCount := 0
			pagePostCount = ptt.ParsePttPageByIndex(page)
			printPageResult(ptt, pagePostCount)

			scanner := bufio.NewScanner(os.Stdin)
			isQuit := false

			for !isQuit {
				fmt.Print("ptt:> ")

				if !scanner.Scan() {
					break
				}

				line := scanner.Text()
				parts := strings.Split(line, " ")
				cmd := parts[0]
				args := parts[1:]

				switch cmd {
				case "quit":
					isQuit = true;
				case "n":
					page = page + 1
					pagePostCount = ptt.ParsePttPageByIndex(page)
					printPageResult(ptt, pagePostCount)
				case "p":
					if page > 0 {
						page = page - 1
					}

					pagePostCount = ptt.ParsePttPageByIndex(page)
					printPageResult(ptt, pagePostCount)
				case "s":
					
				}
			}
		},
	}
}
