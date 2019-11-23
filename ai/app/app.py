from flask import Flask, jsonify
from flask_cors import CORS

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)


@app.route('/', methods=['GET'])
def hello_world():
    return 'Hello World!'


@app.route('/model/predict', methods=['POST'])
def model_predict():
    return "model predict"


@app.route('/model/train', methods=['POST'])
def model_train():
    return "model train"


if __name__ == '__main__':
    app.run(host='0.0.0.0')
