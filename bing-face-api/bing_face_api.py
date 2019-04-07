import os
import glob
import time
import re
import cv2
import requests
import json
from urllib.request import urlretrieve

API_KEY = 'your key'


# 顔認識する画像のファイル名を取得する関数
def get_image_path_list(search_dir):
    # ファイルを取得
    img_path_list = glob.glob(search_dir + '*/*.jpg')
    return img_path_list


# 画像を保存するディレクトリを作成・取得する関数
def get_save_dir(img_path):
    # 画像を保存するディレクトリを指定
    save_dir = re.search(r'./image/original/(.+)/.+', img_path)
    save_dir = "./image/face/" + save_dir.group(1)
    # 画像を保存するディレクトリを作成
    if not os.path.exists(save_dir):
        os.makedirs(save_dir)
    return save_dir


# 保存する画像のファイル名を取得する関数
def get_save_name(img_path):
    save_name = re.search(r'./image/original/.+/(.+)', img_path)
    save_name = save_name.group(1)
    return save_name


# 画像を保存するパスを取得する関数
def get_save_path(save_dir, save_name):
    save_path = save_dir + "/" + save_name
    return save_path


# 顔認識する関数
def detect_image(img_path_list):
    endpoint = 'https://westcentralus.api.cognitive.microsoft.com'
    headers = {
        'Content-Type': 'application/octet-stream',
        'Ocp-Apim-Subscription-Key': API_KEY,
    }
    params = {
        'returnFaceId': 'true',
        'returnFaceLandmarks': 'false',
        # 'returnFaceAttributes': 'age,gender,headPose,smile,facialHair,glasses,emotion,hair,makeup,occlusion,accessories,blur,exposure,noise',
    }
    for img_path in img_path_list:
        # ファイルを開く
        img_file = open(img_path, 'rb').read()
        response = requests.request(
            'POST', endpoint + '/face/v1.0/detect', data=img_file, headers=headers, params=params)
        resources = response.json()
        # 顔の部分を切り取る
        clip_face(resources, img_path)


# 顔の部分を切り取るする関数
def clip_face(resources, img_path):
    # 顔の位置を取得
    try:
        for resource in resources:
            top = resource['faceRectangle']['top']
            left = resource["faceRectangle"]["left"]
            width = resource["faceRectangle"]["width"]
            height = resource["faceRectangle"]["height"]
        # オリジナルの画像を読み込み
        origin_image = cv2.imread(img_path)
        # 顔の部分を抜き取る
        face_img = origin_image[top:top+height, left:left+width]
        # 画像を保存するディレクトリを作成・取得
        save_dir = get_save_dir(img_path)
        # 保存する画像のファイル名
        save_name = "face_" + get_save_name(img_path)
        # 画像を保存するパス
        save_path = get_save_path(save_dir, save_name)

        # 切り取った顔の画像を保存
        cv2.imwrite(save_path, face_img)
        # 3秒スリープ
        time.sleep(3)
        print("[Detect] {0}".format(get_save_name(img_path)))
    except Exception as e:
        print("[Error] {0}".format(get_save_name(img_path)))
