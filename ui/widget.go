package ui

// Widget é o contrato entre um nó de árvore e seu pai: o pai pergunta
// quanto ele mede (Measure), decide um espaço, manda se posicionar nele
// (Layout) e pede pra se desenhar (Draw).
type Widget interface {
	Draw(*Canvas)
	GetMargin() Edge
	Layout(Rect)
	Measure() Size
}
