package actors

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"

	console "github.com/asynkron/goconsole"
)

type InitializerActor struct {
	selfPID        *actor.PID
	coordinatorPID *actor.PID
	aggregatorPID  *actor.PID
}

type CoordinatorActor struct {
	selfPID       *actor.PID
	parentPID     *actor.PID
	aggregatorPID *actor.PID
}

type AggregatorActor struct {
	selfPID   *actor.PID
	parentPID *actor.PID
}

type PidsStruct struct {
	initPID        *actor.PID
	coordinatorPID *actor.PID
	aggregatorPID  *actor.PID
}

func newInitializatorActor() actor.Actor {
	return &InitializerActor{}
}

func newCoordinatorActor() actor.Actor {
	return &CoordinatorActor{}
}

func newAggregatorActor() actor.Actor {
	return &AggregatorActor{}
}

func (state *InitializerActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {

	case *actor.PID:
		coordinatorProps := actor.PropsFromProducer(newCoordinatorActor)
		coordinatorPID := context.Spawn(coordinatorProps)
		aggregatorProps := actor.PropsFromProducer(newAggregatorActor)
		aggregatorPID := context.Spawn(aggregatorProps)

		state.selfPID = msg
		state.coordinatorPID = coordinatorPID
		state.aggregatorPID = aggregatorPID

		context.Send(coordinatorPID, PidsStruct{initPID: state.selfPID, coordinatorPID: state.coordinatorPID, aggregatorPID: state.aggregatorPID})

	}
}

func (state *CoordinatorActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case PidsStruct:
		if msg.initPID == nil {
			state.parentPID = msg.initPID
			state.selfPID = msg.coordinatorPID
			state.aggregatorPID = msg.aggregatorPID
		}
	}
}

func (state *AggregatorActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case PidsStruct:
		if msg.initPID == nil {
			state.parentPID = msg.initPID
			state.selfPID = msg.aggregatorPID
		}
	}
}

func main() {
	system := actor.NewActorSystem()
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.RestartDirective
	}

	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	rootContext := system.Root
	props := actor.
		PropsFromProducer(newInitializatorActor,
			actor.WithSupervisor(supervisor))
	pid := rootContext.Spawn(props)

	rootContext.Send(pid, pid)

	_, _ = console.ReadLine()
}
