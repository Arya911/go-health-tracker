package tui

import (
	"fmt"
	"strings"

	"github.com/guptarohit/asciigraph"
)

func (d dashboard) View() string {

	title := titleStyle.Render("Health Checker")

	header := fmt.Sprintf(
		"%-25s %-10s %-10s %-20s",
		"SERVICE", "STATUS", "LATENCY", "LAST CHECK",
	)
	header = headerStyle.Render(header)

	rows := []string{header}

	for _, r := range d.results {

		var status string
		switch r.Status {
		case "UP":
			status = upStyle.Render("UP")
		case "DOWN":
			status = downStyle.Render("DOWN")
		default:
			status = slowStyle.Render("SLOW")
		}

		row := fmt.Sprintf(
			"%-25s %-10s %-10v %-20s",
			r.URL,
			status,
			r.Latency,
			r.Checked.Format("15:04:05"),
		)

		rows = append(rows, cellStyle.Render(row))
	}

	table := strings.Join(rows, "\n")
	graph := ""

	if d.selected != "" {
		data := d.history[d.selected]

		if len(data) > 1 {
			g := asciigraph.Plot(
				data,
				asciigraph.Height(10),
				asciigraph.Caption("Latency (ms) → "+d.selected),
			)
			graph = "\n" + g
		}
	}

	footer := "\n↑  select service, q to quit"

	content := title + "\n\n" + table + "\n" + graph + footer

	return boxStyle.Render(content)

}
