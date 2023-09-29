package main

import (
	"bytes"
	"federated-learning-project/proto"
	"fmt"
	"io"
	"net/http"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

type Hello struct{ Who string }
type SetBehaviorActor struct {
	behavior actor.Behavior
}

type TrainingActor struct {
	behavior actor.Behavior
}

type Train struct {
}

type WeightsBiases struct {
	Layer_1_weights [13][64]float64
	Layer_1_biases  []float64
	Layer_2_weights [64][128]float64
	Layer_2_biases  []float64
	Layer_3_weights [128][128]float64
	Layer_3_biases  []float64
	Layer_4_weights [128][64]float64
	Layer_4_biases  []float64
	Layer_5_weights [64][32]float64
	Layer_5_biases  []float64
	Layer_6_weights [32][1]float64
	Layer_6_biases  []float64
}

func (state *SetBehaviorActor) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *TrainingActor) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *SetBehaviorActor) One(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
		state.behavior.Become(state.Other)
	}
}

func (state *SetBehaviorActor) Other(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("%v, ey we are now handling messages in another behavior", msg.Who)
	}
}

func NewSetBehaviorActor() actor.Actor {
	act := &SetBehaviorActor{
		behavior: actor.NewBehavior(),
	}
	act.behavior.Become(act.One)
	return act
}

func NewTrainingActor() actor.Actor {
	act := &TrainingActor{}

	return act
}

func NewSetTrainingActorBehavior() actor.Actor {
	actor := &TrainingActor{
		behavior: actor.NewBehavior(),
	}

	actor.behavior.Become(actor.Trainable)
	return actor
}

func (state *TrainingActor) Trainable(context actor.Context) {
	switch msg := context.Message().(type) {
	case *proto.TrainReq:
		fmt.Println("Training started")

		client := &http.Client{}

		data := bytes.NewBuffer(msg.MarshalledWeights)

		req, err := http.NewRequest("POST", "http://127.0.0.1:5000/train", data)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")

		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		content, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Print(string(content))

		var senderPID actor.PID
		senderPID.Address = msg.SenderAddress
		senderPID.Id = msg.SenderId

		context.Send(&senderPID, &proto.TrainResp{
			Data: content,
		})

	}
}

func (state *TrainingActor) Training(context actor.Context) {
	switch msg := context.Message().(type) {
	case *proto.TrainReq:
		fmt.Printf("Training is already in motion: sent by %v \n", msg.SenderId)

	}
}

func main() {
	system := actor.NewActorSystem()
	remoteConfiguration := remote.Configure("192.168.188.14", 8091)
	remoting := remote.NewRemote(system, remoteConfiguration)
	remoting.Start()

	remoting.Register("training_actor", actor.PropsFromProducer(NewSetTrainingActorBehavior))
	console.ReadLine()
}
