package main

import (
	"music/control"
	"music/start"
	"os"
	"time"
)

func main() {

    musics := start.LoadMusics()
    if len(musics) == 0 {
        os.Exit(1)
    }

    next := make(chan bool)
    pause := make(chan bool)
    exit := make(chan bool)
    control.PlayMusic(musics[0],next)

    go control.Menu(next, pause, exit)

    for {
        select {
        case <- next: //播放下一首
            control.PlayNextMusic(musics, next)
        case <- pause: //暂停播放
            control.PauseCurrentMusic()
        case <- exit: //退出播放
            time.Sleep(time.Second)
            control.Close()
            os.Exit(1)
        }
    }
}
