package xgo

import (
	"context"
	"errors"
)

var ErrorDone = errors.New("done")

type Message interface{}
type Iterator struct {
	retCh chan Message
	errCh chan error
	cancel context.CancelFunc
}

// NewIterator 创建一个迭代器
func NewIterator(retCh chan Message, errCh chan error) (*Iterator, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Iterator{
		retCh:  retCh,
		errCh:  errCh,
		cancel: cancel,
	}, ctx
}

func (i *Iterator) Next() (Message, error) {
	var (
		msg Message
		err error
		ok  bool
	)
	select {
	case msg, ok = <-i.retCh:
	default:
		select {
		case msg, ok = <-i.retCh:
		case err, ok = <-i.errCh:
		}
	}
	if !ok {
		return nil, ErrorDone
	}
	return msg, err
}

func (i *Iterator) Close() {
	i.cancel()
}
