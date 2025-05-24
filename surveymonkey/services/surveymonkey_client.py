# services/surveymonkey.py

from services.base_client import BaseAPIClient
from config import SURVEYMONKEY_API_TOKEN, SURVEYMONKEY_API_BASE_URL
from tenacity import retry, stop_after_attempt, wait_fixed
from typing import Any


class SurveyMonkeyAPI(BaseAPIClient):

    def __init__(self):
        super().__init__(SURVEYMONKEY_API_BASE_URL, SURVEYMONKEY_API_TOKEN)

    @retry(stop=stop_after_attempt(3), wait=wait_fixed(2))
    def list_surveys(self) -> Any:
        return self.get("/surveys")

    def get_survey_details(self, survey_id: str) -> Any:
        return self.get(f"/surveys/{survey_id}")


