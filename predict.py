
try:
    import sys, os,string,json

    os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'
    import tensorflow as tf
    from keras.models import load_model

    import numpy as np

    JSONStr =sys.argv[1]
    jsonData = json.loads(JSONStr)

    modelid = jsonData["modelid"]
    features = jsonData["features"]

    model = load_model("ml/models/"+str(modelid)+"/model.h5")

    predicted_value = model.predict(np.array([features]))

    response_predicted = float(predicted_value[0][0])
    response_predicted = float("%.2f" % response_predicted)
    response_predicted = str(response_predicted)


    print(response_predicted)
except:
    pass

