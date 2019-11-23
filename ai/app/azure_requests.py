import json

from requests import Request, Session


endpoint = "https://southcentralus.api.cognitive.microsoft.com/customvision/v3.0/Prediction/0c667159-2b5c-4449-8aa5-a670fb31edd8/classify/iterations/Iteration2/image"
prediction_key = "284caa38d879463b90d0871031c19958"


def request_prediction(file_path):
    # data = open('./persistent-data/images/201402014_LIMI_001255.jpg', 'rb').read()
    data = open(file_path, 'rb').read()
    s = Session()
    headers = {'Content-Type': 'application/octet-stream', 'Prediction-Key': prediction_key}
    prepped = Request('POST', endpoint, headers, None, data, None).prepare()
    resp = s.send(prepped)
    respDict = json.loads(resp.content)
    crackedProb = ""
    uncrackedProb = ""
    predictions = respDict['predictions']
    for item in predictions:
        if item['tagName'] == "cracked":
            crackedProb = item['probability']
        elif item['tagName'] == "uncracked":
            uncrackedProb = item['probability']
    return crackedProb, uncrackedProb