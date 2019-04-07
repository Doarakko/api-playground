import os
import json
import requests
from urllib.request import urlopen, urlretrieve

ENDPOIMT = 'https://api.flickr.com/services/rest/'
API_KEY = ''
DIR = './data/'
MAX_PER_PAGE_N = 500


def search_photos(word, photo_n=10, save_flg=True):
    if photo_n > MAX_PER_PAGE_N:
        per_page_n = MAX_PER_PAGE_N
        page_n = photo_n / MAX_PER_PAGE_N
    else:
        per_page_n = str(photo_n)
        page_n = str(1)

    params = {
        'method': 'flickr.photos.search',
        'api_key': API_KEY,
        'text': word,
        'per_page': per_page_n,
        'page': page_n,
        'format': 'json',
        'nojsoncallback': '1',
    }
    response = requests.get(ENDPOIMT, params=params)
    resource_json = response.json()
    if save_flg:
        save_json(resource_json, file_name='dog.json')
    return resource_json


def save_to_json(resource_json, file_name='hoge.json'):
    file_path = DIR + file_name
    with open(file_path, 'w') as f:
        json.dump(resource_json, f)


def get_url():
    url = 'https://farm{}.staticflickr.com/{}/{}_{}.jpg'.format(
        str(resource['farm']), resource['server'], resource['id'], resource['secret'])
    return url


if not os.path.exists(DIR):
    os.makedirs(DIR)
search_photos('犬', 600)
