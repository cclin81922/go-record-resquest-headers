package main

func loop(done <-chan bool) {
    for {
        select {
        case <-done:
            return
        }
    }
}

func main() {
    parseArgs()
    done := registerSignal()
    HttpServe()
    loop(done)
}
