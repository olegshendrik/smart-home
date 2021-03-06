package core

type MapBind struct {
	Map *Map
}

func (e *MapBind) SetElementState(device *DeviceBind, state string) {
	e.Map.SetElementState(device.model, state)
}

func (e *MapBind) GetElement(device *DeviceBind) *MapElementBind {
	return &MapElementBind{element:e.Map.GetElement(device.model)}
}
