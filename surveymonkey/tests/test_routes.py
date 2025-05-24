import sys
import os
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

import pytest
from app import app
from unittest.mock import patch
from flask.testing import FlaskClient

# Fucntional Testing
@pytest.fixture
def client():
    app.config["TESTING"] = True
    with app.test_client() as client:
        yield client

def test_get_surveys(client: FlaskClient):
    mock_response = {
        "data": [{"id": "123", "title": "Customer Feedback"}]
    }

    with patch("app.api.list_surveys", return_value=mock_response):
        response = client.get("/surveys")
        assert response.status_code == 200
        assert response.json == mock_response

def test_get_survey_by_id(client: FlaskClient):
    mock_survey = {
        "id": "123",
        "title": "Customer Feedback"
    }

    with patch("app.api.get_survey_details", return_value=mock_survey):
        response = client.get("/surveys/123")
        assert response.status_code == 200
        assert response.json == mock_survey

def test_get_survey_by_id_500(client: FlaskClient):
    mock_survey = {
        "id": "123",
        "title": "Customer Feedback"
    }

    with app.test_client() as client:
        with patch("app.api.get_survey_details", return_value=mock_survey, side_effect=Exception("Something went wrong")):
            response = client.get("/surveys/123")
            print(response)
            assert response.status_code == 500
