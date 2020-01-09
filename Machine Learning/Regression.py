import pandas as pd
import quandl
import math
import numpy as np
from sklearn import preprocessing, cross_validation, svm
from sklearn.linear_model import LinearRegression


df = quandl.get('WIKI/GOOGL')       #Data frame (df)

df = df[['Adj. Open',  'Adj. High',  'Adj. Low',  'Adj. Close', 'Adj. Volume']]
df['HL_PCT'] = (df['Adj. High'] - df['Adj. Low'])/df['Adj. Low'] * 100.0       #High vs Low Percent
df['PCT_change'] = (df['Adj. Close'] - df['Adj. Open'])/df['Adj. Open'] * 100.0       #Percent Change
df = df[['Adj. Close','HL_PCT','PCT_change','Adj. Volume']]

forecast_col = 'Adj. Close'         #New Forcast column
df.fillna(-99999, inplace=True)         #Fill 'not available', any not a number data (NaN) with -99999

forecast_out = int(math.ceil(0.01*len(df)))     #Forecast 1 percent of the entire length of the dataset. Math ceil will round up the number.
print(str(forecast_out) + ' days')
df['label'] = df[forecast_col].shift(-forecast_out)
df.dropna(inplace=True)

X = np.array(df.drop(['label']), 1)       #Features array, everything except 'label' column. Features are all in the inputs
y = np.array(df.['label'])       #Labels

X = preprocessing.scale(X)

print(df.head())