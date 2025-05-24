from services.surveymonkey_client import SurveyMonkeyAPI
from unittest.mock import patch

# Unit testing list_surveys() method
# Testing only list_surveys(), mocking out its dependency (get method).
def test_list_surveys_success():
    api = SurveyMonkeyAPI()

    mock_response = {"data": [{"id": "1", "title": "Feedback"}]}
    with patch.object(api, "get", return_value=mock_response) as mock_get:
        response = api.list_surveys()
        assert response == mock_response
        mock_get.assert_called_once_with("/surveys")
