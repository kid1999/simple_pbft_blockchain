from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.route('/api')
def get_config():
    config = {
        "BlockSize":4,
        "BlockInterval":4,
        "Consensus":0,
        "Producers":["N0","N1"]
    }
    return config