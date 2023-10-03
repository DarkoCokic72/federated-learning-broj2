package main

import (
	"encoding/json"
	"federated-learning-project/proto"
	"fmt"
	"os"
	"strconv"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
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

	roundsTrained  uint
	maxRounds      uint
	actorsTraining uint

	children *actor.PIDSet
	weights  []WeightsBiases

	behavior actor.Behavior
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

type WeightsBiases struct {
	Layer_1_weights [13][64]float64
	Layer_1_biases  [64]float64
	Layer_2_weights [64][128]float64
	Layer_2_biases  [128]float64
	Layer_3_weights [128][128]float64
	Layer_3_biases  [128]float64
	Layer_4_weights [128][64]float64
	Layer_4_biases  [64]float64
	Layer_5_weights [64][32]float64
	Layer_5_biases  [32]float64
	Layer_6_weights [32][1]float64
	Layer_6_biases  [1]float64
}

type CalculateAverage struct {
	Weights []WeightsBiases
}

type UpdatedWeights struct{}

type spawnedRemoteActor struct {
	remoteActorPID *actor.PID
}

type startTraining struct{}

type weightsUpdated struct{}

type trainingFinished struct{}

func newInitializatorActor() actor.Actor {
	return &InitializerActor{}
}

func newCoordinatorActor() actor.Actor {
	return &CoordinatorActor{}
}

func newAggregatorActor() actor.Actor {
	return &AggregatorActor{}
}

func NewSetCoordinatorBehavior() actor.Actor {
	act := &CoordinatorActor{
		behavior: actor.NewBehavior(),
	}
	act.behavior.Become(act.Training)
	return act
}

func (state *InitializerActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {

	case *actor.PID:

		coordinatorProps := actor.PropsFromProducer(NewSetCoordinatorBehavior)
		coordinatorPID := context.Spawn(coordinatorProps)
		aggregatorProps := actor.PropsFromProducer(newAggregatorActor)
		aggregatorPID := context.Spawn(aggregatorProps)

		state.selfPID = msg
		state.coordinatorPID = coordinatorPID
		state.aggregatorPID = aggregatorPID

		msg1 := PidsStruct{initPID: msg, coordinatorPID: coordinatorPID, aggregatorPID: aggregatorPID}
		context.Send(coordinatorPID, msg1)

		msg2 := PidsStruct{initPID: msg, aggregatorPID: aggregatorPID}
		context.Send(aggregatorPID, msg2)

		fmt.Printf("Initialized all actors %v\n", time.Now())

	case spawnedRemoteActor:
		// Passing the message to the cooridnator actor
		fmt.Printf("Spawned remote actor %v at %v\n", msg.remoteActorPID, time.Now())
		context.Send(state.coordinatorPID, msg)
	case startTraining:
		// Passing the message to the cooridnator actor
		context.Send(state.coordinatorPID, msg)
	case weightsUpdated:
		// Once the weights are updated, let coordinator know
		context.Send(state.coordinatorPID, msg)
	case trainingFinished:
		// Once the training is finished we can serialize
		// the new weights
		file, err := os.Create("weightModel.json")

		if err != nil {
			panic(err)
		}

		defer file.Close()
		encoder := json.NewEncoder(file)

		encodingError := encoder.Encode(globalModel)
		if encodingError != nil {
			panic(err)
		}
		fmt.Println("Successfully written the weights to 'weightModel.json' file!")
	}
}

func (state *CoordinatorActor) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *CoordinatorActor) Training(context actor.Context) {
	switch msg := context.Message().(type) {
	//internal state
	case *actor.Started:
		state.children = actor.NewPIDSet()
	case PidsStruct:
		if msg.initPID == nil {
			return
		}
		state.parentPID = msg.initPID
		state.selfPID = msg.coordinatorPID
		state.aggregatorPID = msg.aggregatorPID

		state.maxRounds = 10
		state.roundsTrained = 0
		state.actorsTraining = 0
	case startTraining:
		state.weights = []WeightsBiases{}
		fmt.Printf("Starting a new round of training at %v\n", time.Now())
		// If we have reached maximum rounds of training we exit
		if state.roundsTrained >= state.maxRounds {
			fmt.Printf("Reached a maximum of %v training rounds.\n", state.maxRounds)
			msg3 := trainingFinished{}
			context.Send(state.parentPID, msg3)

			return
		}
		// Create a training reqeuest message which will be sent
		// to all the training actors
		weightsJson, marshalErr := json.Marshal(globalModel)

		if marshalErr != nil {
			panic(marshalErr)
		}

		senderAddress := localAddress + ":" + strconv.Itoa(port)
		trainMessage := &proto.TrainReq{
			SenderAddress:     senderAddress,
			SenderId:          state.selfPID.Id,
			MarshalledWeights: weightsJson,
		}
		// Send the message to all the training actors
		state.children.ForEach(func(i int, pid *actor.PID) {
			msg4 := trainMessage
			context.Send(pid, msg4)
			state.actorsTraining += 1
		})
		state.roundsTrained += 1
		fmt.Println("TRAINING ROUND: ", state.roundsTrained)
	case *proto.TrainResp:
		fmt.Printf("Got a training response %v\n", time.Now())
		// Another node finished training
		state.actorsTraining -= 1
		// Convert the weights and add it to the weight list of this round of training
		var deserializedWeights WeightsBiases
		json.Unmarshal(msg.Data, &deserializedWeights)

		weightsToAvg := append(state.weights, deserializedWeights)
		state.weights = weightsToAvg

		// All nodes finished training
		// send the weights to the aggregator
		if state.actorsTraining == 0 {
			fmt.Printf("All actors finished training, the round %v has ended at %v\n", state.roundsTrained, time.Now())
			msg5 := CalculateAverage{Weights: weightsToAvg}
			context.Send(state.aggregatorPID, msg5)
			state.behavior.Become(state.WeightCalculating)
		}
	}
}

// This state signifies that one round of training is over
// so we have to wait until the weights are averaged and updated
// to start a new training round
func (state *CoordinatorActor) WeightCalculating(context actor.Context) {

	switch msg := context.Message().(type) {
	case startTraining:
		fmt.Println("Waiting for new weight model!")
		fmt.Printf("msg: %v\n", msg)
	case weightsUpdated:
		state.behavior.Become(state.Training)
		msg6 := startTraining{}
		context.Send(state.parentPID, msg6)
	}
}

func (state *AggregatorActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case PidsStruct:
		if msg.initPID == nil {
			return
		}
		state.parentPID = msg.initPID
		state.selfPID = msg.aggregatorPID
		// Once all the training actors have finished
	// we send all the weights to the aggregator
	// that calculates the average
	// Once it is done it tell the coordinator that it has finished
	// so it can start another round of training
	case CalculateAverage:
		fmt.Printf("Averaging the weights %v\n", time.Now())
		fmt.Println("Length of all weights array:", len(msg.Weights))
		FederatedAverageAlgo(msg.Weights)
		fmt.Printf("Weights have been averaged %v\n", time.Now())
		msg7 := weightsUpdated{}
		context.Send(state.parentPID, msg7)
	}

}

func calculateAverage_L1w(arr [][13][64]float64) [13][64]float64 {
	len := float64(len(arr))
	var res [13][64]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}

func calculateAverage_L2w(arr [][64][128]float64) [64][128]float64 {
	len := float64(len(arr))
	var res [64][128]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}
func calculateAverage_L3w(arr [][128][128]float64) [128][128]float64 {
	len := float64(len(arr))
	var res [128][128]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}

func calculateAverage_L4w(arr [][128][64]float64) [128][64]float64 {
	len := float64(len(arr))
	var res [128][64]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}

func calculateAverage_L5w(arr [][64][32]float64) [64][32]float64 {
	len := float64(len(arr))
	var res [64][32]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}

func calculateAverage_L6w(arr [][32][1]float64) [32][1]float64 {
	len := float64(len(arr))
	var res [32][1]float64

	for _, arr2d := range arr {
		for i, arr1d := range arr2d {
			for j, val := range arr1d {
				res[i][j] += val
			}
		}
	}

	for i, row := range res {
		for j := range row {
			res[i][j] /= len
		}
	}

	return res
}

func calculateAverage_L1b(arr [][64]float64, res [64]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}

func calculateAverage_L2b(arr [][128]float64, res [128]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}
func calculateAverage_L3b(arr [][128]float64, res [128]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}

func calculateAverage_L4b(arr [][64]float64, res [64]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}

func calculateAverage_L5b(arr [][32]float64, res [32]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}

func calculateAverage_L6b(arr [][1]float64, res [1]float64) {
	len := float64(len(arr))
	for _, row := range arr {
		for i, val := range row {
			res[i] += val
		}
	}

	for i := range res {
		res[i] /= len
	}
}

func FederatedAverageAlgo(weights []WeightsBiases) {
	var layer1_w_calc [][13][64]float64
	var layer2_w_calc [][64][128]float64
	var layer3_w_calc [][128][128]float64
	var layer4_w_calc [][128][64]float64
	var layer5_w_calc [][64][32]float64
	var layer6_w_calc [][32][1]float64

	var layer1_b_calc [64]float64
	var layer2_b_calc [128]float64
	var layer3_b_calc [128]float64
	var layer4_b_calc [64]float64
	var layer5_b_calc [32]float64
	var layer6_b_calc [1]float64

	var layer1_b_arr [][64]float64
	var layer2_b_arr [][128]float64
	var layer3_b_arr [][128]float64
	var layer4_b_arr [][64]float64
	var layer5_b_arr [][32]float64
	var layer6_b_arr [][1]float64

	for _, val := range weights {
		layer1_w_calc = append(layer1_w_calc, val.Layer_1_weights)
		layer2_w_calc = append(layer2_w_calc, val.Layer_2_weights)
		layer3_w_calc = append(layer3_w_calc, val.Layer_3_weights)
		layer4_w_calc = append(layer4_w_calc, val.Layer_4_weights)
		layer5_w_calc = append(layer5_w_calc, val.Layer_5_weights)
		layer6_w_calc = append(layer6_w_calc, val.Layer_6_weights)

		layer1_b_arr = append(layer1_b_arr, val.Layer_1_biases)
		layer2_b_arr = append(layer2_b_arr, val.Layer_2_biases)
		layer3_b_arr = append(layer3_b_arr, val.Layer_3_biases)
		layer4_b_arr = append(layer4_b_arr, val.Layer_4_biases)
		layer5_b_arr = append(layer5_b_arr, val.Layer_5_biases)
		layer6_b_arr = append(layer6_b_arr, val.Layer_6_biases)
	}

	go calculateAverage_L1b(layer1_b_arr, layer1_b_calc)
	go calculateAverage_L2b(layer2_b_arr, layer2_b_calc)
	go calculateAverage_L3b(layer3_b_arr, layer3_b_calc)
	go calculateAverage_L4b(layer4_b_arr, layer4_b_calc)
	go calculateAverage_L5b(layer5_b_arr, layer5_b_calc)
	go calculateAverage_L6b(layer6_b_arr, layer6_b_calc)

	globalModel.Layer_1_weights = calculateAverage_L1w(layer1_w_calc)
	globalModel.Layer_2_weights = calculateAverage_L2w(layer2_w_calc)
	globalModel.Layer_3_weights = calculateAverage_L3w(layer3_w_calc)
	globalModel.Layer_4_weights = calculateAverage_L4w(layer4_w_calc)
	globalModel.Layer_5_weights = calculateAverage_L5w(layer5_w_calc)
	globalModel.Layer_6_weights = calculateAverage_L6w(layer6_w_calc)

	globalModel.Layer_1_biases = layer1_b_calc
	globalModel.Layer_2_biases = layer2_b_calc
	globalModel.Layer_3_biases = layer3_b_calc
	globalModel.Layer_4_biases = layer4_b_calc
	globalModel.Layer_5_biases = layer5_b_calc
	globalModel.Layer_6_biases = layer6_b_calc
}

var globalModel WeightsBiases
var localAddress string
var port int

func main() {
	localAddress = "192.168.188.14"
	port = 8000

	file, openingErr := os.Open("weightModel.json")
	if openingErr != nil {
		panic(openingErr)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	deserializingError := decoder.Decode(&globalModel)
	if deserializingError != nil {
		panic(deserializingError)
	}

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

	remoteConfig := remote.Configure(localAddress, port)
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	spawnResponse, err1 := remoting.SpawnNamed("192.168.188.14:8091", "training_actor", "training_actor", 10*time.Second)
	spawnResponse1, err2 := remoting.SpawnNamed("192.168.188.237:8090", "training_actor", "training_actor", 10*time.Second)

	if err1 != nil {
		fmt.Println("panika 1")
		panic(err1)

	}
	if err2 != nil {
		fmt.Println("panika 2")
		panic(err2)
	}
	spawnedActorMessage := spawnedRemoteActor{remoteActorPID: spawnResponse.Pid}
	spawnedActorMessage1 := spawnedRemoteActor{remoteActorPID: spawnResponse1.Pid}

	rootContext.Send(pid, spawnedActorMessage)
	rootContext.Send(pid, spawnedActorMessage1)

	rootContext.Send(pid, startTraining{})

	_, _ = console.ReadLine()
}
