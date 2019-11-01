package gui

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Tree struct {
	*tview.TreeView
	OriginRoot *tview.TreeNode
}

func NewTree() *Tree {
	t := &Tree{
		TreeView: tview.NewTreeView(),
	}

	t.SetBorder(true).SetTitle("json tree").SetTitleAlign(tview.AlignLeft)
	return t
}

func (t *Tree) UpdateView(g *Gui, i interface{}) {
	g.App.QueueUpdateDraw(func() {
		r := reflect.ValueOf(i)

		var root *tview.TreeNode
		switch r.Kind() {
		case reflect.Map:
			root = tview.NewTreeNode("{object}").SetReference(Object)
		case reflect.Slice:
			root = tview.NewTreeNode("{array}").SetReference(Array)
		default:
			root = tview.NewTreeNode("{value}").SetReference(Key)
		}

		root.SetChildren(t.AddNode(i))
		t.SetRoot(root).SetCurrentNode(root)

		originRoot := *root
		t.OriginRoot = &originRoot
	})
}

func (t *Tree) AddNode(node interface{}) []*tview.TreeNode {
	var nodes []*tview.TreeNode

	switch node := node.(type) {
	case map[string]interface{}:
		for k, v := range node {
			newNode := t.NewNodeWithLiteral(k).
				SetColor(tcell.ColorMediumSlateBlue).
				SetChildren(t.AddNode(v))
			r := reflect.ValueOf(v)

			if r.Kind() == reflect.Slice {
				newNode.SetReference(Array)
			} else if r.Kind() == reflect.Map {
				newNode.SetReference(Object)
			} else {
				newNode.SetReference(Key)
			}

			log.Printf("key:%v value:%v value_kind:%v", k, v, newNode.GetReference())
			nodes = append(nodes, newNode)
		}
	case []interface{}:
		for _, v := range node {
			switch n := v.(type) {
			case map[string]interface{}:
				r := reflect.ValueOf(n)
				if r.Kind() != reflect.Slice {
					objectNode := tview.NewTreeNode("{object}").
						SetChildren(t.AddNode(v)).SetReference(Object)

					log.Printf("value:%v value_kind:%v", v, "object")
					nodes = append(nodes, objectNode)
				}
			default:
				nodes = append(nodes, t.AddNode(v)...)
			}
		}
	default:
		log.Printf("value:%v value_kind:%v", node, "value")
		nodes = append(nodes, t.NewNodeWithLiteral(node).SetReference(Value))
	}
	return nodes
}

func (t *Tree) NewNodeWithLiteral(i interface{}) *tview.TreeNode {
	return tview.NewTreeNode(fmt.Sprintf("%v", i))
}

func (t *Tree) SetKeybindings(g *Gui) {
	t.SetSelectedFunc(func(node *tview.TreeNode) {
		nodeType := node.GetReference().(Type)
		if nodeType == Root || nodeType == Object {
			return
		}
		g.Input(node.GetText(), "filed", func(text string) {
			node.SetText(text)
		})
	})

	t.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			t.GetCurrentNode().SetExpanded(false)
		case 'H':
			t.GetRoot().CollapseAll()
		case 'd':
			t.GetCurrentNode().ClearChildren()
		case 'L':
			t.GetRoot().ExpandAll()
		case 'l':
			t.GetCurrentNode().SetExpanded(true)
		case 'r':
			g.LoadJSON()
		case 's':
			g.SaveJSON()
		case '/':
			g.Search()
		}

		return event
	})
}
