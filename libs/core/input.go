package core

import "github.com/veandco/go-sdl2/sdl"

var KEYBOARD_EVENT = "KEYBOARD_EVENT"

type KeyboardBinding struct {
	MinEventKey int
	MinValue    float32
	MaxEventKey int
	MaxValue    float32
}

type Input interface {
	GetValue(name string) float32
	RegisterInput(name string, inputBinding interface{})
	Handle(sdlEvent sdl.Event)
	Reset()
}

type InputImp struct {
	bindingMap   map[string]interface{}
	bindingValue map[string]float32
}

func NewInput() Input {
	return &InputImp{
		bindingMap:   make(map[string]interface{}),
		bindingValue: make(map[string]float32),
	}
}

func (i *InputImp) GetValue(name string) float32 {
	return i.bindingValue[name]
}

func (i *InputImp) RegisterInput(name string, inputBinding interface{}) {
	i.bindingMap[name] = inputBinding
	i.bindingValue[name] = 0.0
}

func (i *InputImp) Handle(sdlEvent sdl.Event) {
	for key, v := range i.bindingMap {
		switch t := sdlEvent.(type) {
		case *sdl.KeyboardEvent:
			if _, ok := v.(*KeyboardBinding); ok {
				kb := v.(*KeyboardBinding)
				if t.Keysym.Sym == sdl.Keycode(kb.MinEventKey) {
					i.bindingValue[key] = kb.MinValue
				} else if t.Keysym.Sym == sdl.Keycode(kb.MaxEventKey) {
					i.bindingValue[key] = kb.MaxValue
				}
			}
		}
	}
}

func (i *InputImp) Reset() {
	for key := range i.bindingValue {
		i.bindingValue[key] = 0.0
	}
}
