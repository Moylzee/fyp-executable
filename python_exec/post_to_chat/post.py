import requests

def post_to_chat(message, GROUP_URL):
    headers = {
        'Content-Type': 'application/json',
        "acceptType": "application/json",
    }

    payload = {
        "message": message
    }

    try:
        response = requests.post(GROUP_URL, json=payload, headers=headers)
        
        if response.status_code == 200:
            print("Message sent successfully!")
        else:
            print(f"Error sending message. Status code: {response.status_code}")
            print("Response:", response.text)
    
    except requests.exceptions.RequestException as e:
        print(f"Request failed: {e}")
