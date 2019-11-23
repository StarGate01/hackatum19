from flask import Flask, jsonify
from flask_cors import CORS

app = Flask(__name__)
app.config.from_object('config.DevelopmentConfig')
cors = CORS(app)


@app.route('/', methods=['GET'])
def hello_world():
    return 'Hello World!'


if __name__ == '__main__':
    app.run(host='0.0.0.0')
