package app

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

// KeyHandler defines the signature for component-internal key handlers.
// It returns true if the key press was handled, false otherwise.
type KeyHandler func(keyMsg tea.KeyMsg) bool

// MouseHandler defines the signature for component-internal mouse handlers.
type MouseHandler func(msg tea.MouseMsg, childID string) bool

// MsgHandler is for receiving raw tea.Msg messages.
type MsgHandler func(msg tea.Msg) tea.Cmd

// C represents an instance of a functional component (FC).
// It holds the component's ID, focusable state, function reference, props,
// event handlers, and state management for both state and effects.
type C struct {
	id        string
	content   string
	focusable bool

	parent   *C
	children []*C
	props    Props
	layout   Layout

	// Hooks
	states            []any
	effects           []effectRecord
	keyHandlers       []KeyHandler
	globalKeyHandlers []KeyHandler
	mouseHandlers     []MouseHandler
	messageHandlers   []MsgHandler
	onFocused         func(isReverse bool)

	useEffectCounter int
	useStateCounter  int

	width, height int
	x, y          int

	// cursor *tea.Cursor
}

func (c *C) String() string {
	return c.content
}

func (c *Ctx) getCurrentComponent() *C {
	id := c.id.getID()
	instance, exists := c.getComponent(id)
	if !exists {
		panic("UseMsgHandler: component instance not found")
	}
	return instance
}

func (c *Ctx) getComponent(id string) (*C, bool) {
	instance, ok := c.components[id]
	return instance, ok
}

func (c *Ctx) initComponent(id string, props Props) *C {
	instance, ok := c.components[id]
	if !ok {
		instance = &C{
			id:               id,
			states:           make([]any, 0),
			effects:          make([]effectRecord, 0),
			mouseHandlers:    make([]MouseHandler, 0),
			keyHandlers:      make([]KeyHandler, 0),
			messageHandlers:  make([]MsgHandler, 0),
			useEffectCounter: 0,
			useStateCounter:  0,
		}
		c.components[id] = instance
	}
	instance.layout = extractLayoutFromProps(props)
	instance.props = props
	return instance
}

func (c *Ctx) cleanupEffects(removedIDs []string) {
	for _, id := range removedIDs {
		if instance, ok := c.components[id]; ok {
			for i := range instance.effects {
				if instance.effects[i].cleanupFn != nil {
					instance.effects[i].cleanupFn()
					instance.effects[i].cleanupFn = nil
				}
			}
		}
		delete(c.components, id)
	}
}

func (c *Ctx) getAllGlobalKeyHandlers() []KeyHandler {
	var handlers []KeyHandler
	ids := make([]string, 0, len(c.components))
	for id := range c.components {
		ids = append(ids, id)
	}

	for i := len(ids) - 1; i >= 0; i-- {
		instance := c.components[ids[i]]
		handlers = append(handlers, instance.globalKeyHandlers...)
	}
	return handlers
}
