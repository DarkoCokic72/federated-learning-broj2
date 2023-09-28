package actors

import (
	"fmt"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
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
	case CalculateAverage:
		fmt.Println("Averaging weights started")
		fmt.Println("Lenght of all weights: ", len(msg.Weights))
		FederatedAverageAlgo(msg.Weights)
		context.Send(state.parentPID, UpdatedWeights{})
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
