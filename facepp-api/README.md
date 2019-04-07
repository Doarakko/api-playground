# facepp-api-clip-face
## Overview
「[Face++ Detect API](https://www.faceplusplus.com/face-detection/)」を使用して, ローカル画像から顔を検出し, 顔の切り取りを行います.  
顔の角度・検出した顔の数・瞳孔間の距離から, 一定条件を満たす顔のみ切り取りを行うプログラムを作成しました.
![image.png](https://qiita-image-store.s3.amazonaws.com/0/245792/b8de8352-0e4d-6e99-e606-325a105ac94e.png)

## Description
「[Face++ Detect API](https://www.faceplusplus.com/face-detection/)」を使用して, ローカル画像から顔を検出し, 顔の切り取りを行います.  
顔が正面を向いていて, 一定サイズ以上大きい顔画像のみを取得することを目的とにします.  
そのため, 以下の条件を満たす顔画像は顔の切り取りは行いません.

・検出した顔の数が0または複数  
・顔の横揺れの角度が15°以上または-15°以下  
・顔の縦揺れの角度が10°以上または-10°以下  
・瞳孔間の距離が40ピクセル未満  

取得した顔の特徴量を後で使用することを想定するため, レスポンスをJSON形式で保存しました.  
その後, JSONファイルをロードして顔の切り取りを行います.

## Requirement
・mac os 10.13.3  
・anaconda3-5.1.0  
・pip 9.0.3  
・opencv-python 3.4.0.12  

・Face++ Detect API v3.0  
・Face++ アカウント  
　- アカウント作成の際に電話番号が必要  
・ Face++ API Key

## Usage
1. pip で opencv をインストール
```
$ pip install opencv-python
```

2. git clone
```
$ git clone https://github.com/Doarakko/api-challenge
$ cd facepp-api-clip-face
```
```
./
├── clip_face.py
├── data
├── detect_image.py
└── image
    ├── error
    ├── face
    └── original
```

3. API KEY と API SECRET を入力
![facepp.jpg](https://qiita-image-store.s3.amazonaws.com/0/245792/0218f4ba-b158-7f8e-3398-2b1b0d73f52b.jpeg)
```python:detect_image.py
API_KEY = 'aaabbbccc'
API_SECRET = 'xxxyyyzzz'
```

4. 顔認識する画像を ./image/original 下に置く  
数に制限はなく, 名前は自由です
```
./image/original/
├── person0
│   ├── person0_0.jpg
│   ├── person0_1.jpg
│   ├── person0_2.jpg
│   ├── person0_3.jpg
│   └── person0_4.jpg
├── person1
│   ├── person1_0.jpg
│   ├── person1_1.jpg
│   ├── person1_2.jpg
│   ├── person1_3.jpg
│   └── person1_4.jpg
└── person2
    ├── person2_0.jpg
    ├── person2_1.jpg
    ├── person2_2.jpg
    ├── person2_3.jpg
    └── person2_4.jpg
```
5. [Face++ Detect API](https://www.faceplusplus.com/face-detection/) にリクエストを送り, レスポンスをJSON形式で保存します
```
$ python detect_image.py
```

6. 保存した JSONファイルをロードして, 取得した数値から顔を切り取ります
```
$ python clip_face.py
```

## Hints
### detect_image.py  
'return_attributes' は, ',' の後にスペースを入れるとエラーになるので注意してください

```python:detect_image.py
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
```
### clip_face.py
目的に合わせて下部分を変更・追加・削除してください

```python:clip_face.py
#指定した条件を満たす顔画像か判断する関数
def judge_noise_image(resources):
    #検出された顔の数
    face_num = len(resources['faces'])
    # 検出された顔の数が1以外の場合
    if face_num != 1:
        print('[Remove] face_num = {}'.format(face_num), end=' ')
        return -1

    #ランドマークの数
    landmark_num = len(resources['faces'][0]['landmark'])
    # 一部のランドマークが取得できなかった場合
    if landmark_num != 83:
        print('[Remove] landmark_num = {}'.format(landmark_num), end=' ')
        return -1

    #顔の向きを取得
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
    right_eye_pupil_y = resources['faces'][0]['landmark']['right_eye_pupil']['y']
    # 右目の瞳孔のy座標
    right_eye_pupil_x = resources['faces'][0]['landmark']['right_eye_pupil']['x']
    # 瞳孔間の距離
    pupil_dis = math.sqrt((left_eye_pupil_y - right_eye_pupil_y) ** 2 + (left_eye_pupil_x - right_eye_pupil_x) ** 2)
    # 瞳孔間の距離が40ピクセル未満の場合
    if pupil_dis < 40:
        print('[Remove] pupil_dis = {}'.format(pupil_dis), end=' ')
        return -1
    return 0
```
・[「Face++ Detect API」でローカル画像から一定条件を満たす顔を切り取る in Python](https://qiita.com/Doarakko/items/bf33d9eb102871224ce1)
## Author
[Doarakko](https://github.com/Doarakko)

## License
MIT