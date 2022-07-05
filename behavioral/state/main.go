package main

func main() {
	video := &context{}
	video.setState(&readyState{context: video})

	video.play()
	video.play()
	video.pause()
	video.pause()
	video.play()
	video.pause()
	video.play()
	video.setState(&finishedState{context: video})
	video.pause()
	video.play()
	video.play()
	video.pause()
	video.setState(&finishedState{context: video})
	video.play()
}
