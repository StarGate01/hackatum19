from flask import Flask, jsonify, request, send_file
import requests
from flask_cors import CORS
from azure_requests import request_prediction
from dotenv import load_dotenv
import os
import threading

load_dotenv()

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)

endpoint = os.getenv('ENDPOINT')
prediction_key = os.getenv('PREDICTION_KEY')


class AsyncRequest(threading.Thread):
    def __init__(self, filepath, id):
        self.filepath = filepath
        self.id = id

    def run(self):
        try:
            cracked, uncracked = request_prediction(self.filepath, endpoint, prediction_key)
            cracked = int(cracked * 100)
            print(cracked)
            url = "/images/" + self.id + "/probability"
            data = jsonify({
                "probability": cracked
            })
            return requests.post(url, data=data)

        except:
            return None


@app.route('/model/predict', methods=['POST'])
async def model_predict():
    r = request.get_json()
    print(str(r))
    filename = r["id"]
    filepath = "/data/images/" + filename + ".jpg"
    asnycRequest = AsyncRequest(filepath=filepath, id=filename)
    asnycRequest.start()

    return "Request sent"


@app.route('/model/train', methods=['POST'])
def model_train():
    r = request.get_json()
    filename = r["id"]
    is_cracked = r["iscracked"]
    filepath = "/data/images/" + filename + ".jpg"

    return "Model trained"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
