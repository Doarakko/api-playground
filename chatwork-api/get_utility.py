import requests

ENDPOINT = 'https://api.chatwork.com/v2/'
API_TOKEN = 'your api token'
HEADERS = {'X-ChatWorkToken': API_TOKEN}


def get_task_list():
    url = '{}/my/tasks'.format(ENDPOINT)
    params = {
        # 'assigned_by_account_id': xxxx,
        'staus': 'open',
    }
    response = requests.get(url, headers=HEADERS, params=params)
    resources = response.json()


def get_task_info():
    url = '{}/my/tasks'.format(ENDPOINT)
    params = {
        # 'assigned_by_account_id': xxxx,
        'staus': 'open',
    }
    response = requests.get(url, headers=HEADERS, params=params)
    resources = response.json()
