from flask import Flask

app = Flask(__name__)
NUMBERS = []

@app.route("/")
def index():
    return "Server up"

@app.route("/all", methods=['GET'])
def get_all():
    return ','.join(map(str, NUMBERS))

@app.route("/<int:value>", methods=['POST'])
def add(value):
    NUMBERS.append(value)
    return str(len(NUMBERS))

@app.route("/<int:value>", methods=['DELETE'])
def remove(value):
    try:
        NUMBERS.remove(value)
    except ValueError:
        print(f'Number {value} doesn\'t exist')
    return str(len(NUMBERS))

@app.route("/all", methods=['DELETE'])
def remove_all():
    global NUMBERS
    NUMBERS = []
    return str(len(NUMBERS))

# @app.route("/<values>", methods=['PUT'])
# def put_all_values(values):
#     global NUMBERS
#     NUMBERS = [int(v) for v in values.split(',')]
#     return str(len(NUMBERS))
