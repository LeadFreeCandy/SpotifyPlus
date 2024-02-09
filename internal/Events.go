package internal

import "fmt"

type EventHandler[EventData any] interface {
	Handle(EventData)
}

type Event[EventData any] interface {
	Subscribe(EventHandler[EventData]) error
	Unsubscribe(EventHandler[EventData]) error
	Trigger(EventData)
}

type GenericEventEmitter[EventData any] struct {
	handlers map[EventHandler[EventData]]struct{}
}

func (g *GenericEventEmitter[EventData]) Subscribe(handler EventHandler[EventData]) error {

	_, alreadyExists := g.handlers[handler]
	if alreadyExists {
		//TODO Improve the error messages
		return fmt.Errorf("handler already exists")
	}
	g.handlers[handler] = struct{}{}
	return nil
}

func (g *GenericEventEmitter[EventData]) Unsubscribe(handler EventHandler[EventData]) error {
	delete(g.handlers, handler)
	return nil
}

func (g *GenericEventEmitter[EventData]) Trigger(data EventData) {
	for handler := range g.handlers {
		handler.Handle(data)
	}
}
