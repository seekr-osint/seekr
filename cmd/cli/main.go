package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/seekr-osint/seekr/api/client"
	"github.com/seekr-osint/seekr/api/functions"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cristalhq/acmd"
)

// globals
var (
	ID string
)

func renderMarkdown(md string) error {
	out, err := glamour.Render(md, "dark")
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
	return nil
}

func ShowPerson(c *client.Client, id string) {
	person, err := c.GetPerson(id)
	if err != nil {
		panic(err)
	}
	err = renderMarkdown(person.Markdown())
	if err != nil {
		panic(err)
	}
}

func search(c *client.Client, ctx context.Context, args []string) error {
	SelectPersonView(c)
	return nil
}

func person(c *client.Client, ctx context.Context, args []string) error {
	if args[0] != "" {
		ShowPerson(c, args[0])
	} else {
		SelectPersonView(c)
		//fmt.Println("error arguments")
	}

	return nil
}

func ping(c *client.Client, ctx context.Context, args []string) error {
	ping, err := c.Ping()
	if err != nil {
		return err
	}
	fmt.Printf("ping: %s\n", ping)
	return nil
}
func main() {
	c := client.NewClient("localhost", 8569)
	cmds := []acmd.Command{
		{
			Name:        "person",
			Description: "show person",
			ExecFunc: func(ctx context.Context, args []string) error {
				return person(c, ctx, args)
			},
		},
		{
			Name:        "ping",
			Description: "ping api",
			ExecFunc: func(ctx context.Context, args []string) error {
				return person(c, ctx, args)
			},
		},
		{
			Name:        "search",
			Description: "search person",
			ExecFunc: func(ctx context.Context, args []string) error {
				return person(c, ctx, args)
			},
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "cli",
		AppDescription: "seekr cli",
		Version:        "0.4.0",
		// Context - if nil `signal.Notify` will be used
		// Args - if nil `os.Args[1:]` will be used
		// Usage - if nil default print will be used
	})
	if err := r.Run(); err != nil {
		r.Exit(err)
	}
}

func SelectPersonDelete(c *client.Client) {
	people, err := c.GetPeople()
	if err != nil {
		panic(err)
	}
	ShowList(MapToList(people))
	c.DeletePerson(ID)
}

func SelectPersonView(c *client.Client) {
	people, err := c.GetPeople()
	if err != nil {
		panic(err)
	}
	ShowList(MapToList(people))
	ShowPerson(c, ID)
}

func MapToList(mp map[string]string) []list.Item {
	lst := []list.Item{}
	for _, key := range functions.SortMapKeys(mp) {
		it := fmt.Sprintf("%s (%s)", mp[key], key)
		lst = append(lst, item(it))
	}
	return lst
}

func extractID(input string) (string, error) {
	lastOpenBracket := strings.LastIndex(input, "(")
	if lastOpenBracket == -1 {
		return "", fmt.Errorf("opening bracket not found")
	}

	id := input[lastOpenBracket+1 : len(input)-1]

	return id, nil
}

// ui

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		id, err := extractID(m.choice)
		if err != nil {
			panic(err)
		}
		ID = id

		return quitTextStyle.Render(fmt.Sprintf("Selected: %s.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Nothing selected")
	}
	return "\n" + m.list.View()
}

func ShowList(items []list.Item) {
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select person?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
