import requests
from abc import ABC
from typing import Any
from requests.auth import HTTPBasicAuth


class BaseAPIClient(ABC):

    def __init__(self, base_url: str, token: str):
        self.base_url = base_url
        self.token = token
        self.session = requests.Session()
        self.session.headers.update({
            "Authorization": f"Bearer {self.token}",
            "Content-Type": "application/json"
        })

    def get(self, endpoint: str, params: Any = None) -> Any:
        url = f"{self.base_url}{endpoint}"
        response = self.session.get(url, params=params)
        response.raise_for_status()
        return response.json()

    def post(self, endpoint: str, data: Any = None):
        url = f"{self.base_url}{endpoint}"
        response = self.session.post(url, json=data)
        response.raise_for_status()
        return response.json()

    def basic_auth(self):

        token_url = "https://api.surveymonkey.com/oauth/token"
        auth = HTTPBasicAuth("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")

        data = {
            "grant_type": "authorization_code",
            "code": "code",
            "redirect_uri": "redirect_uri",
        }

        response = requests.post(token_url, auth=auth, data=data)
        return response.json()
