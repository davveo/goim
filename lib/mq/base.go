package mq

type Mq interface {
	Push()
	Pull()
}

type DefaultMq struct {

}

func (d DefaultMq) Push() {
	panic("implement me")
}

func (d DefaultMq) Pull() {
	panic("implement me")
}