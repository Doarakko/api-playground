import requests
import re

ENDPOINT = 'https://api.chatwork.com/v2/'
API_TOKEN = 'your api token'
HEADERS = {'X-ChatWorkToken': API_TOKEN}


def post_messege(room_id, message):
    url = '{}rooms/{}/messages'.format(ENDPOINT, room_id)
    params = {'body': message}
    responce = requests.post(url, headers=HEADERS, params=params)
    return responce


def post_local_jpeg(room_id, file_path, message=''):
    url = '{}rooms/{}/files'.format(ENDPOINT, room_id)
    jpeg_bin = open(file_path, 'rb')
    file_name = re.search(r'.+/(.+\.jpg)', file_path)
    files = {
        'file': (file_name, jpeg_bin, 'image/jpeg'),
        'message': message,
    }
    responce = requests.post(url, headers=HEADERS, files=files)
    return responce
