import pandas as pd
import numpy as np

import tensorflow as tf
from tensorflow import keras
from keras import layers

from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler

def train_neural_network(weigts = None):
    df = pd.read_csv('flask-api/heart.csv', delimiter=',')
    columns = df.columns.ravel()

    output = np.where(columns == 'output')[0][0]

    X_train = df.iloc[:, :output]
    Y_train = df.iloc[:, output:]

    scaler = StandardScaler()

    X_train = pd.DataFrame(scaler.fit_transform(X_train), columns=X_train.columns)

    model = keras.Sequential([
        layers.Dense(64, activation='sigmoid', input_shape=(13,)),
        layers.Dense(128, activation='sigmoid'),
        layers.Dense(128, activation='sigmoid'),
        layers.Dense(64, activation='sigmoid'),
        layers.Dense(32, activation='sigmoid'),
        layers.Dense(1, activation='sigmoid')
    ])

    model.compile(optimizer='adagrad', loss='binary_crossentropy', metrics=['accuracy'])

    if weigts is not None:
        weigt_layer1 = weigts["Layer_1_weights"]
        weigt_layer2 = weigts["Layer_2_weights"]
        weigt_layer3 = weigts["Layer_3_weights"]
        weigt_layer4 = weigts["Layer_4_weights"]
        weigt_layer5 = weigts["Layer_5_weights"]
        weigt_layer6 = weigts["Layer_6_weights"]

        bias_layer1 = weigts["Layer_1_biases"]
        bias_layer2 = weigts["Layer_2_biases"]
        bias_layer3 = weigts["Layer_3_biases"]
        bias_layer4 = weigts["Layer_4_biases"]
        bias_layer5 = weigts["Layer_5_biases"]
        bias_layer6 = weigts["Layer_6_biases"]

        model.layers[0].set_weights([np.array(weigt_layer1), np.array(bias_layer1)])
        model.layers[1].set_weights([np.array(weigt_layer2), np.array(bias_layer2)])
        model.layers[2].set_weights([np.array(weigt_layer3), np.array(bias_layer3)])
        model.layers[3].set_weights([np.array(weigt_layer4), np.array(bias_layer4)])
        model.layers[4].set_weights([np.array(weigt_layer5), np.array(bias_layer5)])
        model.layers[5].set_weights([np.array(weigt_layer6), np.array(bias_layer6)])

    model.fit(X_train, Y_train, epochs=30, verbose=1)

    weights = model.get_weights()

    return weights

