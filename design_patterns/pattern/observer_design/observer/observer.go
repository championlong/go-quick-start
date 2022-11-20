package observer

import "fmt"

type Observer interface {
	update(message string)
}

type Subject interface {
	registerObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers(message string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (self *ConcreteSubject) registerObserver(observer Observer) {
	self.observers = append(self.observers, observer)
}

func (self *ConcreteSubject) removeObserver(observer Observer) {
	for i, ob := range self.observers {
		if ob == observer {
			self.observers = append(self.observers[:i], self.observers[i+1:]...)
		}
	}
}

func (self *ConcreteSubject) notifyObservers(message string) {
	for _, v := range self.observers {
		v.update(message)
	}
}

type ConcreteObserverOne struct {
}

func (self *ConcreteObserverOne) update(message string) {
	fmt.Println("ConcreteObserverOne", message)
}

type ConcreteObserverTwo struct {
}

func (self *ConcreteObserverTwo) update(message string) {
	fmt.Println("ConcreteObserverTwo", message)
}
