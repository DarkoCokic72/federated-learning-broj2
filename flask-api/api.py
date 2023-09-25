import train
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/train', methods = ['POST'])
def train_ann():
    weights = train.train_neural_network()

    serialized = [w.tolist() for w in weights]
    return jsonify(serialized), 200

if __name__ == '__main__':
    app.run()

    print(f"Running on port: {app.config['SERVER_NAME'].split(':')[-1]}")