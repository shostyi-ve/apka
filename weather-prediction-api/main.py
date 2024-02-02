from flask import Flask, request
from weather_prediction import get_prediction_for_24h

app = Flask(__name__)


@app.route("/api/weather/get_prediction_24h", methods=["GET"])
def get_prediction():
    json_data = request.get_json()
    hourly_data = json_data["hourly"]
    first_none_index = -1
    for k, v in hourly_data.items():
        if None in v:
            curr_none_index = v.index(None)
            if first_none_index == -1 or curr_none_index < first_none_index:
                first_none_index = curr_none_index
    for k, v in hourly_data.items():
        hourly_data[k] = v[:first_none_index]
    print(hourly_data)
    forecast = get_prediction_for_24h(hourly_data)
    result = {}
    for c in forecast.columns:
        result[c] = forecast[c].tolist()

    return result
