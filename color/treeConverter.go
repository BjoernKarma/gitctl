package color

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

// lipgloss.Color --> https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
// 0-7: standard colors (as in ESC [ 30–37 m)
const (
	GREEN      = lipgloss.Color("2")
	YELLOW     = lipgloss.Color("3")
	RED        = lipgloss.Color("1")
	BLUE       = lipgloss.Color("26")
	FolderIcon = "\U0001F4C1 "
)

func AddRepositoryPathToTree(root *tree.Tree, path string) *tree.Tree {
	repoPath := filepath.Clean(path)

	parts := strings.Split(repoPath, string(os.PathSeparator))

	if root == nil {
		root = tree.New().Root(FolderIcon + parts[0])
	}

	current := root

	for _, part := range parts[1:] { // Skip the root part
		found := false
		for i := 0; i < current.Children().Length(); i++ {
			child := current.Children().At(i)
			if child.Value() == FolderIcon+part {
				current = child.(*tree.Tree)
				found = true
				break
			}
		}
		if !found {
			newChild := tree.New().Root(FolderIcon + part)
			current.Child(newChild)
			current = newChild
		}
	}

	return root
}

func ConvertRepositoryPathToTree(path string, color lipgloss.TerminalColor) *tree.Tree {
	enumeratorStyle := lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)
	rootStyle := lipgloss.NewStyle().Foreground(color)
	itemStyle := lipgloss.NewStyle().Foreground(color)

	parts := strings.Split(path, string(os.PathSeparator))
	t := tree.New().Root(FolderIcon + parts[0]).
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(rootStyle).
		ItemStyle(itemStyle)
	current := t

	for _, part := range parts[1:] {
		newChild := tree.New().Root(FolderIcon + part)
		current.Child(newChild).
			Enumerator(tree.RoundedEnumerator).
			EnumeratorStyle(enumeratorStyle).
			RootStyle(rootStyle).
			ItemStyle(itemStyle)
		current = newChild
	}

	return t
}
