import flask

app = flask.Flask(__name__)

word=open("readme.md", "r")


@app.route('/metrics', methods=['GET'])
def main():
    return word.read()


@app.route('/', methods=['GET'])
def home():
    return "<h1>Hello from Flask</h1>"

app.run(host='localhost',port='8888')