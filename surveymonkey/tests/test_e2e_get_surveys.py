import requests

# End-to-End (E2E) Testing
def test_e2e_get_surveys():
    response = requests.get("http://127.0.0.1:5000/surveys")
    assert response.status_code == 200
    assert "data" in response.json()  # Optional: Validate schema
