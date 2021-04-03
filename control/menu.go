package control

import (
    "fmt"
	"log"
    "github.com/eiannone/keyboard"
)

//监听键盘，当按下回车是换下一首
//func Menu(next, pause chan bool){
//    reader := bufio.NewReader(os.Stdin)
//    for {
//        fmt.Println("---")
//        char, _, err := reader.ReadRune()
//        if err != nil {
//            log.Fatal(err)
//        }
//        switch char{
//        case 'n':
//            next <- true
//        case 'p':
//            pause <- true
//        }
//    }
//}

func Menu(next, pause, exit chan bool){
    if err := keyboard.Open(); err != nil {
        log.Fatal(err)
    }

    defer keyboard.Close()

    for {
        fmt.Println("<Space> 暂停, <Enter> 下一首, <Esc> 退出")
        _, key, err := keyboard.GetSingleKey()

        if err != nil {
            log.Fatal(err)
        }

        if key == keyboard.KeySpace {
            pause <- true
        }else if key == keyboard.KeyEnter{
            next <- true
        }else if key == keyboard.KeyEsc {
            exit <- true
            break
        }
    }
}
