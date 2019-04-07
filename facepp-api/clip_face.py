import time
import os
import re
import cv2
import json
import glob
import math


# 指定した条件を満たす顔画像か判断する関数
def judge_noise_image(resources):
    # 検出された顔の数
    face_num = len(resources['faces'])
    # 検出された顔の数が1以外の場合
    if face_num != 1:
        print('[Remove] face_num = {}'.format(face_num), end=' ')
        return -1

    # ランドマークの数
    landmark_num = len(resources['faces'][0]['landmark'])
    # 一部のランドマークが取得できなかった場合
    if landmark_num != 83:
        print('[Remove] landmark_num = {}'.format(landmark_num), end=' ')
        return -1

    # 顔の向きを取得
    yaw = resources['faces'][0]['attributes']['headpose']['yaw_angle']
    pitch = resources['faces'][0]['attributes']['headpose']['pitch_angle']
    # roll = resources['faces'][0]['attributes']['headpose']['roll_angle']

    # yaw角度が15°以上または-15°以下の場合
    if yaw >= 15 or yaw <= -15:
        print('[Remove] yaw = {}'.format(yaw), end=' ')
        return -1
    # pitch角度が10°以上または-10°以下の場合
    elif pitch >= 10 or pitch <= -10:
        print('[Remove] pitch = {}'.format(pitch), end=' ')
        return -1

    # 左目の瞳孔のy座標
    left_eye_pupil_y = resources['faces'][0]['landmark']['left_eye_pupil']['y']
    # 左目の瞳孔のx座標
    left_eye_pupil_x = resources['faces'][0]['landmark']['left_eye_pupil']['x']
    # 右目の瞳孔のy座標
    right_eye_pupil_y = resources['faces'][0][
        'landmark']['right_eye_pupil']['y']
    # 右目の瞳孔のy座標
    right_eye_pupil_x = resources['faces'][0][
        'landmark']['right_eye_pupil']['x']
    # 瞳孔間の距離
    pupil_dis = math.sqrt((left_eye_pupil_y - right_eye_pupil_y) ** 2 + (
        left_eye_pupil_x - right_eye_pupil_x) ** 2)
    # 瞳孔間の距離が40ピクセル未満の場合
    if pupil_dis < 40:
        print('[Remove] pupil_dis = {}'.format(pupil_dis), end=' ')
        return -1
    return 0


# 顔を切り取る関数
def clip_face(resources, img_path):
    try:
        top = resources['faces'][0]['face_rectangle']['top']
        left = resources['faces'][0]['face_rectangle']['left']
        width = resources['faces'][0]['face_rectangle']['width']
        height = resources['faces'][0]['face_rectangle']['height']
        # オリジナルの画像をロード
        origin_image = cv2.imread(img_path)
        face_img = origin_image[top:top + height, left:left + width]
        save_dir = re.search(r'./image/original/(.+)/.+', img_path)
        save_dir = save_dir.group(1)
        save_dir = './image/face/' + save_dir
        # 画像を保存するディレクトリを作成
        if not os.path.exists(save_dir):
            os.makedirs(save_dir)
        # 画像のファイル名
        file_name = re.search(r'./image/original/.+/(.+).jpg', img_path)
        file_name = file_name.group(1)
        save_name = 'face_' + file_name + '.jpg'
        save_path = save_dir + '/' + save_name
        # 切り取った顔画像を保存
        cv2.imwrite(save_path, face_img)
        print('[Clip] {0}'.format(img_path))
        # 2秒スリープ
        time.sleep(2)
    except Exception as e:
        print('[Error] {0}'.format(img_path))


if __name__ == '__main__':
    # jsonファイルのリストを取得
    json_path_list = glob.glob('./data/*/*.json')
    json_path_list.sort()
    for json_path in json_path_list:
        # jsonファイルをロード
        with open(json_path, 'r') as f:
            resources = json.load(f)
        if judge_noise_image(resources) != - 1:
            # 画像のファイル名
            file_name = re.search(r'./data/.+/(.+).json', json_path)
            file_name = file_name.group(1)
            # 画像の親ディレクトリ名
            dir_name = re.search(r'./data/(.+)/.+', json_path)
            dir_name = dir_name.group(1)
            # 画像のパス
            img_path = './image/original/' + dir_name + '/' + file_name + '.jpg'
            # 顔を切り取る
            clip_face(resources, img_path)
        else:
            print('{}'.format(json_path))
