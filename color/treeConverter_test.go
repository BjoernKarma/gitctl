package color

import (
	"path/filepath"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func TestAddRepositoryPathToTree_MergeCommonParts(t *testing.T) {
	source := tree.New().Root(FolderIcon + "usr").
		Child(
			tree.New().
				Root(FolderIcon + "dev").
				Child(
					tree.New().
						Root(FolderIcon + "github")))

	combinedTree := tree.New().Root(FolderIcon + "usr").
		Child(
			tree.New().
				Root(FolderIcon + "dev").
				Child(FolderIcon + "github").Child(FolderIcon + "middleware"))

	var path = filepath.Join("usr", "dev", "middleware")
	AddRepositoryPathToTree(source, path)

	if source.String() != combinedTree.String() {
		t.Errorf("expected merged tree to be \n%v \n\ngot\n \n%v", combinedTree, source)
	}
}

func TestAddRepositoryPathToTree_EmptyRoot(t *testing.T) {
	var source *tree.Tree

	var path = filepath.Join("usr", "dev", "middleware")
	source = AddRepositoryPathToTree(source, path)

	expectedTree := tree.New().Root(FolderIcon + "usr").
		Child(
			tree.New().
				Root(FolderIcon + "dev").
				Child(FolderIcon + "middleware"))

	if source.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, source)
	}
}

func TestAddRepositoryPathToTree_EmptyPath(t *testing.T) {
	source := tree.New().Root(FolderIcon + "usr")

	AddRepositoryPathToTree(source, "")

	expectedTree := tree.New().Root(FolderIcon + "usr")

	if source.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, source)
	}
}

func TestAddRepositoryPathToTree_RootOnly(t *testing.T) {
	source := tree.New().Root(FolderIcon + "usr")

	AddRepositoryPathToTree(source, "usr")

	expectedTree := tree.New().Root(FolderIcon + "usr")

	if source.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, source)
	}
}

func TestAddRepositoryPathToTree_NestedPath(t *testing.T) {
	source := tree.New().Root(FolderIcon + "usr")

	var path = filepath.Join("usr", "dev", "github", "middleware")
	AddRepositoryPathToTree(source, path)

	expectedTree := tree.New().Root(FolderIcon + "usr").
		Child(
			tree.New().
				Root(FolderIcon + "dev").
				Child(
					tree.New().
						Root(FolderIcon + "github").
						Child(FolderIcon + "middleware")))

	if source.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, source)
	}
}

func TestConvertRepositoryPathToTree_EmptyPath(t *testing.T) {
	result := ConvertRepositoryPathToTree("", BLUE)
	expectedTree := tree.New().Root(FolderIcon).
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
		RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
		ItemStyle(lipgloss.NewStyle().Foreground(BLUE))

	if result.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, result)
	}
}

func TestConvertRepositoryPathToTree_SingleLevelPath(t *testing.T) {
	result := ConvertRepositoryPathToTree("usr", BLUE)
	expectedTree := tree.New().Root(FolderIcon + "usr").
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
		RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
		ItemStyle(lipgloss.NewStyle().Foreground(BLUE))

	if result.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, result)
	}
}

func TestConvertRepositoryPathToTree_MultiLevelPath(t *testing.T) {
	var path = filepath.Join("usr", "dev", "github", "middleware")
	result := ConvertRepositoryPathToTree(path, BLUE)
	expectedTree := tree.New().Root(FolderIcon + "usr").
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
		RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
		ItemStyle(lipgloss.NewStyle().Foreground(BLUE)).
		Child(
			tree.New().Root(FolderIcon + "dev").
				Enumerator(tree.RoundedEnumerator).
				EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
				RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
				ItemStyle(lipgloss.NewStyle().Foreground(BLUE)).
				Child(
					tree.New().Root(FolderIcon + "github").
						Enumerator(tree.RoundedEnumerator).
						EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
						RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
						ItemStyle(lipgloss.NewStyle().Foreground(BLUE)).
						Child(
							tree.New().Root(FolderIcon + "middleware").
								Enumerator(tree.RoundedEnumerator).
								EnumeratorStyle(lipgloss.NewStyle().Foreground(BLUE).MarginRight(1)).
								RootStyle(lipgloss.NewStyle().Foreground(BLUE)).
								ItemStyle(lipgloss.NewStyle().Foreground(BLUE)),
						),
				),
		)

	if result.String() != expectedTree.String() {
		t.Errorf("expected tree to be \n%v \n\ngot\n \n%v", expectedTree, result)
	}
}
