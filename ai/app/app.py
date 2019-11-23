from flask import Flask, jsonify, request, send_file
from flask_cors import CORS
from azure_requests import request_prediction
from azure_requests import upload_pic
from dotenv import load_dotenv
import os

load_dotenv()

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)

endpoint = os.getenv('ENDPOINT')
prediction_key = os.getenv('PREDICTION_KEY')
endpoint_upload = os.getenv('ENDPOINT_UPLOAD')
training_key = os.getenv("TRAINING_KEY")
predction_resource_id = os.getenv("PREDICTION_RESOURCE_ID")

@app.route('/model/predict', methods=['POST'])
def model_predict():
    r = request.get_json()
    print("SAKJHDAKJHDKAJSHD")
    print(r)
    filename = r["id"]
    filepath = "/data/images/" + filename + ".jpg"
    cracked, uncracked = request_prediction(filepath, endpoint, prediction_key)
    return


@app.route('/model/train', methods=['POST'])
def model_train():
    print("TRAIN")
    r = request.get_json()
    filename = r["id"]
    is_cracked = r["iscracked"]
    filepath = "/data/images/" + filename + ".jpg"

    upload_pic(endpoint_upload, training_key)

    return "Model trained"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
