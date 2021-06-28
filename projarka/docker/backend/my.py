import flask

app = flask.Flask(__name__)


@app.route('/', methods=['GET'])
def home():
    return "<h1>Hello from Flask</h1>"

app.run(host='localhost',port='8888')