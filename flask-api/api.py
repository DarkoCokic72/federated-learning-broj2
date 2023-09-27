import train
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/train', methods = ['POST'])
def train_ann():

    received_data = request.json
    weights = train.train_neural_network(received_data)

    serialized = [w.tolist() for w in weights]
    resp = {
        "layer_1_weights": serialized[0],
        "layer_1_biases": serialized[1],
        "layer_2_weights": serialized[2],
        "layer_2_biases": serialized[3],
        "layer_3_weights": serialized[4],
        "layer_3_biases": serialized[5],
        "layer_4_weights": serialized[6],
        "layer_4_biases": serialized[7],
        "layer_5_weights": serialized[8],
        "layer_5_biases": serialized[9],
        "layer_6_weights": serialized[10],
        "layer_6_biases": serialized[11],

    }
    return jsonify(resp), 200

if __name__ == '__main__':
    app.run()

    print(f"Running on port: {app.config['SERVER_NAME'].split(':')[-1]}")