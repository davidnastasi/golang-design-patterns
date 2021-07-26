package example1

type Result int

const (
	WIN Result = iota
	LOOSE
	DRAW
)

type Rook struct {}
type Paper struct {}
type Scissor struct {}

type Visitor interface {
  	visitRook(*Rook) Result
	visitPaper(*Paper) Result
	visitScissor(*Scissor) Result
}

type Visitable interface {
  	Accept(Visitor) Result
}

func (r *Rook) Accept(v Visitor) Result {
	return v.visitRook(r)
}

func (p *Paper) Accept(v Visitor) Result {
	return v.visitPaper(p)
}

func (s *Scissor) Accept(v Visitor) Result {
	return v.visitScissor(s)
}


type RookVisitor struct {}

func (r *RookVisitor) visitRook(rook *Rook) Result {
	return DRAW
}
func (r *RookVisitor) visitPaper(paper *Paper) Result {
	return LOOSE
}

func (r *RookVisitor) visitScissor(scissor *Scissor) Result {
	return WIN
}

type PaperVisitor struct {}

func (p *PaperVisitor) visitRook(rook *Rook) Result {
	return WIN
}
func (p *PaperVisitor) visitPaper(paper *Paper) Result {
	return DRAW
}

func (p *PaperVisitor) visitScissor(scissor *Scissor) Result {
	return LOOSE
}

type ScissorVisitor struct {}

func (s *ScissorVisitor) visitRook(rook *Rook) Result {
	return WIN
}
func (s *ScissorVisitor) visitPaper(paper *Paper) Result {
	return DRAW
}

func (s *ScissorVisitor) visitScissor(scissor *Scissor) Result {
	return LOOSE
}




