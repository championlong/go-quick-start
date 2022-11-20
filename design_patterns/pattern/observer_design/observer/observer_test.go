package observer

import "testing"

func TestObserverTemplate(t *testing.T) {
	concreteSubject := ConcreteSubject{}
	concreteSubject.registerObserver(&ConcreteObserverOne{})
	concreteSubject.registerObserver(&ConcreteObserverTwo{})
	concreteSubject.notifyObservers("通知")
}