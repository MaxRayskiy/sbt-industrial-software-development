import requests

body = {"title": "Test Issue",
        "description": "any!",
        "email": "tester+100500@gmail.com",
        "private": True}
response = requests.post("http://localhost:8080/add", json=body)
assert response.status_code == 200
print(response.text)
