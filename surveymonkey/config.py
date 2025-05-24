from dotenv import load_dotenv
import os

load_dotenv()  # Load from .env

SURVEYMONKEY_API_BASE_URL = str(os.getenv("SURVEYMONKEY_API_BASE_URL"))
SURVEYMONKEY_API_TOKEN = str(os.getenv("SURVEYMONKEY_API_TOKEN"))