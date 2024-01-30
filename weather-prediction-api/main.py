from flask import Flask, request
from weather_prediction import get_prediction_for_24h

app = Flask(__name__)


@app.route("/api/weather/get_prediction_24h", methods=["GET"])
def get_prediction():
    json_data = request.get_json()
    forecast = get_prediction_for_24h(json_data)
    result = {}
    for c in forecast.columns:
        result[c] = forecast[c].tolist()

    return result
