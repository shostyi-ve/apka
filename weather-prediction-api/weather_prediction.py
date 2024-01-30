import pandas as pd
from prophet import Prophet
from datetime import datetime


def predict_column(hourly_data, column_name, hours):
    data = pd.DataFrame()
    data["ds"] = hourly_data["time"]
    data["ds"] = data["ds"].apply(
        lambda timestamp_str: datetime.strptime(
            timestamp_str, "%Y-%m-%dT%H:%M"
        ).strftime("%Y-%m-%d %H:%M:%S")
    )

    data["y"] = hourly_data[column_name]

    model = Prophet()

    model.fit(data)

    future = model.make_future_dataframe(periods=hours, freq="h")

    forecast = model.predict(future)

    forecast = forecast[data["ds"].max() < forecast["ds"]]

    forecast_needed_columns = ["ds", "yhat"]
    forecast = forecast[forecast_needed_columns]
    forecast.rename(columns={"yhat": column_name, "ds": "time"}, inplace=True)

    return forecast


def get_prediction_for_24h(json_data):
    source_data_dict = json_data
    hours = 24
    forecast_all = pd.DataFrame()
    for c in source_data_dict.keys():
        if c == "time":
            continue

        if forecast_all.empty:
            forecast_all = predict_column(source_data_dict, c, hours)
        else:
            forecast_all = forecast_all.merge(
                predict_column(source_data_dict, c, hours), on="time"
            )
    forecast_all["time"] = forecast_all["time"].dt.strftime("%Y-%m-%dT%H:%M")
    return forecast_all
