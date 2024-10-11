package api

import (
	"errors"
	"golang.org/x/term"
	"os"
)

type Keyboard struct{}

type KeyBase int

const (
	KEY_ARROW_UP KeyBase = iota
	KEY_ARROW_DOWN
	KEY_ARROW_LEFT
	KEY_ARROW_RIGHT
	NONE
)

func (k *Keyboard) GetSingleKey() (KeyBase, error) {
	// makes terminal press key without hit 'enter'
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		return NONE, errors.New("fail to init")
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// read typed
	b := make([]byte, 4)

	_, err = os.Stdin.Read(b)

	if err != nil {
		return NONE, errors.New("fail to read")
	}

	var keyMap = map[[4]byte]KeyBase{
		{27, 91, 65, 0}: KEY_ARROW_UP,
		{27, 91, 66, 0}: KEY_ARROW_DOWN,
		{27, 91, 67, 0}: KEY_ARROW_RIGHT,
		{27, 91, 68, 0}: KEY_ARROW_LEFT,
	}

	if key, ok := keyMap[[4]byte(b)]; ok {
		return key, nil
	}

	return NONE, nil
}
