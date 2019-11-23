from flask import Flask, jsonify, request
from flask_cors import CORS
import random

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)


@app.route('/model/predict', methods=['POST'])
def model_predict():
    r = request.get_json()
    filepath = r["filepath"]

    return jsonify({
        "probability": random.randint(0, 99)
    })


@app.route('/model/train', methods=['POST'])
def model_train():
    r = request.get_json()
    filepath = r["filepath"]
    is_cracked = r["iscracked"]
    return "Model trained"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
