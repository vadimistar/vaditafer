package taf

import (
	"fmt"
	"strings"
)

func (w *Wind) String() string {
	if w.Speed == 0 {
		return "штиль"
	}
	if w.Gusts == 0 {
		return fmt.Sprintf("%s %d м/c", w.Direction, w.Speed)
	}
	return fmt.Sprintf("%s %d м/c (порывы %d м/c)", w.Direction, w.Speed, w.Gusts)
}

func (c *CloudLayer) String() string {
	if c.Kind != "" {
		return fmt.Sprintf("%s %s, %d м", c.Quantity, c.Kind, c.Height)
	}
	return fmt.Sprintf("%s, %d м", c.Quantity, c.Height)
}

func (c *ChangeHeader) String() string {
	var s strings.Builder
	if c.Probability != 0 {
		fmt.Fprintf(&s, "с вероятностью %d%% ", c.Probability)
	}
	if c.Kind != "" {
		s.WriteString(c.Kind + " ")
	}
	s.WriteString("с " + c.Start.Format(TimeLayout) + " ")
	if !c.End.IsZero() {
		s.WriteString("до " + c.End.Format(TimeLayout) + " ")
	}
	res := s.String()
	return res[:len(res)-1]
}

var TimeLayout = "15:04"
