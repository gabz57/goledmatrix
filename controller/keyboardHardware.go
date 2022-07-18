package controller

import (
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"strconv"
)

type KeyboardHard struct {
	KeyboardChannel KeyboardEventChannel
	projection      KeyboardProjection
	listening       bool
}

func (kbh *KeyboardHard) Projection() *KeyboardProjection {
	return &kbh.projection
}

func (kbh *KeyboardHard) EventChannel() *KeyboardEventChannel {
	return &kbh.KeyboardChannel
}

func NewKeyboardHard(keyboardChannel *KeyboardEventChannel) *KeyboardHard {
	var keyChannel *KeyboardEventChannel
	if keyboardChannel != nil {
		keyChannel = keyboardChannel
	} else {
		channel := make(KeyboardEventChannel, KeyboardEventChannelSize)
		keyChannel = &channel
	}

	return &KeyboardHard{
		KeyboardChannel: *keyChannel,
		projection:      *NewKeyboardProjection(),
		listening:       false,
	}
}

func (kbh *KeyboardHard) Start() {
	if !kbh.listening {

		kbh.listening = true
		go kbh.listen()
	}
}

func (kbh *KeyboardHard) Stop() {
	kbh.listening = false
}

func (kbh *KeyboardHard) listen() {

	log.Println("Press ESC button or Ctrl-C to exit this program")
	log.Println("Press any key to see their ASCII code follow by Enter")

	err := termbox.Init()
	if err != nil {
		log.Println("termbox.Init() failed, keyboardHardware not available", err)
		return
	}
	defer termbox.Close()
	//termbox.SetInputMode(termbox.InputEsc)
	for kbh.listening {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				log.Println("Exiting...")
				os.Exit(0)
			case termbox.KeyEsc:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "esc")
			case termbox.KeyArrowLeft:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "left")
			case termbox.KeyArrowRight:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "right")
			case termbox.KeyArrowUp:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "up")
			case termbox.KeyArrowDown:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "down")
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "deleteBackward")
			case termbox.KeyDelete:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "deleteForward")
			case termbox.KeyEnter:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, "enter")
			case termbox.KeySpace:
				kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, " ")
			default:
				if ev.Ch != 0 {
					kbh.KeyboardChannel <- NewKeyboardEvent(KeyEventTypeChar, PressKey, strconv.Itoa(int(ev.Ch)))
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
