import requests


def test(self):
    body = {"title": "Test Issue",
            "description": "any!",
            "email": "tester+100500@gmail.com",
            "private": True}
    response = requests.post("http://192.168.1.110:8080/add", json=body)
    assert response.status_code == 201
    print(response.text)