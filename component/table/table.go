// Taken from the Bubble library.

package table

import (
	"strconv"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/mattn/go-runewidth"
)

// Props holds the configuration for the Table component.
type Props struct {
	DataFunc func(c *app.Ctx) (clms []Column, rows []Row)
	KeyMap   KeyMap
	Styles   Styles
	Help     help.Model
	app.Margin
	app.Layout
}

type tableProp func(*Props)

// tableState holds the internal state of the Table component.
type tableState struct {
	cols     []column
	cursor   int
	viewport viewport.Model
}

type Row []string

type ColumnWidth struct {
	Int  int
	Grow bool
}

type Column struct {
	Title string
	Width ColumnWidth
}

func WidthGrow() ColumnWidth {
	return ColumnWidth{Grow: true}
}
func WidthInt(i int) ColumnWidth {
	return ColumnWidth{Int: i}
}

// column is an internal representation with calculated width.
type column struct {
	Title string
	Width int
}

type KeyMap struct {
	LineUp       key.Binding
	LineDown     key.Binding
	PageUp       key.Binding
	PageDown     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	GotoTop      key.Binding
	GotoBottom   key.Binding
}

func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.LineUp, km.LineDown}
}

func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.LineUp, km.LineDown, km.GotoTop, km.GotoBottom},
		{km.PageUp, km.PageDown, km.HalfPageUp, km.HalfPageDown},
	}
}

func defaultKeyMap() KeyMap {
	return KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("b", "pgup"),
			key.WithHelp("b/pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("f", "pgdown", "space"),
			key.WithHelp("f/pgdn", "page down"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("u", "ctrl+u"),
			key.WithHelp("u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("d", "ctrl+d"),
			key.WithHelp("d", "½ page down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GotoBottom: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
	}
}

type Styles struct {
	Base      lipgloss.Style
	BaseFocus lipgloss.Style
	Header    lipgloss.Style
	Cell      lipgloss.Style
	Selected  lipgloss.Style
	Hovered   lipgloss.Style // Currently unused
}

// TODO: Add this to the theme
func defaultStyles(c *app.Ctx) Styles {
	base := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, true, true, true).BorderForeground(c.Theme.Colors.Base600)
	return Styles{
		Base:      base,
		BaseFocus: base.BorderForeground(c.Theme.Colors.Base50),
		Selected:  lipgloss.NewStyle().Bold(true).Foreground(c.Theme.Colors.PrimaryLight).Background(c.Theme.Colors.Base700),
		Hovered:   lipgloss.NewStyle().Bold(true).Foreground(c.Theme.Colors.PrimaryLight).Background(c.Theme.Colors.Base600),
		Header:    lipgloss.NewStyle().Bold(true).BorderStyle(lipgloss.NormalBorder()).BorderForeground(c.Theme.Colors.Base600).BorderBottom(true),
		Cell:      lipgloss.NewStyle(),
	}
}

// NewTable creates a new table component instance.
func New(c *app.Ctx, props ...tableProp) *app.C {
	p := Props{
		KeyMap: defaultKeyMap(),
		Styles: defaultStyles(c),
		Help:   help.New(),
		Layout: app.Layout{
			GrowX: true,
			GrowY: true,
		},
	}
	for _, prop := range props {
		if prop != nil {
			prop(&p)
		}
	}
	p.Styles.Base = app.ApplyMargin(p.Styles.Base, p.Margin)
	p.Styles.BaseFocus = app.ApplyMargin(p.Styles.BaseFocus, p.Margin)

	return c.Render(Table, p)
}

func WithDataFunc(f func(c *app.Ctx) (clms []Column, rows []Row)) tableProp {
	return func(props *Props) {
		props.DataFunc = f
	}
}

// Table is the functional component for rendering a table.
func Table(c *app.Ctx, props app.Props) string {
	p, _ := props.(Props)
	id := app.UseID(c)
	isFocused := app.UseIsFocused(c)
	_, childHoverID := app.UseIsHovered(c)

	state, setState := app.UseState(c, tableState{
		cursor:   -1,
		viewport: viewport.New(),
	})

	rawCols, rows := p.DataFunc(c)

	app.UseKeyHandler(c, func(keyMsg tea.KeyMsg) bool {
		return processInternalKeys(keyMsg, p.KeyMap, rows, state, func(t tableState) {
			setState(t)
		})
	})

	app.UseMouseHandler(c, func(msg tea.MouseMsg, childID string) bool {
		// TODO: Add click for headers to sort
		if childID == "" || msg.Mouse().Button != tea.MouseLeft {
			return false
		}
		if _, ok := msg.(tea.MouseReleaseMsg); ok {

			rowIDFromChild := strings.Split(childID, ":")
			if rowIDFromChild[0] == "row" {
				rowIndex, err := strconv.Atoi(rowIDFromChild[1])
				if err != nil {
					return false
				}
				if rowIndex >= 0 && rowIndex < len(rows) {
					state.cursor = rowIndex
					setState(state)
				}
			}
			return true
		}
		return false
	})

	width, height := app.UseSize(c)

	app.UseEffect(c, func() {
		baseStyleToUse := p.Styles.Base
		if isFocused {
			baseStyleToUse = p.Styles.BaseFocus
		}
		state.cols = columnMapping(width-baseStyleToUse.GetHorizontalFrameSize()-(p.Styles.Header.GetHorizontalFrameSize()*len(rawCols)), rawCols)
		setState(state)
	}, []any{rawCols, width})

	numRows := len(rows)

	// Handle cursor initialization and bounds
	currentCursor := state.cursor
	updatedCursor := currentCursor

	if isFocused && updatedCursor < 0 && numRows > 0 {
		updatedCursor = 0
	}
	if numRows > 0 {
		updatedCursor = clamp(updatedCursor, 0, numRows-1)
	} else {
		updatedCursor = -1 // No rows, no cursor
	}

	if updatedCursor != currentCursor {
		newState := state
		newState.cursor = updatedCursor
		// If cursor moves due to focus/data change, ensure viewport is valid
		if numRows > 0 {
			if newState.cursor < newState.viewport.YOffset() {
				newState.viewport.SetYOffset(newState.cursor)
			} else if newState.cursor >= newState.viewport.YOffset()+newState.viewport.Height() {
				newState.viewport.SetYOffset(newState.cursor - newState.viewport.Height() + 1)
			}
		}
		setState(newState)
	}

	headersViewStr := generateHeadersView(state.cols, p.Styles)

	currentBaseStyle := p.Styles.Base
	if isFocused {
		currentBaseStyle = p.Styles.BaseFocus
	}

	state.viewport.SetHeight(height - lipgloss.Height(headersViewStr) - currentBaseStyle.GetVerticalFrameSize())
	state.viewport.SetWidth(width - currentBaseStyle.GetHorizontalFrameSize())

	updateViewportContent(&state.viewport, rows, state, childHoverID, p.Styles, c, id)

	return currentBaseStyle.Render(headersViewStr + "\n" + state.viewport.View())
}

// processInternalKeys contains the logic for handling key presses for table navigation.
func processInternalKeys(keyMsg tea.KeyMsg, km KeyMap, rows []Row, currentTableState tableState, setState func(tableState)) bool {
	numRows := len(rows)

	if numRows == 0 && !(key.Matches(keyMsg, km.LineUp) || key.Matches(keyMsg, km.LineDown)) {
		if key.Matches(keyMsg, km.LineUp) || key.Matches(keyMsg, km.LineDown) {
			return true
		}
		return false
	}

	switch {
	case key.Matches(keyMsg, km.LineUp):
		moveUp(currentTableState, setState, 1, numRows)
	case key.Matches(keyMsg, km.LineDown):
		moveDown(currentTableState, setState, 1, numRows)
	case key.Matches(keyMsg, km.PageUp):
		moveUp(currentTableState, setState, currentTableState.viewport.Height(), numRows)
	case key.Matches(keyMsg, km.PageDown):
		moveDown(currentTableState, setState, currentTableState.viewport.Height(), numRows)
	case key.Matches(keyMsg, km.HalfPageUp):
		moveUp(currentTableState, setState, currentTableState.viewport.Height()/2, numRows)
	case key.Matches(keyMsg, km.HalfPageDown):
		moveDown(currentTableState, setState, currentTableState.viewport.Height()/2, numRows)
	case key.Matches(keyMsg, km.GotoTop):
		gotoTop(currentTableState, setState, numRows)
	case key.Matches(keyMsg, km.GotoBottom):
		gotoBottom(currentTableState, setState, numRows)
	default:
		return false // No key matched
	}
	return true
}
func calculateClampedCursorValue(currentCursor, delta, numRows int) int {
	if numRows == 0 {
		return -1
	}
	newCursor := currentCursor + delta
	return clamp(newCursor, 0, numRows-1)
}

func moveUp(currentState tableState, setState func(tableState), n int, numRows int) {
	if numRows == 0 {
		return
	}
	newCursor := calculateClampedCursorValue(currentState.cursor, -n, numRows)

	newYOffset := currentState.viewport.YOffset()
	if newCursor < newYOffset { // Cursor moved above viewport top
		newYOffset = newCursor
	}

	if newCursor != currentState.cursor || newYOffset != currentState.viewport.YOffset() {
		state := currentState
		state.cursor = newCursor
		state.viewport.SetYOffset(newYOffset) // Modifies the viewport copy in state
		setState(state)
	}
}

func moveDown(currentState tableState, setState func(tableState), n int, numRows int) {
	if numRows == 0 {
		return
	}
	newCursor := calculateClampedCursorValue(currentState.cursor, n, numRows)

	newYOffset := currentState.viewport.YOffset()
	// Cursor moved below viewport bottom
	if newCursor >= newYOffset+currentState.viewport.Height() {
		newYOffset = newCursor - currentState.viewport.Height() + 1
	}

	if newCursor != currentState.cursor || newYOffset != currentState.viewport.YOffset() {
		state := currentState
		state.cursor = newCursor
		state.viewport.SetYOffset(newYOffset)
		setState(state)
	}
}

func gotoTop(currentState tableState, setState func(tableState), numRows int) {
	if numRows == 0 {
		return
	}
	newCursor := calculateClampedCursorValue(0, 0, numRows) // Effectively clamp(0,0,numRows-1)

	if newCursor != currentState.cursor || currentState.viewport.YOffset() != 0 {
		state := currentState
		state.cursor = newCursor
		state.viewport.GotoTop() // Modifies the viewport copy in state
		setState(state)
	}
}

func gotoBottom(currentState tableState, setState func(tableState), numRows int) {
	if numRows == 0 {
		return
	}
	newCursor := calculateClampedCursorValue(numRows-1, 0, numRows)

	// Calculate the YOffset that would make the last item visible.
	// This logic is tricky with GotoBottom() as it scrolls to the very end of content.
	// We need to ensure the viewport's YOffset is correctly set if it changes.
	tempViewport := currentState.viewport // Copy to simulate GotoBottom
	tempViewport.GotoBottom()             // This sets YOffset to show the end of content.
	newYOffset := tempViewport.YOffset()

	if newCursor != currentState.cursor || newYOffset != currentState.viewport.YOffset() {
		state := currentState
		state.cursor = newCursor
		state.viewport.SetYOffset(newYOffset) // Use the calculated YOffset
		setState(state)
	}
}

func generateHeadersView(cols []column, styles Styles) string {
	s := make([]string, 0, len(cols))
	for _, col := range cols {
		if col.Width <= 0 {
			continue
		}
		style := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width)
		renderedCell := style.Render(runewidth.Truncate(col.Title, col.Width, "…"))
		s = append(s, styles.Header.Render(renderedCell))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, s...)
}

func generateRenderedRow(rowIndex int, rowData Row, state tableState, childHoverID string, styles Styles, c *app.Ctx, tableID string) string {
	rowStyle := styles.Cell

	s := make([]string, 0, len(state.cols))
	for i, value := range rowData {
		if i >= len(state.cols) || state.cols[i].Width <= 0 {
			continue
		}
		rowClmStyle := styles.Cell.Width(state.cols[i].Width).MaxWidth(state.cols[i].Width)
		s = append(s, rowClmStyle.Render(runewidth.Truncate(value, state.cols[i].Width, "…")))
	}
	rowElementID := "row:" + strconv.Itoa(rowIndex)

	if rowElementID == childHoverID {
		rowStyle = rowStyle.Inherit(styles.Hovered)
	} else if rowIndex == state.cursor {
		rowStyle = rowStyle.Inherit(styles.Selected)
	}

	return c.MouseZoneChild(rowElementID, rowStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, s...)))
}

func updateViewportContent(vp *viewport.Model, rows []Row, state tableState, childHoverID string, styles Styles, c *app.Ctx, tableID string) {
	if len(rows) == 0 {
		vp.SetContent("")
		return
	}

	renderedRows := make([]string, len(rows))
	for i, rowData := range rows {
		renderedRows[i] = generateRenderedRow(i, rowData, state, childHoverID, styles, c, tableID)
	}

	vp.SetContent(
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
	)
}

func columnMapping(width int, clms []Column) []column {
	numberOfGrowers := 0
	sizeOfStatic := 0
	for _, clm := range clms {
		if clm.Width.Grow {
			numberOfGrowers++
		} else {
			sizeOfStatic += clm.Width.Int
		}
	}

	growWidth := width - sizeOfStatic
	baseSizePerGrower := 0
	remainder := 0

	if numberOfGrowers > 0 {
		if growWidth < 0 {
			growWidth = 0
		}
		baseSizePerGrower = growWidth / numberOfGrowers
		remainder = growWidth % numberOfGrowers
	}

	columns := make([]column, len(clms))
	for i, clm := range clms {
		colWidth := 0
		if clm.Width.Grow {
			colWidth = baseSizePerGrower
			if remainder > 0 {
				colWidth++
				remainder--
			}
		} else {
			colWidth = clm.Width.Int
		}

		if colWidth < 0 {
			colWidth = 0
		}
		columns[i] = column{
			Title: clm.Title,
			Width: colWidth,
		}
	}
	return columns
}

func clamp(v, low, high int) int {
	if high < low {
		return min(max(v, low), high)
	}
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
