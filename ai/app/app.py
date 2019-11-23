from flask import Flask, jsonify, request, send_file
from flask_cors import CORS
from azure_requests import request_prediction
import random

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)


@app.route('/model/predict', methods=['POST'])
def model_predict():
    r = request.get_json()
    filename = r["filename"]
    filepath = "/data/images/" + filename + ".jpg"
    cracked, uncracked = request_prediction(filepath)
    cracked = int(cracked * 100)
    return jsonify({
        "probability": cracked
    })


@app.route('/model/train', methods=['POST'])
def model_train():
    r = request.get_json()
    filename = r["filename"]
    is_cracked = r["iscracked"]
    filepath = "/data/images/" + filename

    return "Model trained"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
