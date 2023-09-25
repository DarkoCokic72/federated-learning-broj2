import pandas as pd
import numpy as np

import tensorflow as tf
from tensorflow import keras
from keras import layers

from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler

def train_neural_network():
    df = pd.read_csv('heart.csv', delimiter=',')
    columns = df.columns.ravel()

    output = np.where(columns == 'output')[0][0]

    X = df.iloc[:, :output]
    Y = df.iloc[:, output:]

    X_train, X_test, Y_train, Y_test = train_test_split(X, Y, train_size=0.8, shuffle=True)

    scaler = StandardScaler()

    X_train = pd.DataFrame(scaler.fit_transform(X_train), columns=X_train.columns)
    X_test = pd.DataFrame(scaler.fit_transform(X_test), columns=X_test.columns)

    model = keras.Sequential([
        layers.Dense(64, activation='sigmoid', input_shape=(13,)),
        layers.Dense(128, activation='sigmoid'),
        layers.Dense(128, activation='sigmoid'),
        layers.Dense(64, activation='sigmoid'),
        layers.Dense(32, activation='sigmoid'),
        layers.Dense(1, activation='sigmoid')
    ])

    model.compile(optimizer='adagrad', loss='binary_crossentropy', metrics=['accuracy'])

    model.fit(X_train, Y_train, epochs=30, verbose=1)

    y_pred = model.predict(X_test)

    weights = model.get_weights()

    return weights

