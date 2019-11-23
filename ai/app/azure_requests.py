import json
import os

from requests import Request, Session


def request_prediction(file_path, endpoint, prediction_key):
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