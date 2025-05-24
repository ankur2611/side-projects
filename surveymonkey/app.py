from flask import Flask, jsonify
from services.surveymonkey_client import SurveyMonkeyAPI
import logging
import traceback

app = Flask(__name__)
api = SurveyMonkeyAPI()

# Global Exception Handler
@app.errorhandler(Exception)
def handle_exception(e: Exception):
    logger.error("Unhandled Exception: %s", str(e))
    logger.debug(traceback.format_exc())
    
    return jsonify({
        "error": "An internal server error occurred.",
        "message": str(e)
    }), 500

# Set up basic logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@app.route("/")
def home():
    return {"message": "SurveyMonkey API Wrapper Running"}

@app.route("/surveys", methods=["GET"])
def get_surveys():
    surveys = api.list_surveys()
    return jsonify(surveys)

@app.route("/surveys/<survey_id>", methods=["GET"])
def get_survey(survey_id: str):
    survey = api.get_survey_details(survey_id)
    return jsonify(survey)

if __name__ == "__main__":
    app.run(debug=True)
