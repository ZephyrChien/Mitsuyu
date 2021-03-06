package terminal

import (
	ui "github.com/gizak/termui/v3"
)

func (t *Terminal) renderInfo(idch chan string) {
	ch := t.manager.GetRecorder().GetChan()
	info := t.element.info
	infoData := t.element.infoData
	ctrl := 5
	for {
		select {
		case msg := <-ch:
			infoData = append(infoData, msg)
			info.Rows = infoData
			if ctrl == 5 {
				info.ScrollBottom()
			} else {
				ctrl++
			}
		case id := <-idch:
			ctrl = 0
			switch id {
			case "<Up>", "<MouseWheelUp>":
				info.ScrollUp()
			case "<Down>", "<MouseWheelDown>":
				info.ScrollDown()
			}
		}
		ui.Render(t.grid)
	}
}
