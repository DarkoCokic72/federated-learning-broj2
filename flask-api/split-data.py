import pandas as pd
import numpy as np

df = pd.read_csv('heart.csv', delimiter=',')

df = df.sample(frac=1).reset_index(drop=True)

row, col = df.shape

middle = row // 2

df_1 = df.iloc[:middle, :]
df_2 = df.iloc[middle:, :]


df_1 = df_1.reset_index(drop=True)
df_2 = df_2.reset_index(drop=True)

df_1.to_csv('heart1.csv')
df_2.to_csv('heart2.csv')
