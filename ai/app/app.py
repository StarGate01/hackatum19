from flask import Flask, jsonify, request, send_file
from flask_cors import CORS
from azure_requests import request_prediction
from dotenv import load_dotenv
import os

load_dotenv()

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)

endpoint = os.getenv('ENDPOINT')
prediction_key = os.getenv('PREDICTION_KEY')

@app.route('/model/predict', methods=['POST'])
def model_predict():
    r = request.get_json()
    print(str(r))
    filename = r["id"]
    filepath = "/data/images/" + filename + ".jpg"
    cracked, uncracked = request_prediction(filepath, endpoint, prediction_key)
    cracked = int(cracked * 100)
    print(cracked)
    return jsonify({
        "probability": cracked
    })


@app.route('/model/train', methods=['POST'])
def model_train():
    r = request.get_json()
    filename = r["id"]
    is_cracked = r["iscracked"]
    filepath = "/data/images/" + filename + ".jpg"

    return "Model trained"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
