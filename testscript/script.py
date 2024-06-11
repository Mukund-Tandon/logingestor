import requests
import json
import time
from concurrent.futures import ThreadPoolExecutor

url = "http://127.0.0.1:8080/log"
data = {
  "timestamp": "2023-05-18T15:23:42Z",
  "level": "error",
  "message": "Lol",
  "resourceID": "server-1234"
}

headers = {
    "Content-Type": "application/json"
}

def send_post_request(url, data, headers):
    try:
        requests.post(url, data=json.dumps(data), headers=headers)
    except requests.exceptions.RequestException as e:
        print(f"Request failed: {e}")

def main():
    start_time = time.time()  # Record the start time
    with ThreadPoolExecutor(max_workers=50) as executor:
        for i in range(1000):
            executor.submit(send_post_request, url, data, headers)
    print("All requests have been sent.")
    end_time = time.time()  # Record the end time
    print(f"Time taken: {end_time - start_time} seconds")

if __name__ == "__main__":
    main()
