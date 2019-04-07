import time
import os
import re
import json
import base64
import requests
import shutil
import glob


API_KEY = 'aaabbbccc'
API_SECRET = 'xxxyyyzzz'


# APIにリクエストを送る関数
def detect_image(img_path):
    endpoint = 'https://api-us.faceplusplus.com'
    img_file = base64.encodestring(open(img_path, 'rb').read())
    try:
        response = requests.post(
            endpoint + '/facepp/v3/detect',
            {
                'api_key': API_KEY,
                'api_secret': API_SECRET,
                # 'image_url': img_url,
                'image_base64': img_file,
                'return_landmark': 1,
                'return_attributes': 'headpose,eyestatus,facequality,mouthstatus,eyegaze'
            }
        )
        # 5秒スリープ
        time.sleep(5)
        # レスポンスのステータスコードが200以外の場合
        if response.status_code != 200:
            # 画像の親ディレクトリ名を取得
            dir_name = re.search(r'./image/original/(.+)/.+', img_path)
            dir_name = dir_name.group(1)
            move_dir = './image/error/' + dir_name

            # 移動先のディレクトリを作成
            if not os.path.exists(move_dir):
                os.makedirs(move_dir)
            # 画像を移動
            shutil.move(img_path, move_dir)
            print('[Error] {} {}'.format(img_path, response))
            return -1
        resources = response.json()
        return resources
    except Exception as e:
        # 画像の親ディレクトリ名を取得
        dir_name = re.search(r'./image/original/(.+)/.+', img_path)
        dir_name = dir_name.group(1)
        move_dir = './image/error/' + dir_name
        # 移動先のディレクトリを作成
        if not os.path.exists(move_dir):
            os.makedirs(move_dir)
        # 画像を移動
        shutil.move(img_path, move_dir)
        print('[Error] {}'.format(img_path))
        return -1


# APIのResponceをjson形式で保存する関数
def save_json(resources, img_path):
    # 画像の親ディレクトリ名を取得
    save_dir = re.search(r'./image/original/(.+)/.+', img_path)
    save_dir = './data/' + save_dir.group(1)
    # 保存先のディレクトリを作成
    if not os.path.exists(save_dir):
        os.makedirs(save_dir)
    # 画像のファイル名を取得
    file_name = re.search(r'./image/original/.+/(.+).jpg', img_path)
    file_name = file_name.group(1)
    save_name = file_name + '.json'
    file_path = save_dir + '/' + save_name
    with open(file_path, 'w') as f:
        json.dump(resources, f)
    print('[Save] {}'.format(save_name))


if __name__ == '__main__':
    # 画像のリストを取得
    img_path_list = glob.glob('./image/original/*/*.jpg')
    img_path_list.sort()
    for img_path in img_path_list:
        save_dir = re.search(r'./image/original/(.+)/.+', img_path)
        save_dir = './data/' + save_dir.group(1)
        # 画像のファイル名を取得
        file_name = re.search(r'./image/original/.+/(.+).jpg', img_path)
        file_name = file_name.group(1)
        save_name = file_name + '.jpg'
        file_path = save_dir + '/' + save_name
        if os.path.exists(file_path):
            print('[Skip] {}'.format(file_path))
        else:
            # APIにリクエストを送る
            resources = detect_image(img_path)
            if resources != -1:
                save_json(resources, img_path)
