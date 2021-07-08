package control

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var index int = 0
var currentMusic *os.File
var ctrl *beep.Ctrl

//播放下一首音乐
func PlayNextMusic(musics []string, next chan bool){
    index = (index + 1) % len(musics)
    currentMusic.Close()
    PlayMusic(musics[index], next)
}

//暂停播放
func PauseCurrentMusic() {
    speaker.Lock()
    ctrl.Paused = !ctrl.Paused
    speaker.Unlock()
}

func Close(){
    currentMusic.Close()
}


//根据歌名播放音乐
func PlayMusic(musicPath string,next chan bool){

    //显示当前播放的音乐名称
    level := strings.Split(musicPath, "/")
    fmt.Println("当前播放音乐:",level[len(level) - 1])

    var err error
    currentMusic, err = os.Open(musicPath)
    if err != nil {
        log.Fatal(err)
    }

    streamer, format, err := mp3.Decode(currentMusic)
    if err != nil {
        log.Fatal(err)
    }

    ctrl = &beep.Ctrl{Streamer: streamer, Paused: false}

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second))
    speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
        next <- true // 音乐结束时播放下一首
    })))
}
