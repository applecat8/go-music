package start

import (
	"io/ioutil"
	"log"
	"strings"
)

const MUSIC_DIR string = "/home/applecat/Recreation/music/"

//加载所有音乐文件,获得mp3文件的路径
func LoadMusics() []string {
    dir, err := ioutil.ReadDir(MUSIC_DIR)
    if err != nil {
       log.Fatal(err)
    }

    var musics []string

    for _, file := range dir {
        //进行筛选
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp3") {
            musics = append(musics, MUSIC_DIR + file.Name())
        }
    }
    return musics
}
