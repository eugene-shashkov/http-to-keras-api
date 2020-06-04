try:
    import json,sys,os
    import numpy as np
    from datetime import datetime
    import time,os,threading,calendar,json,sys,math
    import tensorflow as tf

    from keras.models import Sequential
    from keras.layers import Dense, Activation
    from keras.models import load_model
    from pathlib import Path
    from keras.callbacks import ModelCheckpoint

    JSONStr =sys.argv[1]
    jsonData = json.loads(JSONStr)


    features = jsonData["features"]
    labels = jsonData["labels"]
    modelid = jsonData["modelid"]

    trainX = np.array(features)
    trainY = np.array(labels)

    model = Sequential()

    model.add(Dense(10, input_dim=len(features[0]), activation='relu'))
    model.add(Dense(1))

    model.compile(loss='mean_squared_error', optimizer='adam')

    filepath="weights.best.hdf5"
    checkpoint = ModelCheckpoint(filepath, monitor='val_acc', verbose=1, save_best_only=True, mode='max')
    callbacks_list = [checkpoint]
    model.fit(trainX, trainY, nb_epoch=1000, batch_size=2, verbose=2)

    os.mkdir("ml/models/"+str(modelid))
    model.save("ml/models/"+str(modelid)+"/model.h5")  # creates a HDF5 file 'model.h5'

except:
    pass
