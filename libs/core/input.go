package core

type Input interface {
	GetBindValue(bind string) float64
	SetBindValue(bind string, value float64)
	Reset()
}

type InputImp struct {
	binds map[string]float64
}

func NewInput() Input {
	return &InputImp{
		binds: map[string]float64{},
	}
}

func (i *InputImp) GetBindValue(bind string) float64 {
	return i.binds[bind]
}

func (i *InputImp) SetBindValue(bind string, value float64) {
	i.binds[bind] = value
}

func (i *InputImp) Reset() {
	for key := range i.binds {
		i.binds[key] = 0.0
	}
}
